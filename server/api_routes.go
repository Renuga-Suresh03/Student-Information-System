package server

import (
	"controllers/backend/controllers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SetupAPIRoutes(router *gin.Engine, adminController *controllers.AdminController, markController *controllers.MarkController, attendanceController *controllers.AttendanceController, studentController *controllers.StudentController) {
	apiRoutes := router.Group("/api")
	{
		adminRoutes := apiRoutes.Group("/admin")
		{
			adminRoutes.POST("/login", func(c *gin.Context) {
				var loginRequest struct {
					AdminID  string `json:"admin_id"`
					Password string `json:"password"`
				}
				if err := c.BindJSON(&loginRequest); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				admin, err := adminController.AuthenticateAdmin(loginRequest.AdminID, loginRequest.Password)
				if err != nil {
					c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
					return
				}

				c.JSON(http.StatusOK, admin)
			})

			adminRoutes.GET("/profile", func(c *gin.Context) {
				var AdminID string

				if err := c.BindJSON(&AdminID); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				admin, err := adminController.GetAdminProfile(AdminID)
				if err != nil {
					c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
					return
				}

				c.JSON(http.StatusOK, admin)
			})
		}

		markRoutes := apiRoutes.Group("/mark")
		{
			markRoutes.POST("/add", func(c *gin.Context) {
				var requestBody struct {
					RegNo       string `json:"reg_no"`
					ExamNo      int    `json:"exam_no"`
					SubjectCode string `json:"subject_code"`
					Subject     string `json:"subject"`
					Mark        int    `json:"mark"`
				}
				if err := c.BindJSON(&requestBody); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
					return
				}

				if err := markController.AddMark(requestBody.RegNo, requestBody.ExamNo, requestBody.SubjectCode, requestBody.Subject, requestBody.Mark); err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add mark"})
					return
				}

				c.JSON(http.StatusOK, gin.H{"message": "Mark added successfully"})
			})

			markRoutes.GET("/get", func(c *gin.Context) {
				regNo := c.Query("reg_no")
				examNo := c.Query("exam_no")

				if regNo == "" || examNo == "" {
					c.JSON(http.StatusBadRequest, gin.H{"error": "Registration number and exam number are required"})
					return
				}

				// Convert examNo to int
				examNoInt, err := strconv.Atoi(examNo)
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid exam number"})
					return
				}

				marks, err := markController.GetMarks(regNo, examNoInt)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch marks"})
					return
				}

				c.JSON(http.StatusOK, gin.H{"marks": marks})
			})
		}

		attendanceRoutes := apiRoutes.Group("/attendance")
		{
			attendanceRoutes.GET("/get", func(c *gin.Context) {
				regNo := c.Query("reg_no")
				if regNo == "" {
					c.JSON(http.StatusBadRequest, gin.H{"error": "Registration number is required"})
					return
				}

				attendanceRecords, err := attendanceController.GetAttendance(regNo)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch attendance records"})
					return
				}

				c.JSON(http.StatusOK, gin.H{"attendance_records": attendanceRecords})
			})

			attendanceRoutes.POST("/add", func(c *gin.Context) {
				var addAttendanceRequest struct {
					RegNo  string `json:"reg_no"`
					Status string `json:"status"`
				}
				if err := c.BindJSON(&addAttendanceRequest); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				err := attendanceController.AddAttendance(addAttendanceRequest.RegNo, addAttendanceRequest.Status)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}

				c.JSON(http.StatusOK, gin.H{"message": "Attendance added successfully"})
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
