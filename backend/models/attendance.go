package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Attendance struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	StudentID        string             `bson:"student_id"`
	Date             time.Time          `bson:"date"`
	AttendanceStatus string             `bson:"attendance_status"`
}
