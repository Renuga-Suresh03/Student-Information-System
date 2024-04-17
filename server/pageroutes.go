
package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupPageRoutes(router *gin.Engine) {
	router.LoadHTMLGlob("public/*.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	// Admin routes
	router.GET("/admin/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/login.html", gin.H{})
	})

	router.GET("/admin/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/home.html", gin.H{})
	})

	router.GET("/admin/profile", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/profile.html", gin.H{})
	})

	router.GET("/admin/marks", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/marks.html", gin.H{})
	})

	router.GET("/admin/attendance", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/attendance.html", gin.H{})
	})

	// Student routes
	router.GET("/student/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "student/login.html", gin.H{})
	})

	router.GET("/student/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "student/home.html", gin.H{})
	})

	router.GET("/student/profile", func(c *gin.Context) {
		c.HTML(http.StatusOK, "student/profile.html", gin.H{})
	})

	router.GET("/student/marks", func(c *gin.Context) {
		c.HTML(http.StatusOK, "student/marks.html", gin.H{})
	})

	router.GET("/student/attendance", func(c *gin.Context) {
		c.HTML(http.StatusOK, "student/attendance.html", gin.H{})
	})
}
