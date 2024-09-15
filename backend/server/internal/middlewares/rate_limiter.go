package middlewares

import (
	"net/http"
	"sync"
	"time"

	"go.uber.org/ratelimit"
)

type RateLimiterEntry struct {
	limiter    ratelimit.Limiter
	lastAccess time.Time
}

var limiters = sync.Map{}

var globalLimiter = ratelimit.New(15)

const cleanupInterval = time.Minute * 5
const expirationDuration = time.Minute * 10

func GetRateLimiter(ip string) ratelimit.Limiter {
	now := time.Now()
	entry, ok := limiters.Load(ip)
	if !ok {
		newLimiter := ratelimit.New(5) // Set per-IP limit here
		limiters.Store(ip, RateLimiterEntry{limiter: newLimiter, lastAccess: now})
		return newLimiter
	}
	limiterEntry := entry.(RateLimiterEntry)
	limiterEntry.lastAccess = now
	limiters.Store(ip, limiterEntry)
	return limiterEntry.limiter
}

func StartCleanupRoutine() {
	ticker := time.NewTicker(cleanupInterval)
	go func() {
		for range ticker.C {
			now := time.Now()
			limiters.Range(func(key, value interface{}) bool {
				limiterEntry := value.(RateLimiterEntry)
				if now.Sub(limiterEntry.lastAccess) > expirationDuration {
					// Remove limiter if it's older than the expiration duration
					limiters.Delete(key)
				}
				return true
			})
		}
	}()
}

func RateLimiterMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		globalLimiter.Take() // Wait until the global rate limit allows this request

		// Extract the client's IP address
		ip := r.RemoteAddr
		limiter := GetRateLimiter(ip)
		limiter.Take()
		next.ServeHTTP(w, r)
	})
}
