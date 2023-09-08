package web

import "ngacak-go/model/domain"

type PokemonUpdateRequest struct {
	Id     int         `json:"id"`
	Name   string      `json:"name"`
	Types  []string    `json:"types"`
	Weight int         `json:"weight"`
	Height int         `json:"height"`
	Stats  domain.Stats `json:"stats"`
}
