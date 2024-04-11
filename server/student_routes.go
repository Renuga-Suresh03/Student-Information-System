// server/student_routes.go

package server

import (
	"controllers/backend/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupStudentRoutes(router *gin.Engine, studentController *controllers.StudentController) {
	studentRoutes := router.Group("/student")
	{
		studentRoutes.POST("/login", func(c *gin.Context) {
			var loginRequest struct {
				RegNo string `json:"reg_no"`
				DOB   string `json:"dob"`
			}
			if err := c.BindJSON(&loginRequest); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			student, err := studentController.AuthenticateStudent(loginRequest.RegNo, loginRequest.DOB)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				return
			}

			// Authentication successful, return student details
			c.JSON(http.StatusOK, student)
		})

		studentRoutes.GET("/profile/:reg_no", func(c *gin.Context) {
			regNo := c.Param("reg_no")
			student, err := studentController.GetStudentProfile(regNo)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, student)
		})
		// Add other student routes here
	}
}
