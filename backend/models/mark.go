package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Mark struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	StudentID primitive.ObjectID `bson:"student_id"`
	ExamNo    int                `bson:"exam_no"`
	Subjects  []SubjectMark      `bson:"subjects"`
}
