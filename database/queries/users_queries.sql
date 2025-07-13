-- name: CreateUser :one
INSERT INTO users
    (id, email, hash_password)
VALUES
    ($1, $2, $3)
RETURNING id, created_at;
