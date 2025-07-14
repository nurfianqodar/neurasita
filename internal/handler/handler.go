package handler

import "net/http"

// Handler adalah Interface bantuan untuk memastikan semua handler
// memiliki method RegisterRouter agar routing bisa didaftarkan ke
// mux (*http.ServeMux)
type Handler interface {
	// Mendaftarkan semua routing dari handler ke *http.ServeMux
	RegisterRouter(*http.ServeMux)
}
