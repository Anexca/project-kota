package services

import (
	goaway "github.com/TwiN/go-away"

	"common/constants"
)

// ProfanityService is an implementation of ProfanityServiceInterface
type ProfanityService struct {
	profanityDetector *goaway.ProfanityDetector
}

// NewProfanityService creates a new instance of ProfanityService
func NewProfanityService() *ProfanityService {
	profanityDetector := goaway.NewProfanityDetector().WithCustomDictionary(
		constants.DefaultProfanities, constants.FalsePositiveProfanities, constants.FalseNegativeProfanities,
	)

	return &ProfanityService{
		profanityDetector: profanityDetector,
	}
}

// IsProfane checks if the provided string contains any profane words
func (p *ProfanityService) IsProfane(s string) bool {
	return p.profanityDetector.IsProfane(s)
}
