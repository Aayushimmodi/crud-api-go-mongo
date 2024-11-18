package main

import (
	"fmt"
	"log"
	"os"
	"student-teacher-api/config"
	"student-teacher-api/controller"
	"student-teacher-api/manager"
	"student-teacher-api/routes"
	"student-teacher-api/service"

	_ "github.com/lib/pq"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

// CustomValidator implements Echo's Validator interface
type CustomValidator struct {
	Validator *validator.Validate
}

// Validate validates the struct using go-playground/validator
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default configurations.")
	}
	// Connect to the database
	client, err := config.ConnectDatabase()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	// Set the student collection in the service
	service.SetStudentCollection(client, "students")

	// Initialize the Echo instance
	e := echo.New()

	// Set the default validator to the Echo instance
	e.Validator = &CustomValidator{Validator: validator.New()}

	// Initialize the movie service
	studentService := &manager.StudentService{}

	// Initialize the controller with the movieService instance
	StudentController := &controller.StudentController{Service: studentService}

	// Setup routes with the movieController
	routes.SetupRoutes(e, StudentController)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "9091" // default port
	}
	address := fmt.Sprintf(":%s", port)
	if err := e.Start(address); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
