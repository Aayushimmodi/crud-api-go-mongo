package utils

import (
	models "student-teacher-api/model"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStudent(student *models.Student) error {
	return validate.Struct(student)
}
