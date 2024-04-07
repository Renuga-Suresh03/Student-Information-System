package controllers

import (
	"context"
	"controllers/backend/models"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type StudentController struct {
	DB *mongo.Database
}

func NewStudentController(db *mongo.Database) *StudentController {
	return &StudentController{DB: db}
}

func (sc *StudentController) AuthenticateStudent(regNo string, dob time.Time) (models.Student, error) {
	var student models.Student

	// Fetch student details from the database based on regNo and dob
	collection := sc.DB.Collection("students")
	err := collection.FindOne(context.Background(), bson.M{"reg_no": regNo, "date_of_birth": dob}).Decode(&student)
	if err != nil {
		return models.Student{}, errors.New("student not found")
	}

	// Authentication successful, return student details
	return student, nil
}

func (sc *StudentController) GetStudentProfile(regNo string) (models.Student, error) {
	var student models.Student

	// Fetch student details from the database based on regNo
	collection := sc.DB.Collection("students")
	err := collection.FindOne(context.Background(), bson.M{"reg_no": regNo}).Decode(&student)
	if err != nil {
		return models.Student{}, errors.New("student not found")
	}

	// Student found, return student details
	return student, nil
}
