package controller

import (
	"log"
	"net/http"
	"student-teacher-api/Request"
	"student-teacher-api/Response"
	"student-teacher-api/manager"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// StudentController handles HTTP requests related to Students
type StudentController struct {
	Service *manager.StudentService
}

// GetStudents handler to fetch all Students
func (c *StudentController) GetStudents(ctx echo.Context) error {
	students, err := c.Service.GetStudents()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	// Convert Students to response models
	StudentResponses := make([]Response.StudentResponse, 0, len(students))
	for _, student := range students {
		StudentResponses = append(StudentResponses, Response.FromModel(student))
	}
	log.Println("Returned all students")
	return ctx.JSON(http.StatusOK, StudentResponses)
}

// GetStudentByID handler to fetch a Student by ID
func (c *StudentController) GetStudentByID(ctx echo.Context) error {
	id := ctx.Param("id")
	Student, err := c.Service.GetStudent(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err.Error())
	}
	log.Println("Returned student by ID")
	return ctx.JSON(http.StatusOK, Response.FromModel(Student))
}

// CreateStudent handler to create a new Student
func (c *StudentController) CreateStudent(ctx echo.Context) error {
	var req Request.StudentRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	// Validate the request
	if err := ctx.Validate(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	// Convert request model to main model
	student := req.ToModel()
	result, err := c.Service.CreateStudent(student)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	log.Println("Created new student")
	return ctx.JSON(http.StatusCreated, Response.FromModel(result))
}

func (c *StudentController) UpdateStudent(ctx echo.Context) error {
	// Get the ID from the URL parameter
	id := ctx.Param("id")

	// Convert string ID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid ID format")
	}

	var req Request.StudentRequest

	// Bind the request data to the StudentRequest struct
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	// Validate the request
	if err := ctx.Validate(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	// Convert request model to main model (ensure id is correctly set)
	student := req.ToModel()
	student.ID = objectID // Set the Student ID as ObjectID

	// Call the service to update the Student
	if err := c.Service.UpdateStudent(id, student); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	log.Println("Updated student successfully")

	// Return the updated Student with the ID
	return ctx.JSON(http.StatusOK, student)
}

// DeleteStudent handler to delete a Student
func (c *StudentController) DeleteStudent(ctx echo.Context) error {
	id := ctx.Param("id")
	if err := c.Service.DeleteStudent(id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, "Student deleted")
}
