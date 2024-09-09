package services

import (
	"common/constants"

	goaway "github.com/TwiN/go-away"
)

type ProfanityService struct {
	profanityDetector *goaway.ProfanityDetector
}

func NewProfanityService() *ProfanityService {
	profanityDetector := goaway.NewProfanityDetector().WithCustomDictionary(
		constants.DefaultProfanities, constants.FalsePositiveProfanities, constants.FalseNegativeProfanities,
	)

	return &ProfanityService{
		profanityDetector: profanityDetector,
	}
}

func (p *ProfanityService) IsProfane(s string) bool {
	return p.profanityDetector.IsProfane(s)
}
