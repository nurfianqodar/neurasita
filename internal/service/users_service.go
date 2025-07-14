// User service package

package service

import (
	"context"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/nurfianqodar/neurasita/internal/dto"
	"github.com/nurfianqodar/neurasita/internal/repository"
	"github.com/nurfianqodar/neurasita/pkg/errorw"
	"github.com/nurfianqodar/neurasita/pkg/global"
	"golang.org/x/crypto/bcrypt"
)

// Constructor untuk *UserService
func NewUserService(q *repository.Queries) *UserService {
	return &UserService{q}
}

// Struct untuk melakukan validasi, invoke repository dan semua proses dari handling
// request pada user management endpoints.
type UserService struct {
	q *repository.Queries
}

// CreateUser membuat satu user (signup) mengembalikan user id dan createdat jika berhasil
func (us *UserService) CreateUser(ctx context.Context, req *dto.CreateUserRequest) (*dto.CreateUserReponse, error) {
	// Validate request body
	if err := global.Validate.Struct(req); err != nil {
		if vErr, ok := err.(validator.ValidationErrors); ok {
			return nil, errorw.NewValidationError(vErr)
		} else {
			log.Printf("[error] unexpected behavior on valiator: %v\n", err)
			return nil, err
		}
	}

	// Create new uuid v7
	uid, err := uuid.NewV7()
	if err != nil {
		log.Printf("[error] failed to create uuid: %v\n", err)
		return nil, err
	}

	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("[error] hash password failed: %v\n", err)
		return nil, err
	}

	// Insert user to database
	result, err := us.q.CreateUser(ctx, &repository.CreateUserParams{
		ID:           uid,
		Email:        req.Email,
		HashPassword: string(hash),
	})
	if err != nil {
		log.Printf("[error] insert user error: %v\n", err)
		return nil, err
	}

	return &dto.CreateUserReponse{
		ID:        result.ID,
		CreatedAt: result.CreatedAt.Time,
	}, nil
}
