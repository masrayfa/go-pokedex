package web

import "ngacak-go/model/domain"

type PokemonCreateRequest struct {
	Name   string       `json:"name"`
	Types  string       `json:"types"`
	About  domain.About `json:"about"`
	Stats  domain.Stats `json:"stats"`
}