package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Attendance struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty"`
	StudentID            string             `bson:"student_id"`
	AttendanceRecords    []AttendanceRecord `bson:"attendance_records"`
	AttendancePercentage float64            `bson:"attendance_percentage"`
	TotalPresentDays     int                `bson:"total_present_days"`
	TotalWorkingDays     int                `bson:"total_working_days"`
}

type AttendanceRecord struct {
	Date   time.Time `bson:"date"`
	Status string    `bson:"status"`
}
