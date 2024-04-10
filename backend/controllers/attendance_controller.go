package controllers

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"controllers/backend/models"
)

type AttendanceController struct {
	DB *mongo.Database
}

func NewAttendanceController(db *mongo.Database) *AttendanceController {
	return &AttendanceController{DB: db}
}

func (ac *AttendanceController) AddAttendance(regNo, status string) error {
	student, err := ac.GetStudentByRegNo(regNo)
	if err != nil {
		return err
	}

	attendanceDate := time.Now().UTC().Truncate(24 * time.Hour)
	attendanceRecord := models.AttendanceRecord{Date: attendanceDate, Status: status}

	attendanceCollection := ac.DB.Collection("Attendance")
	filter := bson.M{"student_id": student.ID.Hex()}
	update := bson.M{"$push": bson.M{"attendance_records": attendanceRecord}}
	opts := options.Update().SetUpsert(true)
	_, err = attendanceCollection.UpdateOne(context.Background(), filter, update, opts)
	if err != nil {
		return errors.New("failed to add attendance")
	}

	_, err = attendanceCollection.UpdateMany(context.Background(), filter,
		bson.M{"$pull": bson.M{"attendance_records": bson.M{"date": bson.M{"$lt": time.Now().Add(-7 * 24 * time.Hour)}}}})
	if err != nil {
		return errors.New("failed to remove old attendance records")
	}

	return nil
}

func (ac *AttendanceController) GetAttendance(regNo string) (*models.Attendance, error) {
	student, err := ac.GetStudentByRegNo(regNo)
	if err != nil {
		return nil, err
	}

	attendanceCollection := ac.DB.Collection("Attendance")
	var attendance models.Attendance
	err = attendanceCollection.FindOne(context.Background(), bson.M{"student_id": student.ID.Hex()}).Decode(&attendance)
	if err != nil {
		return nil, errors.New("attendance record not found")
	}

	return &attendance, nil
}

func (ac *AttendanceController) GetStudentByRegNo(regNo string) (models.Student, error) {
	var student models.Student

	collection := ac.DB.Collection("Student")
	err := collection.FindOne(context.Background(), bson.M{"reg_no": regNo}).Decode(&student)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.Student{}, errors.New("student not found")
		}
		return models.Student{}, err
	}

	return student, nil
}
