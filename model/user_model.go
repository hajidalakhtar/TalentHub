package model

type UserResponse struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
type LoginResponse struct {
	Token    string `json:"token"`
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
