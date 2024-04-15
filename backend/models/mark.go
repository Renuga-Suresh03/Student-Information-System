package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Mark struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"` // Unique identifier for the mark entry
	StudentID string             `bson:"student_id"`    // ID of the student associated with these marks
	ExamNo    int                `bson:"exam_no"`       // Exam number
	Subjects  []SubjectMark      `bson:"subjects"`      // Marks obtained in different subjects
}

// SubjectMark represents the marks obtained by a student in a specific subject.
type SubjectMark struct {
	SubjectCode string `bson:"subject_code"` // Code of the subject
	Subject     string `bson:"subject"`      // Name of the subject
	Mark        int    `bson:"mark"`         // Marks obtained in the subject
}
