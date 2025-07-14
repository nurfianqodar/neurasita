package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/nurfianqodar/neurasita/internal/dto"
	"github.com/nurfianqodar/neurasita/internal/service"
	"github.com/nurfianqodar/neurasita/pkg/errorw"
	"github.com/nurfianqodar/neurasita/pkg/response"
)

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService}
}

var _ Handler = (*UserHandler)(nil)

type UserHandler struct {
	userService *service.UserService
}

func (uh *UserHandler) RegisterRouter(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/users", Make(uh.handleCreateUser))
}

func (uh *UserHandler) handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	// Encode request body
	req := new(dto.CreateUserRequest)
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("[error] encode json failed: %v\n", err)
		return errorw.NewMalformedRequestBody()
	}

	// invoke user service with timeout
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*5)
	defer cancel()
	user, err := uh.userService.CreateUser(ctx, req)
	if err != nil {
		return err
	}

	// Send response success
	response.WriteJSON(w, response.NewJSON(true, http.StatusCreated, &dto.SingleUserData{
		User: user,
	}))
	return nil
}
