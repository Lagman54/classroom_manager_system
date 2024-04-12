package entity

import (
	"FinalProject/internal/classroom-app/validator"
	"context"
	"database/sql"
	"log"
	"time"
)

type Classroom struct {
	Id          int    `json:"id"`
	CreatedAt   string `json:"createdAt"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Students    []User `json:"users"`
	Teachers    []User `json:"teachers"`
}

type ClassroomModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

// Insert new classroom into the database
func (c ClassroomModel) Insert(classroom *Classroom) error {
	query := `
		INSERT INTO classroom (name, description) 
		VALUES($1, $2)
		RETURNING id, created_at
		`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	args := []any{classroom.Name, classroom.Description}
	return c.DB.QueryRowContext(ctx, query, args...).Scan(&classroom.Id, &classroom.CreatedAt)
}

// Get classroom from the database
func (c ClassroomModel) Get(id int) (*Classroom, error) {
	query := `
		SELECT id, name, description, created_at FROM classroom 
		WHERE id = $1
		`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var classRoom Classroom
	row := c.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(&classRoom.Id, &classRoom.Name, &classRoom.Description, &classRoom.CreatedAt)

	// TODO add Students and Teachers to classRoom

	if err != nil {
		return nil, err
	}
	return &classRoom, nil
}

// Update classroom in the database
func (c ClassroomModel) Update(classroom *Classroom) error {
	query := `
		UPDATE classroom 
		SET name=$1, description=$2
		WHERE id=$3
	`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	args := []any{classroom.Name, classroom.Description, classroom.Id}

	_, err := c.DB.ExecContext(ctx, query, args...)
	return err
}

// Delete classroom from database
func (c ClassroomModel) Delete(id int) error {
	query := `
		DELETE FROM classroom
		WHERE id=$1
	`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := c.DB.ExecContext(ctx, query, id)
	return err
}

func ValidateClassroom(v *validator.Validator, classroom *Classroom) {
	v.Check(classroom.Name != "", "title", "must be provided")
	v.Check(len(classroom.Name) <= 50, "title", "must be no more than 50 bytes long")
	v.Check(len(classroom.Description) <= 1000, "description", "must be no more than 1000 bytes long")
}
