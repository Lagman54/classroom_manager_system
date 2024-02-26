package entity

import "database/sql"

type Models struct {
	Classrooms ClassroomModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Classrooms: ClassroomModel{
			DB: db,
		},
	}
}
