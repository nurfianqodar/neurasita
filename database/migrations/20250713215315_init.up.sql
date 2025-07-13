CREATE TABLE users (
    -- PK
    id uuid PRIMARY KEY,
    -- Main Data
    email varchar(255) NOT NULL,
    hash_password varchar(255) NOT NULL,
    -- Timestamp
    created_at timestamptz DEFAULT current_timestamp,
    updated_at timestamptz DEFAULT current_timestamp,
    deleted_at timestamptz
);

CREATE UNIQUE INDEX uq_users_email_deleted_at
    ON users(email)
    WHERE deleted_at IS NULL;

CREATE INDEX idx_users_email
    ON users(email);

CREATE INDEX idx_users_deleted_at
    ON users(deleted_at);
