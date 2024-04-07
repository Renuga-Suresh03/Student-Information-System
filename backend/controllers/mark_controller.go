package controllers

import (
	"context"
	"controllers/backend/models"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MarkController struct {
	DB *mongo.Database
}

func NewMarkController(db *mongo.Database) *MarkController {
	return &MarkController{DB: db}
}

func (mc *MarkController) AddMark(regNo string, examNo int, subjectCode string, mark int) error {
	// Fetch student details from the database based on regNo
	studentCollection := mc.DB.Collection("students")
	var student models.Student
	err := studentCollection.FindOne(context.Background(), bson.M{"reg_no": regNo}).Decode(&student)
	if err != nil {
		return errors.New("student not found")
	}

	// Check if examNo is valid
	if examNo < 1 || examNo > 3 {
		return errors.New("invalid exam number")
	}

	// Find the index of the mark for the specified exam
	var markIndex int
	for i, m := range student.Marks {
		if m.ExamNo == examNo {
			markIndex = i
			break
		}
	}

	// Update the mark for the subject in the specified exam
	var subjectFound bool
	for i, s := range student.Marks[markIndex].Subjects {
		if s.SubjectCode == subjectCode {
			student.Marks[markIndex].Subjects[i].Mark = mark
			subjectFound = true
			break
		}
	}
	if !subjectFound {
		return errors.New("subject not found")
	}

	// Update student document in the database
	_, err = studentCollection.ReplaceOne(context.Background(), bson.M{"reg_no": regNo}, student)
	if err != nil {
		return errors.New("failed to update marks")
	}

	return nil
}

func (mc *MarkController) GetMarks(regNo string, examNo int) ([]models.SubjectMark, error) {
	// Fetch student details from the database based on regNo
	studentCollection := mc.DB.Collection("students")
	var student models.Student
	err := studentCollection.FindOne(context.Background(), bson.M{"reg_no": regNo}).Decode(&student)
	if err != nil {
		return nil, errors.New("student not found")
	}

	// Check if examNo is valid
	if examNo < 1 || examNo > 3 {
		return nil, errors.New("invalid exam number")
	}

	// Find the marks for the specified exam
	var marks []models.SubjectMark
	for _, m := range student.Marks {
		if m.ExamNo == examNo {
			marks = m.Subjects
			break
		}
	}

	return marks, nil
}
