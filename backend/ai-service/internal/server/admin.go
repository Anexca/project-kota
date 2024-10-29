package server

import (
	"common/ent"
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"go.uber.org/ratelimit"
)

var rlGenerateExamQuestionAndPopulateCache = ratelimit.New(1, ratelimit.Per(time.Minute))

func (s *Server) GenerateExamQuestionAndPopulateCache(w http.ResponseWriter, r *http.Request) {
	rlGenerateExamQuestionAndPopulateCache.Take()

	examId, err := ParseIDParam(r, "id")
	if err != nil {
		s.HandleError(w, err, "invalid exam id", http.StatusBadRequest)
		return
	}

	generatedExamData, err := s.examService.GenerateExamQuestionAndPopulateCache(r.Context(), examId)
	if err != nil {
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			s.HandleError(w, err, "exam not found", http.StatusNotFound)
			return
		}

		s.HandleError(w, err, "internal server error", http.StatusInternalServerError)
		return
	}

	err = s.WriteJson(w, http.StatusOK, &Response{Data: generatedExamData})
	if err != nil {
		s.HandleError(w, err, "something went wrong", http.StatusInternalServerError)
	}
}

func (s *Server) GenerateAllDescriptiveQuestions(w http.ResponseWriter, r *http.Request) {
	go func() {
		questions, err := s.examService.GenerateAllDescriptiveQuestions(context.Background())
		if err != nil {
			log.Println("something went wrong, ", err)
			return
		}

		log.Printf("generated %d questions\n", len(questions))
	}()

	err := s.WriteJson(w, http.StatusOK, &Response{Message: "exam generation started successfully"})
	if err != nil {
		s.HandleError(w, err, "something went wrong", http.StatusInternalServerError)
	}
}
