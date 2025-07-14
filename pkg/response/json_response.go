package response

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Membuat json response standar untuk app ini
func NewJSON(success bool, statusCode int, data any) *JSONResponse {
	return &JSONResponse{
		StatusCode: statusCode,
		AccessedAt: time.Now(),
		Success:    success,
		Data:       data,
	}
}

// Bentuk json response standard pada aplikasi ini
type JSONResponse struct {
	StatusCode int `json:"statusCode"`
	Success    bool
	AccessedAt time.Time
	Data       any `json:"data,omitempty"`
}

// Mengirim *JSON response ke http.ResponseWriter
//
// *Note: pastikan data yang dikirim sudah benar dan hindari pemanggilan
// ganda pada fungsi ini!
func WriteJSON(w http.ResponseWriter, res *JSONResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.StatusCode)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("[error] unable to write response: %v\n", err)
	}
}
