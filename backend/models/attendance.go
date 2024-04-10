package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Attendance struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty"`
	StudentID            string             `bson:"student_id"`
	AttendanceRecords    []AttendanceRecord `bson:"attendance_records"`
	AttendancePercentage int32              `bson:"attendance_percentage"`
}

type AttendanceRecord struct {
	Date   time.Time `bson:"date"`
	Status string    `bson:"status"`
}
