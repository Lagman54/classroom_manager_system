package entity

import (
	"context"
	"database/sql"
	"time"
)

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UserModel struct {
	DB *sql.DB
}

func (u UserModel) Insert(user *User) error {
	query := `
			INSERT INTO users (first_name, last_name) 
			VALUES ($1, $2) 
			RETURNING id
`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	args := []any{user.FirstName, user.LastName}
	return u.DB.QueryRowContext(ctx, query, args...).Scan(&user.Id)
}

func (u UserModel) Get(id int) (*User, error) {
	query := `
			SELECT * FROM users
			WHERE id=$1
`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user User
	row := u.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(&user.Id, &user.FirstName, &user.LastName)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (u UserModel) Update(user *User) error {
	query := `
			UPDATE users 
			SET first_name=$1, last_name=$2
			WHERE id=$3
`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	args := []any{user.FirstName, user.LastName, user.Id}
	_, err := u.DB.ExecContext(ctx, query, args...)
	return err
}

func (u UserModel) Delete(id int) error {
	query := `
			DELETE FROM users
			WHERE id=$1
`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := u.DB.ExecContext(ctx, query, id)
	return err
}
