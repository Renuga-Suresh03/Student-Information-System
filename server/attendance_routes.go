// server/attendance_routes.go

/*
package server

import (

	"controllers/backend/controllers"
	"net/http"



	"github.com/gin-gonic/gin"

)

	func SetupAttendanceRoutes(router *gin.Engine, attendanceController *controllers.AttendanceController) {
		attendanceRoutes := router.Group("/api/attendance")
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
		}
	}
*/
package server

import (
	"controllers/backend/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupAttendanceRoutes(router *gin.Engine, attendanceController *controllers.AttendanceController) {
	attendanceRoutes := router.Group("/api/attendance")
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
}
