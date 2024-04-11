package controllers

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"controllers/backend/models"
)

type AttendanceController struct {
	DB *mongo.Database
}

func NewAttendanceController(db *mongo.Database) *AttendanceController {
	return &AttendanceController{DB: db}
}

func (ac *AttendanceController) AddAttendance(regNo, status string) error {
	// Check if the student exists
	collection := ac.DB.Collection("Student")
	var student models.Student
	err := collection.FindOne(context.Background(), bson.M{"reg_no": regNo}).Decode(&student)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return errors.New("student not found")
		}
		return err
	}

	// Update total working days
	attendanceCollection := ac.DB.Collection("Attendance")
	filter := bson.M{"student_id": student.ID.Hex()} // Convert ObjectID to hex string
	update := bson.M{"$inc": bson.M{"total_working_days": 1}}
	_, err = attendanceCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return errors.New("failed to update total working days")
	}

	// Increment total present days if status is "Present"
	if status == "Present" {
		update = bson.M{"$inc": bson.M{"total_present_days": 1}}
		_, err = attendanceCollection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			return errors.New("failed to update total present days")
		}
	}

	// Check if there are existing attendance records
	var attendance models.Attendance
	err = attendanceCollection.FindOne(context.Background(), filter).Decode(&attendance)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			// If no records exist, create a new one
			attendance = models.Attendance{
				StudentID:         student.ID.Hex(), // Convert ObjectID to hex string
				AttendanceRecords: []models.AttendanceRecord{},
			}
			_, err = attendanceCollection.InsertOne(context.Background(), attendance)
			if err != nil {
				return errors.New("failed to create new attendance record")
			}
		} else {
			return err
		}
	}

	// Add the new attendance record
	attendanceDate := time.Now().UTC().Truncate(24 * time.Hour)
	newRecord := models.AttendanceRecord{Date: attendanceDate, Status: status}
	update = bson.M{"$push": bson.M{"attendance_records": newRecord}}
	_, err = attendanceCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return errors.New("failed to add new attendance record")
	}

	// Update attendance percentage
	err = ac.updateAttendancePercentage(student.ID)
	if err != nil {
		return err
	}

	return nil
}

func (ac *AttendanceController) updateAttendancePercentage(studentID primitive.ObjectID) error {
	// Fetch student's attendance record
	attendanceCollection := ac.DB.Collection("Attendance")
	filter := bson.M{"student_id": studentID.Hex()} // Convert ObjectID to hex string
	var attendance models.Attendance
	err := attendanceCollection.FindOne(context.Background(), filter).Decode(&attendance)
	if err != nil {
		return err
	}

	// Calculate attendance percentage
	var attendancePercentage float64
	if attendance.TotalWorkingDays > 0 {
		attendancePercentage = float64(attendance.TotalPresentDays) / float64(attendance.TotalWorkingDays) * 100
	}

	// Update attendance percentage in the database
	update := bson.M{"$set": bson.M{"attendance_percentage": attendancePercentage}}
	_, err = attendanceCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return errors.New("failed to update attendance percentage")
	}

	return nil
}
func (ac *AttendanceController) GetAttendance(regNo string) ([]models.AttendanceRecord, error) {
	// Find the student
	collection := ac.DB.Collection("Student")
	var student models.Student
	err := collection.FindOne(context.Background(), bson.M{"reg_no": regNo}).Decode(&student)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("student not found")
		}
		return nil, err
	}

	// Find the attendance record
	attendanceCollection := ac.DB.Collection("Attendance")
	filter := bson.M{"student_id": student.ID.Hex()} // Convert ObjectID to hex string
	var attendance models.Attendance
	err = attendanceCollection.FindOne(context.Background(), filter).Decode(&attendance)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("attendance record not found")
		}
		return nil, err
	}

	// Limit the number of records to the last 7 if there are more than 7
	records := attendance.AttendanceRecords
	if len(records) > 7 {
		records = records[len(records)-7:]
	}

	// Return only the date and status
	var result []models.AttendanceRecord
	for _, record := range records {
		result = append(result, models.AttendanceRecord{
			Date:   record.Date,
			Status: record.Status,
		})
	}

	return result, nil
}
