package server

import (
	"controllers/backend/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupAdminRoutes(router *gin.Engine, adminController *controllers.AdminController) {
	adminRoutes := router.Group("/admin")
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

			// Authentication successful, return admin details
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

		/*adminRoutes.GET("/profile", func(c *gin.Context) {
			adminID := c.Query("admin_id")
			admin, err := adminController.GetAdminProfile(adminID)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, admin)
		})*/

		// Add other admin routes here
	}
}
