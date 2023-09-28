package service

import (
	"context"
	"database/sql"
	"fmt"
	"ngacak-go/exception"
	"ngacak-go/helper"
	"ngacak-go/model/domain"
	"ngacak-go/model/web"
	"ngacak-go/repository"

	"github.com/go-playground/validator"
)

type PokemonServiceImpl struct {
	PokemonRepository repository.PokemonRepository
	UserRepository repository.UserRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewPokemonService(pokemonRepository repository.PokemonRepository, DB *sql.DB, Validate *validator.Validate) PokemonService {
	return &PokemonServiceImpl{
		PokemonRepository: pokemonRepository,
		DB:                DB,
		Validate:          Validate,
	}
}

// Create implements PokemonService.
func (service *PokemonServiceImpl) Create(ctx context.Context, req web.PokemonCreateRequest) (web.PokemonResponse, error) {
	err := service.Validate.Struct(req)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	pokemon := domain.Pokemon{
		Name: req.Name,
		About: domain.About{
			Height: req.About.Height,
			Weight: req.About.Weight,
			Species: req.About.Species,
			Abilities: req.About.Abilities,
		},

		Stats: domain.Stats{
			Attack: req.Stats.Attack,
			Defense: req.Stats.Defense,
			HP: req.Stats.HP,
			Speed: req.Stats.Speed,
		},
		Types: req.Types,
		UserId: req.UserId,
	}

	// _, err = service.UserRepository.FindById(ctx, tx, req.UserId)
	// if err != nil {
	// 	panic(exception.NewNotFoundError(err.Error()))
	// }

	// pokemon := web.PokemonResponseToDomain(req)
	res, err := service.PokemonRepository.Save(ctx, tx, pokemon) 
	helper.PanicIfError(err)

	return helper.ToPokemonResponse(res), nil
}

// FindById implements PokemonService.
func (service *PokemonServiceImpl) FindById(ctx context.Context, pokemonId int) (web.PokemonResponse, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	pokemon, err := service.PokemonRepository.FindById(ctx, tx, pokemonId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	
	return helper.ToPokemonResponse(pokemon), nil
}

// FindByName implements PokemonService.
func (service *PokemonServiceImpl) FindByName(ctx context.Context, name string) (web.PokemonResponse, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	pokemon, err := service.PokemonRepository.FindByName(ctx, tx, name)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToPokemonResponse(pokemon), nil
}

// Update implements PokemonService.
func (service *PokemonServiceImpl) Update(ctx context.Context, req web.PokemonUpdateRequest) (web.PokemonResponse, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	pokemon, err := service.PokemonRepository.FindById(ctx, tx, req.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	pokemon.Name = req.Name

	service.PokemonRepository.Update(ctx, tx, pokemon)

	return helper.ToPokemonResponse(pokemon), nil
}


func (service *PokemonServiceImpl) FindAll(ctx context.Context) ([]web.PokemonResponse, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	fmt.Println("ini dari service")

	var pokemonResponses[]web.PokemonResponse
	pokemons, err := service.PokemonRepository.FindAll(ctx, tx)
	helper.PanicIfError(err)

	for _, pokemon := range pokemons {
		pokemonResponses = append(pokemonResponses, helper.ToPokemonResponse(pokemon))
	}

	return pokemonResponses, nil
}

// FindCollections implements PokemonService.
func (service *PokemonServiceImpl) FindCollections(ctx context.Context, userId int) ([]web.PokemonResponse, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	fmt.Println("ini dari service", userId)

	// _, err = service.UserRepository.FindById(ctx, tx, userId)
	// helper.PanicIfError(err)

	var pokemonResponses[]web.PokemonResponse
	pokemons, err := service.PokemonRepository.FindCollections(ctx, tx, userId)
	helper.PanicIfError(err)

	for _, pokemon := range pokemons {
		pokemonResponses = append(pokemonResponses, helper.ToPokemonResponse(pokemon))
	}

	return pokemonResponses, nil
}

func (service *PokemonServiceImpl) Delete(ctx context.Context, id int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	pokemon, err := service.PokemonRepository.FindById(ctx, tx, id)
	if err != nil {
		exception.NewNotFoundError(err.Error())
	}

	service.PokemonRepository.Delete(ctx, tx, pokemon.Id)
}