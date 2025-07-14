package handler

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/nurfianqodar/neurasita/internal/repository"
	"github.com/nurfianqodar/neurasita/pkg/errorw"
	"github.com/nurfianqodar/neurasita/pkg/response"
)

// Internal handler adalah http.HandlerFunc yang mengembalikan error
// tipe ini di wrap oleh Make func untuk mengubahnya menjadi http.HandlerFunc
//
// *Note: jangan panggil w.WriteHeader(...) jika InternalHandler mengembalikan error
// pastikan panggil w.WriteHeader hanya jika semua operasi sukses
// karena error akan dihandle oleh Make function
type InternalHandler func(w http.ResponseWriter, r *http.Request) error

// Make func mengubah InternalHandler yang mengembalikan error menjadi HandleFunc
// yang kompatibel dengan http.ServeMux. Fungsi ini juga menangani error yang
// dikembalikan oleh InternalHandler (jika operasi tidak berhasil) mengubahnya
// menjadi response yang sesuai
func Make(h InternalHandler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)

		// Berhenti disini jika tidak ada error
		if err == nil {
			return
		}

		// Error dihandle disini untuk menghindari nested if

		// Error yang bisa dikonversi
		// (diperlukan untuk switch case dengan errors.As)
		var (
			apiErr *errorw.APIError
			pgErr  *pgconn.PgError
		)

		log.Printf("[error] %v\n", err)

		switch {

		// APIError langsung write
		case errors.As(err, &apiErr):
			response.WriteJSON(w, apiErr.Response())

		// PgError dihandle handlePgError
		case errors.As(err, &pgErr):
			handlePgError(w, pgErr)

		// Context timeout error
		case errors.Is(err, context.DeadlineExceeded):
			response.WriteJSON(w, errorw.ErrRequestTimeout.Response())

		// Error lainnya dikonversi ke internal server error
		default:
			response.WriteJSON(w, errorw.ErrInternalServer.Response())
		}
	})
}

// menagani seluruh error yang dihasilkan dari operasi repository
func handlePgError(w http.ResponseWriter, pgErr *pgconn.PgError) {
	// Unique index violation pada email
	switch pgErr.Code {
	// Unique constraint violation
	case "23505":
		switch pgErr.ConstraintName {
		case repository.Cst_UqidxUsersEmailDeletedAt:
			res := errorw.New(http.StatusConflict, "email not avaliable", nil).Response()
			response.WriteJSON(w, res)
		default:
			response.WriteJSON(w, errorw.ErrConflictUniqueConstraint.Response())
		}
	case "22P02":
		response.WriteJSON(w, errorw.ErrInvalidTextRepr.Response())
	default:
		response.WriteJSON(w, errorw.ErrInternalServer.Response())
	}
}
