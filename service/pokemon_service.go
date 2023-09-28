package service

import (
	"context"
	"ngacak-go/model/web"
)

type PokemonService interface {
	FindAll(ctx context.Context) ([]web.PokemonResponse, error)
	FindCollections(ctx context.Context, userId int) ([]web.PokemonResponse, error)
	FindById(ctx context.Context, id int) (web.PokemonResponse, error)
	FindByName(ctx context.Context, name string) (web.PokemonResponse, error)
	Create(ctx context.Context, req web.PokemonCreateRequest) (web.PokemonResponse, error)
	Update(ctx context.Context, req web.PokemonUpdateRequest) (web.PokemonResponse, error) 
	Delete(ctx context.Context, id int)
}