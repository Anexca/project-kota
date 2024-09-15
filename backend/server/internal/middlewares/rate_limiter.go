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

const cleanupInterval = time.Minute * 5
const expirationDuration = time.Minute * 10

func GetRateLimiter(ip string) ratelimit.Limiter {
	now := time.Now()
	entry, ok := limiters.Load(ip)
	if !ok {
		newLimiter := ratelimit.New(5)
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
					limiters.Delete(key)
				}
				return true
			})
		}
	}()
}

func RateLimiterMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		limiter := GetRateLimiter(ip)
		limiter.Take()
		next.ServeHTTP(w, r)
	})
}
