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

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}
// Create implements UserService.
func (service *UserServiceImpl) Create(ctx context.Context, req web.UserCreateRequest) (web.UserResponse, error) {
	// 1. Validate request
	err := service.Validate.Struct(req)
	helper.PanicIfError(err)

	// 2. start transaction
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// 3. save to db
	user := domain.User{
		Username: req.Username,
		Password: req.Password,
	}
	res, err := service.UserRepository.Save(ctx, tx, user)
	helper.PanicIfError(err)

	// 4. return response
	return web.UserResponse{
		Id:       res.Id,
		Username: res.Username,
	}, nil
}

// Delete implements UserService.
func (service *UserServiceImpl) Delete(ctx context.Context, id int) {
	panic("unimplemented")
}

// FindAll implements UserService.
func (service *UserServiceImpl) FindAll(ctx context.Context) ([]web.UserResponse, error) {
	panic("unimplemented")
}

// FindById implements UserService.
func (service *UserServiceImpl) FindById(ctx context.Context, id int) (web.UserResponse, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return web.UserResponse{
		Id: 	 user.Id,
		Username: user.Username,
	}, nil
}

// FindByUsername implements UserService.
func (service *UserServiceImpl) FindByUsername(ctx context.Context, username string) (web.UserResponse, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	fmt.Println("dari service: ",username)

	user, err := service.UserRepository.FindByUsername(ctx, tx, username)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return web.UserResponse{
		Id:       user.Id,
		Username: user.Username,
	}, nil
}

// FindByUsernameAndPassword implements UserService.
func (service *UserServiceImpl) FindByUsernameAndPassword(ctx context.Context, username string, password string) (web.UserResponse, error) {
	panic("unimplemented")
}

// Update implements UserService.
func (service *UserServiceImpl) Update(ctx context.Context, req web.UserUpdateRequest) (web.UserResponse, error) {
	panic("unimplemented")
}