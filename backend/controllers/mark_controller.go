// working successfully
package controllers

import (
	"context"
	"errors"

	"controllers/backend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MarkController struct {
	DB *mongo.Database
}

func NewMarkController(db *mongo.Database) *MarkController {
	return &MarkController{DB: db}
}

func (mc *MarkController) AddMark(regNo string, examNo int, subjectCode string, subject string, mark int) error {
	// Fetch student details from the database based on regNo
	student, err := mc.GetStudentByRegNo(regNo)
	if err != nil {
		return err
	}

	// Check if examNo is valid
	if examNo < 1 || examNo > 3 {
		return errors.New("invalid exam number")
	}

	// Update the existing mark document for the student
	markCollection := mc.DB.Collection("Mark")
	filter := bson.M{"student_id": student.ID.Hex(), "exam_no": examNo}
	update := bson.M{
		"$set": bson.M{
			"student_id": student.ID.Hex(),
			"exam_no":    examNo,
		},
		"$push": bson.M{
			"subjects": bson.M{
				"subject_code": subjectCode,
				"subject":      subject,
				"mark":         mark,
			},
		},
	}
	opts := options.Update().SetUpsert(true)
	_, err = markCollection.UpdateOne(context.Background(), filter, update, opts)
	if err != nil {
		return errors.New("failed to add mark")
	}

	return nil
}

/*func (mc *MarkController) GetMarks(regNo string, examNo int) ([]models.Mark, error) {
	// Fetch student details from the database based on regNo
	student, err := mc.GetStudentByRegNo(regNo)
	if err != nil {
		return nil, err
	}

	// Fetch marks from the Mark collection based on student ID and exam number
	markCollection := mc.DB.Collection("Mark")
	filter := bson.M{"student_id": student.ID.Hex(), "exam_no": examNo}
	cursor, err := markCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var marks []models.Mark
	err = cursor.All(context.Background(), &marks)
	if err != nil {
		return nil, err
	}

	return marks, nil
}*/

func (mc *MarkController) GetMarks(regNo string, examNo int) ([]models.SubjectMark, error) {
	// Fetch student details from the database based on regNo
	student, err := mc.GetStudentByRegNo(regNo)
	if err != nil {
		return nil, err
	}

	// Fetch marks from the Mark collection based on student ID and exam number
	markCollection := mc.DB.Collection("Mark")
	filter := bson.M{"student_id": student.ID.Hex(), "exam_no": examNo}
	var mark models.Mark
	err = markCollection.FindOne(context.Background(), filter).Decode(&mark)
	if err != nil {
		return nil, err
	}

	// Extract subject marks from the Mark object
	subjectMarks := mark.Subjects

	return subjectMarks, nil
}

func (mc *MarkController) GetStudentByRegNo(regNo string) (models.Student, error) {
	var student models.Student

	collection := mc.DB.Collection("Student")
	err := collection.FindOne(context.Background(), bson.M{"reg_no": regNo}).Decode(&student)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.Student{}, errors.New("student not found")
		}
		return models.Student{}, err
	}

	return student, nil
}
