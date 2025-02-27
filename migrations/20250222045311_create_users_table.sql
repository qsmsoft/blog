-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS users CASCADE;

CREATE
EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE
EXTENSION IF NOT EXISTS CITEXT;

CREATE TABLE users
(
    id         UUID PRIMARY KEY            DEFAULT uuid_generate_v4(),
    first_name VARCHAR(32)        NOT NULL CHECK ( first_name <> '' ),
    last_name  VARCHAR(32)        NOT NULL CHECK ( last_name <> '' ),
    email      VARCHAR(64) UNIQUE NOT NULL CHECK ( email <> '' ),
    password   VARCHAR(250)       NOT NULL CHECK ( octet_length(password) <> 0 ),
    role       VARCHAR(10)        NOT NULL DEFAULT 'user',
    avatar     VARCHAR(512),
    created_at TIMESTAMP WITH TIME ZONE    DEFAULT (NOW() AT TIME ZONE 'UTC'),
    updated_at TIMESTAMP WITH TIME ZONE    DEFAULT (NOW() AT TIME ZONE 'UTC'),
    login_date TIMESTAMP WITH TIME ZONE    DEFAULT (NOW() AT TIME ZONE 'UTC')
);

CREATE
OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at
= NOW();
RETURN NEW;
END;
$$
LANGUAGE plpgsql;

CREATE TRIGGER update_users_updated_at
    BEFORE UPDATE
    ON users
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER IF EXISTS update_users_updated_at ON users;
DROP FUNCTION IF EXISTS update_updated_at_column();
DROP TABLE IF EXISTS users;

-- +goose StatementEnd
