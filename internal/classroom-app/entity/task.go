package entity

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type Task struct {
	Id          int    `json:"id"`
	Header      string `json:"header"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"UpdatedAt"`
}

type TaskModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func (t *TaskModel) Insert(task *Task) error {
	query := `
		INSERT INTO task (header, description)
		VALUES($1, $2)
		RETURNING id, created_at, updated_at
`
	args := []any{task.Header, task.Description}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return t.DB.QueryRowContext(ctx, query, args...).Scan(&task.Id, &task.CreatedAt, &task.UpdatedAt)
}

func (t *TaskModel) Get(id int) (*Task, error) {
	query := `
		SELECT id, header, description, created_at, updated_at FROM task
		WHERE id=$1
`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var task Task
	row := t.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(&task.Id, &task.Header, &task.Description, &task.CreatedAt, &task.UpdatedAt)

	if err != nil {
		return nil, err
	}
	return &task, err
}

func (t *TaskModel) Update(task *Task) error {
	query := `
		UPDATE task
		SET header=$1, description=$2
		WHERE id=$3
		RETURNING updated_at
`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	args := []any{task.Header, task.Description, task.Id}

	return t.DB.QueryRowContext(ctx, query, args...).Scan(&task.UpdatedAt)
}

func (t *TaskModel) Delete(id int) error {
	query := `
		DELETE FROM task
		WHERE id=$1
`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := t.DB.ExecContext(ctx, query, id)
	return err
}
