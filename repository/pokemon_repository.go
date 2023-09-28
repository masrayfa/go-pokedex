package repository

import (
	"context"
	"database/sql"
	"ngacak-go/model/domain"
	"ngacak-go/model/web"
)

type PokemonRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Pokemon, error)
	FindCollections(ctx context.Context, tx *sql.Tx, userId int) ([]domain.Pokemon, error)
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Pokemon, error)
	FindByName(ctx context.Context, tx *sql.Tx, name string) (domain.Pokemon, error)
	FindAllStatsByPokemonId(ctx context.Context, tx *sql.Tx, pokemonId int64) ([]domain.Stats, error)
	FindAllTypesByPokemonId(ctx context.Context, tx *sql.Tx, pokemonId int64) ([]string, error)
	FindAllAbilitiesByPokemonId(ctx context.Context, tx *sql.Tx, pokemonId int64) ([]string, error)
	FindAllSpeciesByPokemonId(ctx context.Context, tx *sql.Tx, pokemonId int64) ([]string, error)
	Save(ctx context.Context, tx *sql.Tx, pokemon domain.Pokemon) (domain.Pokemon, error)
	Update(ctx context.Context, tx *sql.Tx, pokemon domain.Pokemon) (web.PokemonResponse, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) 
}
