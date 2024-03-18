-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS books (
id SERIAL PRIMARY KEY,
name VARCHAR(255) NOT NULL,
author VARCHAR(255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE books;
-- +goose StatementEnd
