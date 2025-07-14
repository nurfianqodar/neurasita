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

// APIError merupakan standar json error pada aplikasi ini
// pastikan konversi error ke tipe ini jika error merupakan
// kesalahan client seperti error validasi dan pelanggaran unique
// constraint database.
type APIError struct {
	StatusCode int
	Message    string
	Detail     any
}

func (ae *APIError) Error() string {
	return ae.Message
}

// Response method mengembalikan JSONResponse yang siap dikirim
// melalui fungsi WriteJSON.
func (ae *APIError) Response() *response.JSONResponse {
	return response.NewJSON(false, ae.StatusCode, &apiErrorData{
		Error: &apiErrorDataContent{
			Message: ae.Message,
			Detail:  ae.Detail,
		},
	})
}

// Type helper untuk standarisasi error response (private)
type apiErrorData struct {
	Error *apiErrorDataContent `json:"error"`
}

// Type helper untuk standarisasi error response (private)
type apiErrorDataContent struct {
	Message string `json:"message"`
	Detail  any    `json:"detail,omitempty"`
}

// Mengubah validator.ValidationErrors menjadi APIError
func FromValidationError(vErr validator.ValidationErrors) *APIError {
	detail := make([]map[string]string, len(vErr))

	for i, fErr := range vErr {
		detail[i] = map[string]string{
			"field":     fErr.Field(),
			"violation": fErr.Translate(global.Translator),
		}
	}

	return New(http.StatusBadRequest, "validation error", detail)
}
