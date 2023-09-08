package repository

import (
	"context"
	"database/sql"
	"errors"
	"ngacak-go/helper"
	"ngacak-go/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepositoryImpl() UserRepositoryImpl {
	return UserRepositoryImpl{}
}

func (userRepositoryImpl UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.User, error) {
	script := "SELECT id, name from user"
	rows, err := tx.QueryContext(ctx, script)
	helper.PanicIfError(err)
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.Id, &user.Username)
		helper.PanicIfError(err)
		users = append(users, user)
	}

	return users, nil
}

func (userRepositoryImpl UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int64) (domain.User, error) {
	script := "SELECT id, name from user where id = ?"
	rows, err := tx.QueryContext(ctx, script, id)
	helper.PanicIfError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}

func (userRepositoryImpl UserRepositoryImpl) FindByUsername(ctx context.Context, tx *sql.Tx, username string) (domain.User, error) {
	script := "SELECT id, name from user where username = ?"
	row := tx.QueryRowContext(ctx, script, username)

	var user domain.User
	err := row.Scan(&user.Id, &user.Username)
	helper.PanicIfError(err)

	return user, nil
}

func (userRepositoryImpl UserRepositoryImpl) FindByUsernameAndPassword(ctx context.Context, tx *sql.DB, username string, password string) (domain.User, error) {
	script := "SELECT id, name from user where username = ? and password = ?"
	row := tx.QueryRowContext(ctx, script, username, password)

	var user domain.User
	err := row.Scan(&user.Id, &user.Username)
	helper.PanicIfError(err)

	return user, nil
}

func (userRepositoryImpl UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	script := "INSERT INTO user (username, password) VALUES (? ?)"
	result, err := tx.ExecContext(ctx, script, user.Username, user.Password)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	user.Id = id

	return user, nil
}

func (userRepositoryImpl UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	script := "UPDATE user SET name = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, script, user.Username, user.Id)
	helper.PanicIfError(err)

	return user, nil
}

func (userRepositoryImpl UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int64) error {
	script := "DELETE FROM user WHERE id = ?"
	_, err := tx.ExecContext(ctx, script, id)
	helper.PanicIfError(err)

	return nil
}