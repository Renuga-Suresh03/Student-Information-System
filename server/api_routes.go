package server

import (
	"controllers/backend/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupAPIRoutes(router *gin.Engine, adminController *controllers.AdminController, markController *controllers.MarkController, attendanceController *controllers.AttendanceController, studentController *controllers.StudentController) {
	apiRoutes := router.Group("/api")
	{
		adminRoutes := apiRoutes.Group("/admin")
		{
			adminRoutes.POST("/login", func(c *gin.Context) {
				// Admin login route implementation...
			})

			adminRoutes.GET("/profile", func(c *gin.Context) {
				// Admin profile route implementation...
			})
		}

		markRoutes := apiRoutes.Group("/mark")
		{
			markRoutes.POST("/add", func(c *gin.Context) {
				// Add mark route implementation...
			})

			markRoutes.GET("/get", func(c *gin.Context) {
				// Get marks route implementation...
			})
		}

		attendanceRoutes := apiRoutes.Group("/attendance")
		{
			attendanceRoutes.GET("/get", func(c *gin.Context) {
				// Get attendance route implementation...
			})

			attendanceRoutes.POST("/add", func(c *gin.Context) {
				// Add attendance route implementation...
			})
		}

		studentRoutes := apiRoutes.Group("/student")
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
		}
	}
}
