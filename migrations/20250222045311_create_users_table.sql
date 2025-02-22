-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS users CASCADE;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS CITEXT;

CREATE TABLE users
(
    id          UUID PRIMARY KEY                  DEFAULT uuid_generate_v4(),
    first_name  VARCHAR(32)              NOT NULL CHECK ( first_name <> '' ),
    last_name   VARCHAR(32)              NOT NULL CHECK ( last_name <> '' ),
    email       VARCHAR(64) UNIQUE       NOT NULL CHECK ( email <> '' ),
    password    VARCHAR(250)             NOT NULL CHECK ( octet_length(password) <> 0 ),
    role        VARCHAR(10)              NOT NULL DEFAULT 'user',
    avatar      VARCHAR(512),
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_ate TIMESTAMP WITH TIME ZONE          DEFAULT CURRENT_TIMESTAMP,
    login_date  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
