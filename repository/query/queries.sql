-- name: GetBook :one
SELECT * FROM books
WHERE id = $1 LIMIT 1;

-- name: ListBooks :many
SELECT * FROM books
ORDER BY name;

-- name: CreateBook :one
INSERT INTO books (
  name, author
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateBook :exec
UPDATE books
  set name = $2,
  author = $3
WHERE id = $1;

-- name: DeleteBook :exec
DELETE FROM books
WHERE id = $1;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: CheckUser :one
SELECT * FROM users
WHERE username = $1 and password = $2 LIMIT 1;

-- name: ListUser :many
SELECT * FROM users
ORDER BY username;

-- name: CreateUser :one
INSERT INTO users (
  username,password, role
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

---

-- name: BringBorrowedByUserName :many
SELECT ub.id,u.username,b.name,ub.date,ub.status FROM user_books ub 
INNER JOIN users u ON ub.user_id=u.id
INNER JOIN books b ON ub.book_id=b.id
WHERE u.username = $1;

-- name: BringBorrowedByBookId :many
SELECT ub.id,u.username,b.name,ub.date,ub.status FROM user_books ub
INNER JOIN users u ON ub.user_id=u.id
INNER JOIN books b ON ub.book_id=b.id
WHERE ub.book_id = $1 ;

-- name: ListBorrow :many
SELECT ub.id,u.username,b.name,ub.date,ub.status FROM user_books ub
INNER JOIN users u ON ub.user_id=u.id
INNER JOIN books b ON ub.book_id=b.id;

-- name: BringActiveBorrowByUser :one
SELECT * FROM user_books ub
WHERE ub.user_id=$1 and ub.status=$2 LIMIT 1;

-- name: BringBorrowedById :one
SELECT * FROM user_books ub
WHERE ub.id=$1 LIMIT 1;

-- name: CreateBorrow :one
INSERT INTO user_books (
  user_id, book_id, date, status
) 
SELECT 
  $1, $2, NOW(), $3
WHERE EXISTS (
  SELECT 1 FROM books WHERE id = $2
) AND EXISTS (
  SELECT 1 FROM users WHERE id = $1
)
RETURNING *;

-- name: UpdateStatusBorrow :exec
UPDATE user_books
set status = false
WHERE id = $1;

-- name: DeleteBorrow :exec
DELETE FROM user_books
WHERE id = $1;

-- name: DeleteBorrowByUserId :exec
DELETE FROM user_books
WHERE user_id = $1;

-- name: DeleteBorrowByBookId :exec
DELETE FROM user_books
WHERE book_id = $1;

-- name: TooLateBringBorrowed :many
SELECT ub.id,u.username, b.name, ub.date,ub.status
FROM user_books ub
INNER JOIN users u ON ub.user_id=u.id
INNER JOIN books b ON ub.book_id=b.id
WHERE NOW() >= date + INTERVAL '7 days' and status = $1;
