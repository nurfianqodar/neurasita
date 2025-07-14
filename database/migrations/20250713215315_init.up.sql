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

CREATE UNIQUE INDEX IF NOT EXISTS uqidx_users_email_deleted_at
ON users (email)
WHERE deleted_at IS NULL;

CREATE INDEX IF NOT EXISTS idx_users_email
ON users (email);

CREATE INDEX IF NOT EXISTS idx_users_deleted_at
ON users (deleted_at);

-- ==================================================================== --
-- Roles Table
-- ==================================================================== --
CREATE TABLE IF NOT EXISTS roles (
    -- PK
    id uuid PRIMARY KEY,
    -- Main Data
    name varchar(32) NOT NULL,
    -- Timestamp
    created_at timestamptz DEFAULT current_timestamp,
    updated_at timestamptz DEFAULT current_timestamp,

    CONSTRAINT uq_roles_name
    UNIQUE (name)
);


-- many to many user roles relation
CREATE TABLE IF NOT EXISTS user_roles (
    user_id uuid NOT NULL,
    role_id uuid NOT NULL,
    -- Timestamp
    created_at timestamptz DEFAULT current_timestamp,

    CONSTRAINT pk_user_roles
    PRIMARY KEY (user_id, role_id),

    CONSTRAINT fk_user_roles_user_id
    FOREIGN KEY (user_id) REFERENCES users (id)
    ON DELETE CASCADE,

    CONSTRAINT fk_user_roles_role_id
    FOREIGN KEY (role_id) REFERENCES roles (id)
    ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_user_roles_user_id
ON user_roles (user_id);

CREATE INDEX IF NOT EXISTS idx_user_roles_role_id
ON user_roles (role_id);

-- ==================================================================== --
-- Permissions Table
-- ==================================================================== --
CREATE TABLE IF NOT EXISTS permissions (
    -- PK
    id uuid PRIMARY KEY,
    -- Main Data
    permission varchar(64) NOT NULL,
    -- Timestamp
    created_at timestamptz DEFAULT current_timestamp,
    updated_at timestamptz DEFAULT current_timestamp,

    CONSTRAINT uq_permissions_permission
    UNIQUE (permission)
);


-- many to many user permissions relation
CREATE TABLE IF NOT EXISTS user_permissions (
    -- Main Data
    user_id uuid NOT NULL,
    permission_id uuid NOT NULL,
    -- Timestamp
    created_at timestamptz DEFAULT current_timestamp,

    CONSTRAINT pk_user_permissions
    PRIMARY KEY (user_id, permission_id),

    CONSTRAINT fk_user_permissions_user_id
    FOREIGN KEY (user_id) REFERENCES users (id)
    ON DELETE CASCADE,

    CONSTRAINT fk_user_permissions_permission_id
    FOREIGN KEY (permission_id) REFERENCES permissions (id)
    ON DELETE CASCADE
);

CREATE INDEX idx_user_permissions_user_id
ON user_permissions (user_id);

-- many to many role permissions relation
CREATE TABLE IF NOT EXISTS role_permissions (
    -- Main Data
    role_id uuid NOT NULL,
    permission_id uuid NOT NULL,
    -- Timestamp
    created_at timestamptz DEFAULT current_timestamp,

    CONSTRAINT pk_role_permissions
    PRIMARY KEY (role_id, permission_id),

    CONSTRAINT fk_role_permissions_role_id
    FOREIGN KEY (role_id) REFERENCES roles (id)
    ON DELETE CASCADE,

    CONSTRAINT fk_role_permissions_permission_id
    FOREIGN KEY (permission_id) REFERENCES permissions (id)
    ON DELETE CASCADE
);

CREATE INDEX idx_role_permissions_role_id
ON role_permissions (role_id);

-- ==================================================================== --
-- Profiles Table
-- ==================================================================== --
CREATE TYPE gender AS ENUM (
    'm',
    'f'
);

CREATE TABLE profiles (
    -- PK
    id uuid PRIMARY KEY,
    -- Main Data
    user_id uuid NOT NULL,
    display_name varchar(127),
    occupation varchar(63),
    organization varchar(127),
    gender gender,
    photo_url varchar(255),
    bio varchar(255),
    -- Timestamp
    created_at timestamptz DEFAULT current_timestamp,
    updated_at timestamptz DEFAULT current_timestamp,

    CONSTRAINT uq_profiles_user_id
    UNIQUE (user_id),

    CONSTRAINT fk_profiles_user_id
    FOREIGN KEY (user_id) REFERENCES users (id)
    ON DELETE CASCADE
);


-- ==================================================================== --
-- Files Table
-- ==================================================================== --
CREATE TABLE files (
    -- PK
    id uuid PRIMARY KEY,
    -- Main Data
    url varchar(255) NOT NULL,
    name varchar(255) NOT NULL,
    type varchar(64) NOT NULL,
    hash varchar(64) NOT NULL,
    -- Timestamp
    created_at timestamptz DEFAULT current_timestamp,
    updated_at timestamptz DEFAULT current_timestamp,
    deleted_at timestamptz,

    CONSTRAINT uq_files_hash
    UNIQUE (hash)
);

-- ==================================================================== --
-- Plants Table
-- ==================================================================== --

-- ==================================================================== --
-- Varietes Table
-- ==================================================================== --

-- ==================================================================== --
-- Plant Breedeng Records Table
-- ==================================================================== --

-- ==================================================================== --
-- Climate Records Table
-- ==================================================================== --

-- ==================================================================== --
-- Yield Records Table
-- ==================================================================== --

-- ==================================================================== --
-- Disease Records Table
-- ==================================================================== --
