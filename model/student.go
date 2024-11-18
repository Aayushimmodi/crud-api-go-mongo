package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Student struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id" `
	Student_name string             `bson:"student_name" json:"student_name" validate:"required"`
	Age          int                `bson:"age" json:"age" validate:"required"`
	Class        string             `bson:"class" json:"class" validate:"required"`
}
