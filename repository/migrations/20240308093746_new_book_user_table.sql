-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_books (
id SERIAL PRIMARY KEY,
user_id INT NOT NULL,
book_id INT NOT NULL,
date DATE NOT NULL,
status VARCHAR(100) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_books;
-- +goose StatementEnd
