package handler

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/nurfianqodar/neurasita/pkg/errorw"
	"github.com/nurfianqodar/neurasita/pkg/response"
)

type InternalHandler func(w http.ResponseWriter, r *http.Request) error

func Make(h InternalHandler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		if err != nil {
			if apiErr, ok := err.(*errorw.APIError); ok {
				// Handle api error
				response.WriteJSON(w, apiErr.Response())
				return
			} else if pgErr, ok := err.(*pgconn.PgError); ok {
				// Handle postgres error
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
