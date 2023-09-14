package service

import (
	"context"
	"ngacak-go/model/web"
)

type UserService interface {
	FindAll(ctx context.Context) ([]web.UserResponse, error)
	FindById(ctx context.Context, id int) (web.UserResponse, error)
	FindByUsername(ctx context.Context, username string) (web.UserResponse, error)
	FindByUsernameAndPassword(ctx context.Context, username string, password string) (web.UserResponse, error)
	Create(ctx context.Context, req web.UserCreateRequest) (web.UserResponse, error)
	Update(ctx context.Context, req web.UserUpdateRequest) (web.UserResponse, error)
	Delete(ctx context.Context, id int)
}