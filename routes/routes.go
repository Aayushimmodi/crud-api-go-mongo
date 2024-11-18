package routes

import (
	"student-teacher-api/controller"

	"github.com/labstack/echo/v4"
)

// SetupRoutes sets up the routes for the Student API
func SetupRoutes(e *echo.Echo, controller *controller.StudentController) {
	e.GET("/students", controller.GetStudents)
	e.GET("/students/:id", controller.GetStudentByID)
	e.POST("/students", controller.CreateStudent)
	e.PUT("/students/:id", controller.UpdateStudent)
	e.DELETE("/students/:id", controller.DeleteStudent)
}
