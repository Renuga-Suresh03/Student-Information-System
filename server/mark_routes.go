// server/mark_routes.go
//get works here
/*package server

import (
	"controllers/backend/controllers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SetupMarkRoutes(router *gin.Engine, markController *controllers.MarkController) {
	markRoutes := router.Group("/api/mark")
	{
		markRoutes.POST("/add", func(c *gin.Context) {
			var requestBody map[string]interface{}
			if err := c.BindJSON(&requestBody); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
				return
			}

			regNo, ok := requestBody["reg_no"].(string)
			if !ok {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid registration number"})
				return
			}

			examNo, ok := requestBody["exam_no"].(int)
			if !ok {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid exam number"})
				return
			}

			subjectCode, ok := requestBody["subject_code"].(string)
			if !ok {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid subject code"})
				return
			}

			subject, ok := requestBody["subject"].(string)
			if !ok {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid subject"})
				return
			}

			mark, ok := requestBody["mark"].(int)
			if !ok {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid mark"})
				return
			}

			if err := markController.AddMark(regNo, examNo, subjectCode, subject, mark); err != nil {
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
}
*/

// demo //get, post working successfully
package server

import (
	"controllers/backend/controllers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SetupMarkRoutes(router *gin.Engine, markController *controllers.MarkController) {
	markRoutes := router.Group("/api/mark")
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
}
