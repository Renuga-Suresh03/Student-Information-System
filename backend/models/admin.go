package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Admin struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	AdminID         string             `bson:"admin_id"`
	Name            string             `bson:"name"`
	DateOfBirth     time.Time          `bson:"date_of_birth"`
	Department      string             `bson:"department"`
	Designation     string             `bson:"designation"`
	YearOfJoining   time.Time          `bson:"year_of_joining"`
	InchargeClassID string             `bson:"incharge_class_id"`
}
