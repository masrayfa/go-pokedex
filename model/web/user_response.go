package web

import "ngacak-go/model/domain"

type UserResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Pokemon  []domain.Pokemon `json:"pokemon"`
}