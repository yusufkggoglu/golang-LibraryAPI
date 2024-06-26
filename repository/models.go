// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package repository

import (
	"time"
)

type Book struct {
	ID     int32  `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

type User struct {
	ID       int32  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UserBook struct {
	ID     int32     `json:"id"`
	UserID int32     `json:"user_id"`
	BookID int32     `json:"book_id"`
	Date   time.Time `json:"date"`
	Status string    `json:"status"`
}
