-- ==================================================================== --
-- Users Table
-- ==================================================================== --
-- users table berisi data kredensial user digunakan untuk autentikasi
-- pada aplikasi.

CREATE TABLE IF NOT EXISTS users (
    -- PK
    id uuid PRIMARY KEY,
    -- Main Data
    email varchar(255) NOT NULL,
    hash_password varchar(255) NOT NULL,
    email_active boolean NOT NULL DEFAULT FALSE,
    -- Timestamp
    created_at timestamptz DEFAULT current_timestamp,
    updated_at timestamptz DEFAULT current_timestamp,
    deleted_at timestamptz
);

CREATE UNIQUE INDEX IF NOT EXISTS uq_users_email_deleted_at
ON users (email)
WHERE deleted_at IS NULL;

CREATE INDEX IF NOT EXISTS idx_users_email
ON users (email);

CREATE INDEX IF NOT EXISTS idx_users_deleted_at
ON users (deleted_at);
