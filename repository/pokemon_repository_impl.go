package repository

import (
	"context"
	"database/sql"
	"errors"
	"ngacak-go/helper"
	"ngacak-go/model/domain"
	"ngacak-go/model/web"
)

type PokemonRepositoryImpl struct {
}


func NewPokemonRepository() PokemonRepository {
	return &PokemonRepositoryImpl{}
}

func (pokemonRepositoryImpl PokemonRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Pokemon, error) {
	script := `SELECT id, name, types, species, height, weight, abilities, hp, attack, defense, speed
	FROM pokemon 
	INNER JOIN about ON pokemon.id = about.id_pokemon 
	JOIN stats ON pokemon.id = stats.id_pokemon`

	rows, err := tx.QueryContext(ctx, script)
	helper.PanicIfError(err)
	defer rows.Close()

	var pokemons []domain.Pokemon
	for rows.Next() {
		var pokemon domain.Pokemon
		err := rows.Scan(&pokemon.Id, &pokemon.Name, &pokemon.Types, &pokemon.About.Species, &pokemon.About.Height, &pokemon.About.Weight, &pokemon.About.Abilities, &pokemon.Stats.HP, &pokemon.Stats.Attack, &pokemon.Stats.Defense, &pokemon.Stats.Speed)
		helper.PanicIfError(err)
		pokemons = append(pokemons, pokemon)
	}

	return pokemons, nil
}

func (pokemonRepositoryImpl PokemonRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Pokemon, error) {
	script := "SELECT id, name, types, height, weight, attack from pokemon inner join about on pokemon.id = about.id_pokemon join stats on pokemon.id = stats.id_pokemon where id = ?"
	rows, err := tx.QueryContext(ctx, script, id)
	helper.PanicIfError(err)
	defer rows.Close()

	pokemon := domain.Pokemon{}
	if rows.Next() {
		err := rows.Scan(&pokemon.Id, &pokemon.Name, &pokemon.Types, &pokemon.About.Height, &pokemon.About.Weight, &pokemon.Stats.Attack)
		helper.PanicIfError(err)
		return pokemon, nil
	} else {
		return pokemon, errors.New("pokemon not found")
	}
}

func (pokemonRepositoryImpl PokemonRepositoryImpl) FindByName(ctx context.Context, tx *sql.Tx, name string) (domain.Pokemon, error) {
	script := "SELECT id, name, types, height, weight, attack from pokemon inner join about on pokemon.id = about.id_pokemon join stats on pokemon.id = stats.id_pokemon where name = ?"
	row := tx.QueryRowContext(ctx, script, name)

	var pokemon domain.Pokemon
	err := row.Scan(&pokemon.Id, &pokemon.Name, &pokemon.Types, &pokemon.About.Height, &pokemon.About.Weight, &pokemon.Stats.Attack)
	helper.PanicIfError(err)

	return pokemon, nil
}

func (pokemonRepositoryImpl PokemonRepositoryImpl) FindAllStatsByPokemonId(ctx context.Context, tx *sql.Tx, pokemonId int64) ([]domain.Stats, error) {
	script := "SELECT id_stats, hp, attack, defense, speed from stats where id_pokemon = ?"
	rows, err := tx.QueryContext(ctx, script, pokemonId)
	helper.PanicIfError(err)
	defer rows.Close()

	var stats []domain.Stats
	for rows.Next() {
		var stat domain.Stats
		err := rows.Scan(&stat.Id_stats, &stat.HP, &stat.Attack, &stat.Defense, &stat.Speed)
		helper.PanicIfError(err)
		stats = append(stats, stat)
	}

	return stats, nil
}

func (pokemonRepositoryImpl PokemonRepositoryImpl) FindAllTypesByPokemonId(ctx context.Context, tx *sql.Tx, pokemonId int64) ([]string, error) {
	script := "SELECT type from pokemon where id_pokemon = ?"
	rows, err := tx.QueryContext(ctx, script, pokemonId)
	helper.PanicIfError(err)
	defer rows.Close()

	var types []string
	for rows.Next() {
		var t string
		err := rows.Scan(&t)
		helper.PanicIfError(err)
		types = append(types, t)
	}

	return types, nil
}

