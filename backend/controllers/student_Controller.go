package controllers

import (
	"context"
	"controllers/backend/models"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type StudentController struct {
	DB *mongo.Database
}

func NewStudentController(db *mongo.Database) *StudentController {
	return &StudentController{DB: db}
}

// AuthenticateStudent authenticates a student based on registration number and date of birth.
func (sc *StudentController) AuthenticateStudent(regNo string, dob string) (models.Student, error) {
	var student models.Student

	collection := sc.DB.Collection("Student")
	err := collection.FindOne(context.Background(), bson.M{"reg_no": regNo, "date_of_birth": dob}).Decode(&student)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.Student{}, errors.New("student not found")
		}
		return models.Student{}, err
	}

	return student, nil
}

// GetStudentProfile retrieves the profile of a student based on registration number.
func (sc *StudentController) GetStudentProfile(regNo string) (models.Student, error) {
	var student models.Student

	collection := sc.DB.Collection("Student")
	err := collection.FindOne(context.Background(), bson.M{"reg_no": regNo}).Decode(&student)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.Student{}, errors.New("student not found")
		}
		return models.Student{}, err
	}

	return student, nil
}
