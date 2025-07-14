package errorw

import "net/http"

var (
	ErrInternalServer           = New(http.StatusInternalServerError, "internal server error", nil)
	ErrInvalidCredential        = New(http.StatusUnauthorized, "email atau password salah", nil)
	ErrMalformedRequestBody     = New(http.StatusBadRequest, "malformed request body", nil)
	ErrRequestTimeout           = New(http.StatusRequestTimeout, "request timeout", nil)
	ErrConflictUniqueConstraint = New(http.StatusConflict, "conflict data on unique constraint", nil)
	ErrInvalidTextRepr          = New(http.StatusUnprocessableEntity, "invalid text representation", nil)
	ErrInvalidUUID              = New(http.StatusBadRequest, "invalid uuid", nil)
)
