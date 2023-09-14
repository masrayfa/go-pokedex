package web

type UserUpdateRequest struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}