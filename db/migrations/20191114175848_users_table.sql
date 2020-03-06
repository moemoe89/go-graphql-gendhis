
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE users (
    id VARCHAR(20) NOT NULL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    name VARCHAR(50) NOT NULL DEFAULT '',
    email VARCHAR(50) NOT NULL DEFAULT '',
    phone VARCHAR(20) NOT NULL DEFAULT '',
    address text NOT NULL DEFAULT ''
);
CREATE INDEX idx_deleted_at ON users (deleted_at);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF exists users;
DROP INDEX idx_deleted_at;