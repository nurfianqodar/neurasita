package dto

import (
	"time"

	"github.com/google/uuid"
)

// SingleUserData merupakan wrapper response agar response
// distandarisasi, gunakan ini jika telah melakukan operasi user
// yang mengembalikan satu entitas user, dignuakan pada
// func response.NewJSON pada parameter data
// Penggunaan:
//
//	response.NewJSON(true, 201, &dto.SingleUserData{
//		User: &CreateUserReponse{ ... } // atau response lainnya yang bukan array/slice
//	})
//
// Selalu gunakan wrapper untuk standarisasi!
type SingleUserData struct {
	User any `json:"user"`
}

// ManyUserData merupakan wrapper response agar response
// distandarisasi, gunakan ini jika telah melakukan operasi user
// yang mengembalikan banyak entitas user, dignuakan pada
// func response.NewJSON pada parameter data.
//
// Penggunaan:
//
//	response.NewJSON(true, 201, &dto.SingleUserData{
//		User: []&CreateUserReponse{ ... } // atau response array/slice
//	})
//
// Selalu gunakan wrapper untuk standarisasi!
type ManyUserData struct {
	Users any `json:"users"`
}

// CreateUserRequest merupakan bentuk request body yang dikirim user
// untuk membuat satu user
type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// CreateUserReponse merupakan bentuk data response jika sign up (create one user)
// berhasil. CreateUserReponse akan di wrap oleh struct SingleUserData untuk standarisasi
// bentuk response body yang dikirim ke user.
//
// *Note: Pastikan wrap dengan SingleUserData sebelum dimasukkan ke parameter response.NewJSON
type CreateUserReponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}
