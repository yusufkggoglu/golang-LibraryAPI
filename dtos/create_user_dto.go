package dtos

type CreateUserDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
