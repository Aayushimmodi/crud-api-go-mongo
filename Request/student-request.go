package Request

import (
	models "student-teacher-api/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// StudentRequest is the request structure for creating and updating a Student
type StudentRequest struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Student_name string             `bson:"student_name" json:"student_name" validate:"required"`
	Age          int                `bson:"age" json:"age" validate:"required"`
	Class        string             `bson:"class" json:"class" validate:"required"`
}

// ToModel converts a StudentRequest to a Student model
func (req *StudentRequest) ToModel() models.Student {
	return models.Student{
		ID:           req.ID,
		Student_name: req.Student_name,
		Age:          req.Age,
		Class:        req.Class,
	}
}
