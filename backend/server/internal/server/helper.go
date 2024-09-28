package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"server/pkg/constants"
)

type ValidationError struct {
	Errors map[string]string
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("Validation failed for %d fields", len(v.Errors))
}

type ResponsePagination struct {
	CurrentPage int `json:"current_page,omitempty"`
	TotalPages  int `json:"total_pages,omitempty"`
	PerPage     int `json:"per_page,omitempty"`
	TotalItems  int `json:"total_items,omitempty"`
}

type Response struct {
	Error      bool               `json:"error"`
	Message    string             `json:"message,omitempty"`
	Data       any                `json:"data,omitempty"`
	Pagination ResponsePagination `json:"pagination,omitempty"`
}

func (app *Server) ReadJson(w http.ResponseWriter, r *http.Request, data any) error {
	maxSize := 1048576 // 1MB

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxSize))

	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(data); err != nil {
		return err
	}

	if err := dec.Decode(&struct{}{}); err != io.EOF {
		return errors.New("body must have single JSON object")
	}

	return nil
}

func (app *Server) WriteJson(w http.ResponseWriter, status int, data *Response, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for k, v := range headers[0] {
			w.Header()[k] = v
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func (app *Server) ErrorJson(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload Response
	payload.Error = true
	payload.Message = err.Error()

	fmt.Println("Error: ", payload.Message)

	return app.WriteJson(w, statusCode, &payload)
}

func (app *Server) SetCookie(w http.ResponseWriter, cookieName, cookieValue string, expiry time.Time) {
	cookie := &http.Cookie{
		Name:     cookieName,
		Value:    cookieValue,
		Expires:  expiry,
		HttpOnly: true, // Prevents access to the cookie via JavaScript
		Secure:   true, // Ensures the cookie is sent only over HTTPS
		Path:     "/",
		SameSite: http.SameSiteNoneMode, // Allows cross-site cookies
	}

	http.SetCookie(w, cookie)
}

func ValidateInput(data interface{}) error {
	validate := validator.New()
	err := validate.Struct(data)
	if err == nil {
		return nil
	}

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		errors := make(map[string]string)
		for _, fieldError := range validationErrors {
			errors[fieldError.Field()] = getErrorMessage(fieldError)
		}
		return &ValidationError{Errors: errors}
	}

	return err
}

func GetHttpRequestContextValue(r *http.Request, key any) (string, error) {
	value := r.Context().Value(key)
	if value == nil {
		return "", errors.New("value for key not found in context")
	}

	stringValue, ok := value.(string)
	if !ok {
		return "", errors.New("value is not a string")
	}

	return stringValue, nil
}

func getErrorMessage(fieldError validator.FieldError) string {
	switch fieldError.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return fmt.Sprintf("Value is too short, minimum is %s", fieldError.Param())
	case "max":
		return fmt.Sprintf("Value is too long, maximum is %s", fieldError.Param())
	case "gte":
		return fmt.Sprintf("Value is too small, minimum is %s", fieldError.Param())
	case "lte":
		return fmt.Sprintf("Value is too large, maximum is %s", fieldError.Param())
	default:
		return "Invalid value"
	}
}

func (s *Server) HandleError(w http.ResponseWriter, err error, msg string, statusCode int) {
	if err != nil {
		log.Println(err)

		responseError := s.ErrorJson(w, errors.New(msg), statusCode)
		if responseError != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
		}
	}
}

func ParseIDParam(r *http.Request, paramName string) (int, error) {
	idParam := chi.URLParam(r, paramName)
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		return 0, errors.New("invalid id parameter")
	}
	return id, nil
}

func (s *Server) GetUserIdFromRequest(r *http.Request) (string, error) {
	userId, err := GetHttpRequestContextValue(r, constants.UserIDKey)
	if err != nil {
		return "", errors.New("unauthorized")
	}
	return userId, nil
}
