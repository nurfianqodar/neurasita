package handler

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/nurfianqodar/neurasita/pkg/errorw"
	"github.com/nurfianqodar/neurasita/pkg/response"
)

// Internal handler adalah http.HandlerFunc yang mengembalikan error
// tipe ini di wrap oleh Make func untuk mengubahnya menjadi http.HandlerFunc
type InternalHandler func(w http.ResponseWriter, r *http.Request) error

// Make func mengubah InternalHandler yang mengembalikan error menjadi HandleFunc
// yang kompatibel dengan http.ServeMux. Fungsi ini juga menangani error yang
// dikembalikan oleh InternalHandler (jika operasi tidak berhasil) mengubahnya
// menjadi response yang sesuai
func Make(h InternalHandler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		if err != nil {
			// Handle api error (langsung write saja)
			if apiErr, ok := err.(*errorw.APIError); ok {
				response.WriteJSON(w, apiErr.Response())
				return
			} else
			// Handle postgres error (error yang dikembalikan dari operasi database)
			if pgErr, ok := err.(*pgconn.PgError); ok {
				// email sudah digunakan
				if pgErr.ConstraintName == "uqidx_users_email_deleted_at" {
					response.WriteJSON(w, errorw.New(http.StatusConflict, "email not avaliable", nil).Response())
					return
				}
			}
			// Default error
			response.WriteJSON(w, errorw.NewInternalServerError().Response())
		}
	})
}
