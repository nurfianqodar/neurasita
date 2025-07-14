package handler

import "net/http"

type Handler interface {
	RegisterRouter(*http.ServeMux)
}
