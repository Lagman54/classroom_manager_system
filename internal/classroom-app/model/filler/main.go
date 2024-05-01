package filler

import "FinalProject/internal/classroom-app/model"

func PopulateDatabase(models model.Models) error {
	for _, class := range classrooms {
		err := models.Classrooms.Insert(&class)
		if err != nil {
			return err
		}
	}
	return nil
}

var classrooms = []model.Classroom{
	{Name: "Golang application development", Description: "Golang class"},
	{Name: "Calculus 1", Description: "Calculus class"},
	{Name: "OOP", Description: "Object-Oriented Programming class"},
	{Name: "Physics 1", Description: "Physics class"},
	{Name: "Statistics", Description: "Boring class"},
	{Name: "Electronics", Description: "Electronics class"},
	{Name: "FEE", Description: "Foundations of Electric Engineering"},
	{Name: "Linear algebra", Description: "Linear algebra class"},
	{Name: "English C1", Description: "English C1 class"},
	{Name: "Java Spring", Description: "Java spring class"},
	{Name: "Calculus 2", Description: "Calculus class"},
	{Name: "Physics 2", Description: "Physics class"},
	{Name: "History", Description: "History class"},
	{Name: "Philosophy", Description: "Philosophy class"},
	{Name: "Cryptography", Description: "Cryptography class"},
}
