package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	RegNo         string             `bson:"reg_no"`
	Name          string             `bson:"name"`
	DateOfBirth   time.Time          `bson:"date_of_birth"`
	Year          time.Time          `bson:"year"`
	Department    string             `bson:"department"`
	Section       string             `bson:"section"`
	YearOfJoining time.Time          `bson:"year_of_joining"`
	YearOfPassing time.Time          `bson:"year_of_passing"`
	CurrentSem    int                `bson:"current_sem"`
	ClassID       string             `bson:"class_id"`
}
