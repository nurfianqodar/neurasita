-- name: CreateUser :one
INSERT INTO users
(id, email, hash_password)
VALUES
($1, $2, $3)
RETURNING id, created_at;


-- name: ListUser :many
SELECT
    id,
    email,
    created_at,
    updated_at
FROM users
LIMIT $1 OFFSET $2;


-- name: GetUserByID :one
SELECT
    id,
    email,
    email_active,
    created_at,
    updated_at
FROM users
WHERE id = $1 AND deleted_at IS NULL
LIMIT 1;


-- name: GetUserHashPassword :one
SELECT
    id,
    hash_password
FROM users
WHERE email = $1 AND deleted_at IS NULL
LIMIT 1;


-- name: UpdateUserPassword :one
UPDATE users
SET
    hash_password = $2,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1 AND deleted_at IS NULL
RETURNING id, updated_at;


-- name: UpdateUserEmail :one
UPDATE users
SET
    email = $2,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1 AND deleted_at IS NULL
RETURNING id, updated_at;


-- name: SoftDeleteUser :one
UPDATE users
SET
    deleted_at = CURRENT_TIMESTAMP,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1 AND deleted_at IS NULL
RETURNING id, deleted_at;


-- name: HardDeleteUser :one
DELETE FROM users
WHERE id = $1
RETURNING id;
