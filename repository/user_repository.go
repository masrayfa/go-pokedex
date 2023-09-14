package repository

import (
	"context"
	"database/sql"
	"ngacak-go/model/domain"
)

type UserRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.User, error)
	FindById(ctx context.Context, tx *sql.Tx, id int64) (domain.User, error)
	FindByUsername(ctx context.Context, tx *sql.Tx, username string) (domain.User, error)
	FindByUsernameAndPassword(ctx context.Context, tx *sql.DB, username string, password string) (domain.User, error)
	Save(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error)
	Update(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error)
	Delete(ctx context.Context, tx *sql.Tx, id int64) error
}
