package manager

import (
	"log"
	models "student-teacher-api/model"
	"student-teacher-api/service"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// StudentService handles operations related to Students
type StudentService struct {
	Validator *validator.Validate
}

// ValidateStudent validates the Student struct
func (s *StudentService) ValidateStudent(Student *models.Student) error {
	return s.Validator.Struct(Student)
}

// GetStudents fetches all Students
func (s *StudentService) GetStudents() ([]models.Student, error) {
	return service.GetStudents()
}

// GetStudent fetches a Student by ID
func (s *StudentService) GetStudent(id string) (models.Student, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Student{}, err
	}
	return service.GetStudentByID(objectID)
}

// CreateStudent inserts a new Student and returns the created Student
func (s *StudentService) CreateStudent(Student models.Student) (models.Student, error) {
	return service.InsertStudent(Student)
}

// UpdateStudent updates an existing Student
func (s *StudentService) UpdateStudent(id string, Student models.Student) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return service.UpdateStudent(objectID, Student)
}

// DeleteStudent deletes a Student by ID
func (s *StudentService) DeleteStudent(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	log.Println("Deleted Student!")
	return service.DeleteStudent(objectID)
}
