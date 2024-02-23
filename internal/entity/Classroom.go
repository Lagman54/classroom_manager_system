package entity

import (
	"context"
	"database/sql"
	"time"
)

type Classroom struct {
	Id       int    `json:"id"`
	Name     string `json:"title"`
	Students []User `json:"users"`
	Teachers []User `json:"teachers"`
}

type ClassroomModel struct {
	DB *sql.DB
}

func (c ClassroomModel) Insert(classroom *Classroom) error {
	query := `INSERT INTO classroom (name) VALUES($1) RETURNING id`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return c.DB.QueryRowContext(ctx, query, classroom.Name).Scan(&classroom.Id)
}

func (c ClassroomModel) Get(id int) (*Classroom, error) {
	query := `SELECT * FROM classroom WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var classRoom Classroom
	row := c.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(&classRoom.Id, &classRoom.Name)

	// TODO add Students and Teachers to classRoom

	if err != nil {
		return nil, err
	}
	return &classRoom, nil
}
