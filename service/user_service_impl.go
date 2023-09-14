package service

import (
	"context"
	"database/sql"
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
func (*UserServiceImpl) Create(ctx context.Context, req web.UserCreateRequest) (web.UserResponse, error) {
	panic("unimplemented")
}

// Delete implements UserService.
func (*UserServiceImpl) Delete(ctx context.Context, id int) {
	panic("unimplemented")
}

// FindAll implements UserService.
func (*UserServiceImpl) FindAll(ctx context.Context) ([]web.UserResponse, error) {
	panic("unimplemented")
}

// FindById implements UserService.
func (*UserServiceImpl) FindById(ctx context.Context, id int) (web.UserResponse, error) {
	panic("unimplemented")
}

// FindByUsername implements UserService.
func (*UserServiceImpl) FindByUsername(ctx context.Context, username string) (web.UserResponse, error) {
	panic("unimplemented")
}

// FindByUsernameAndPassword implements UserService.
func (*UserServiceImpl) FindByUsernameAndPassword(ctx context.Context, username string, password string) (web.UserResponse, error) {
	panic("unimplemented")
}

// Update implements UserService.
func (*UserServiceImpl) Update(ctx context.Context, req web.UserUpdateRequest) (web.UserResponse, error) {
	panic("unimplemented")
}

