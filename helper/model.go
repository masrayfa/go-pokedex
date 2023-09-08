package helper

import (
	"ngacak-go/model/domain"
	"ngacak-go/model/web"
)

func ToPokemonResponse(pokemon domain.Pokemon) web.PokemonResponse {
	return web.PokemonResponse{
		Id: int(pokemon.Id),
		Name: pokemon.Name,
		Height: pokemon.About.Height,
		Weight: pokemon.About.Weight,
		Stats: pokemon.Stats,
		Types: pokemon.Types,
		About: pokemon.About,
	}
}