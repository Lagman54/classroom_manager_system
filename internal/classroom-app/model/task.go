package model

import (
	"FinalProject/internal/classroom-app/validator"
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

func (t *TaskModel) Insert(task *Task, classroomIds []int) error {
	query := `
		INSERT INTO task (header, description)
		VALUES($1, $2)
		RETURNING id, created_at, updated_at
`

	args := []any{task.Header, task.Description}

	tx, err := t.DB.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	tx.QueryRowContext(ctx, query, args...).Scan(&task.Id, &task.CreatedAt, &task.UpdatedAt)

	query = `
		INSERT INTO classroom_task (class_id, task_id) 
		VALUES ($1, $2)
`
	for _, classId := range classroomIds {
		ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
		_, err = tx.ExecContext(ctx, query, classId, task.Id)
		if err != nil {
			return err
		}
	}

	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
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

func (t *TaskModel) GetTasksOfClass(classId int) (*[]Task, error) {
	query := `
		SELECT task_id FROM classroom_task
		WHERE class_id=$1
`
	rows, err := t.DB.Query(query, classId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var taskId *int
		err = rows.Scan(&taskId)
		if err != nil {
			return &tasks, err
		}

		var task Task
		query = `
				SELECT id, header, description, created_at, updated_at FROM task
				WHERE id=$1
			`

		row := t.DB.QueryRow(query, taskId)
		err = row.Scan(&task.Id, &task.Header, &task.Description, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &tasks, nil
}

func (t *TaskModel) Update(task *Task) error {
	query := `
		UPDATE task
		SET header=$1, description=$2, updated_at=current_timestamp
		WHERE id=$3 and updated_at=$4
		RETURNING updated_at
`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	args := []any{task.Header, task.Description, task.Id, task.UpdatedAt}

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

func ValidateTask(v *validator.Validator, task *Task) {
	v.Check(task.Header != "", "header", "must be provided")
	v.Check(len(task.Header) <= 50, "header", "must be no more than 50 bytes long")
	v.Check(len(task.Description) <= 3000, "description", "must be no more than 1000 bytes long")
}
