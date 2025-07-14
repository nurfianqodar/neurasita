package errorw

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/nurfianqodar/neurasita/pkg/global"
	"github.com/nurfianqodar/neurasita/pkg/response"
)

func New(statusCode int, message string, detail any) *APIError {
	return &APIError{
		StatusCode: statusCode,
		Message:    message,
		Detail:     detail,
	}
}

type APIError struct {
	StatusCode int
	Message    string
	Detail     any
}

func (ae *APIError) Error() string {
	return ae.Message
}

func (ae *APIError) Response() *response.JSONResponse {
	return response.NewJSON(false, ae.StatusCode, &apiErrorData{
		Error: &apiErrorDataContent{
			Message: ae.Message,
			Detail:  ae.Detail,
		},
	})
}

// Type helper
type apiErrorData struct {
	Error *apiErrorDataContent `json:"error"`
}

// Type helper
type apiErrorDataContent struct {
	Message string `json:"message"`
	Detail  any    `json:"detail,omitempty"`
}

func NewInternalServerError() *APIError {
	return New(http.StatusInternalServerError, "internal server error", nil)
}

func NewMalformedRequestBody() *APIError {
	return New(http.StatusBadRequest, "malformed request body", nil)
}

func NewValidationError(vErr validator.ValidationErrors) *APIError {
	detail := make([]map[string]string, len(vErr))

	for i, fErr := range vErr {
		detail[i] = map[string]string{
			"field":     fErr.Field(),
			"violation": fErr.Translate(global.Translator),
		}
	}

	return New(http.StatusBadRequest, "validation error", detail)
}
