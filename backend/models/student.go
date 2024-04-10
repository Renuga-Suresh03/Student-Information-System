package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	RegNo         string             `bson:"reg_no"`
	Name          string             `bson:"name"`
	DateOfBirth   string             `bson:"date_of_birth"`
	Year          int32              `bson:"year"`
	Department    string             `bson:"department"`
	Section       string             `bson:"section"`
	YearOfJoining int32              `bson:"year_of_joining"`
	YearOfPassing int32              `bson:"year_of_passing"`
	CurrentSem    int                `bson:"current_sem"`
	ClassID       string             `bson:"class_id"`
}