func (pokemonRepositoryImpl PokemonRepositoryImpl) FindAllAbilitiesByPokemonId(ctx context.Context, tx *sql.Tx, pokemonId int64) ([]string, error) {
	script := "SELECT ability from about where id_pokemon = ?"
	rows, err := tx.QueryContext(ctx, script, pokemonId)
	helper.PanicIfError(err)
	defer rows.Close()

	var abilities []string
	for rows.Next() {
		var ability string
		err := rows.Scan(&ability)
		helper.PanicIfError(err)
		abilities = append(abilities, ability)
	}

	return abilities, nil
}

func (pokemonRepositoryImpl PokemonRepositoryImpl) FindAllSpeciesByPokemonId(ctx context.Context, tx *sql.Tx, pokemonId int64) ([]string, error) {
	script := "SELECT species from about where id_pokemon = ?"
	rows, err := tx.QueryContext(ctx, script, pokemonId)
	helper.PanicIfError(err)
	defer rows.Close()

	var species []string
	for rows.Next() {
		var s string
		err := rows.Scan(&s)
		helper.PanicIfError(err)
		species = append(species, s)
	}

	return species, nil
}

func (pokemonRepositoryImpl PokemonRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, pokemon domain.Pokemon, userId int) (domain.Pokemon, error) {
	scriptPokemon := "INSERT INTO pokemon(name, types, id_user) VALUES (?, ?, ?)"
	resultPokemonTable, err := tx.ExecContext(ctx, scriptPokemon, pokemon.Name, pokemon.Types, userId)
	helper.PanicIfError(err)

	lastInsertId, err := resultPokemonTable.LastInsertId()
	helper.PanicIfError(err)

	scriptStats := "INSERT INTO stats (hp, attack, defense, speed, id_pokemon) VALUES (?, ?, ?, ?, ?)"
	_, err = tx.ExecContext(ctx, scriptStats, pokemon.Stats.HP, pokemon.Stats.Attack, pokemon.Stats.Defense, pokemon.Stats.Speed, lastInsertId)
	helper.PanicIfError(err)

	scriptAbout := "INSERT INTO about (species, height, weight, abilities, id_pokemon) VALUES (?, ?, ?, ?, ?)"
	_, err = tx.ExecContext(ctx, scriptAbout, pokemon.About.Species, pokemon.About.Height, pokemon.About.Weight, pokemon.About.Abilities, lastInsertId)
	helper.PanicIfError(err)


	pokemon.Id = int(lastInsertId)
	return pokemon, nil

}

func (pokemonRepositoryImpl PokemonRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) {
	script := "DELETE FROM pokemon WHERE id = ?"
	_, err := tx.ExecContext(ctx, script, id)
	helper.PanicIfError(err)
}

// Update implements PokemonRepository.
func (*PokemonRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, pokemon domain.Pokemon) (web.PokemonResponse, error) {
	script := "UPDATE pokemon SET name = ?, types = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, script, pokemon.Name, pokemon.Types, pokemon.Id)
	helper.PanicIfError(err)

	script = "UPDATE about SET species = ?, height = ?, weight = ?, abilities = ? WHERE id_pokemon = ?"
	_, err = tx.ExecContext(ctx, script, pokemon.About.Species, pokemon.About.Height, pokemon.About.Weight, pokemon.About.Abilities, pokemon.Id)
	helper.PanicIfError(err)

	script = "UPDATE stats SET hp = ?, attack = ?, defense = ?, speed = ? WHERE id_pokemon = ?"
	_, err = tx.ExecContext(ctx, script, pokemon.Stats.HP, pokemon.Stats.Attack, pokemon.Stats.Defense, pokemon.Stats.Speed, pokemon.Id)
	helper.PanicIfError(err)

	return helper.ToPokemonResponse(pokemon), nil
}