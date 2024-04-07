package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SubjectMark struct {
	SubjectCode string `bson:"subject_code"`
	Subject     string `bson:"subject"`
	Mark        int    `bson:"mark"`
}

type Mark struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	StudentID  string             `bson:"student_id"`
	Subjects   []SubjectMark      `bson:"subjects"`
	PassOrFail string             `bson:"pass_or_fail"`
	Total      int                `bson:"total"`
	Rank       int                `bson:"rank"`
}
