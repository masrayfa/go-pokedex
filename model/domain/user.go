package domain

type User struct {
	Id       int64     `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Pokemon  []Pokemon `json:"pokemon"`
}