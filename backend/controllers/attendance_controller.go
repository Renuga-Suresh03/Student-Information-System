package controllers

import (
	"context"
	"controllers/backend/models"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AttendanceController struct {
	DB *mongo.Database
}

func NewAttendanceController(db *mongo.Database) *AttendanceController {
	return &AttendanceController{DB: db}
}

func (ac *AttendanceController) AddAttendance(regNo, status string) error {
	// Fetch student details from the database based on regNo
	studentCollection := ac.DB.Collection("students")
	var student models.Student
	err := studentCollection.FindOne(context.Background(), bson.M{"reg_no": regNo}).Decode(&student)
	if err != nil {
		return errors.New("student not found")
	}

	// Add attendance record for the current date
	attendanceDate := time.Now().UTC().Truncate(24 * time.Hour)
	attendanceRecord := models.AttendanceRecord{Date: attendanceDate, Status: status}

	// Create or update the attendance document for the student
	attendanceCollection := ac.DB.Collection("attendance")
	filter := bson.M{"student_id": student.ID.Hex()}
	update := bson.M{"$push": bson.M{"attendance_records": attendanceRecord}}
	opts := options.Update().SetUpsert(true)
	_, err = attendanceCollection.UpdateOne(context.Background(), filter, update, opts)
	if err != nil {
		return errors.New("failed to add attendance")
	}

	// Get the last 7 days of attendance records
	last7Days := time.Now().Add(-7 * 24 * time.Hour)
	cursor, err := attendanceCollection.Find(context.Background(), bson.M{
		"student_id":              student.ID.Hex(),
		"attendance_records.date": bson.M{"$gte": last7Days},
	})
	if err != nil {
		return errors.New("failed to fetch attendance records")
	}
	defer cursor.Close(context.Background())

	var attendance models.Attendance
	if err := cursor.All(context.Background(), &attendance); err != nil {
		return errors.New("failed to decode attendance records")
	}

	// Calculate attendance percentage
	var presentDays int
	for _, record := range attendance.AttendanceRecords {
		if record.Status == "Present" {
			presentDays++
		}
	}
	attendance.AttendancePercentage = float64(presentDays) / 7.0 * 100

	// Update the attendance document with the new percentage
	_, err = attendanceCollection.UpdateOne(context.Background(), filter, bson.M{"$set": bson.M{"attendance_percentage": attendance.AttendancePercentage}})
	if err != nil {
		return errors.New("failed to update attendance percentage")
	}

	return nil
}

func (ac *AttendanceController) GetAttendance(regNo string) (*models.Attendance, error) {
	// Fetch student details from the database based on regNo
	studentCollection := ac.DB.Collection("students")
	var student models.Student
	err := studentCollection.FindOne(context.Background(), bson.M{"reg_no": regNo}).Decode(&student)
	if err != nil {
		return nil, errors.New("student not found")
	}

	// Fetch attendance record for the student
	attendanceCollection := ac.DB.Collection("attendance")
	var attendance models.Attendance
	err = attendanceCollection.FindOne(context.Background(), bson.M{"student_id": student.ID.Hex()}).Decode(&attendance)
	if err != nil {
		return nil, errors.New("attendance record not found")
	}

	return &attendance, nil
}
