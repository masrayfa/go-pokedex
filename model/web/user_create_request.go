package web

type UserCreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}