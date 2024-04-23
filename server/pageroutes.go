package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupPageRoutes(router *gin.Engine) {
	// Set base URL for static assets
	router.Static("/assets", "./frontend/assets")

	router.LoadHTMLGlob("frontend/*.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"baseURL": "/",
		})
	})

	router.GET("/admin/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/login.html", gin.H{
			"baseURL": "/",
		})
	})

	// Add other routes...
	router.GET("/admin/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/home.html", gin.H{})
	})

	router.GET("/admin/profile", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/profile.html", gin.H{})
	})

	router.GET("/admin/mark-IAT1", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/mark-IAT1.html", gin.H{})
	})

	router.GET("/admin/mark-IAT2", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/mark-IAT2.html", gin.H{})
	})

	router.GET("/admin/mark-model", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/mark-model.html", gin.H{})
	})

	router.GET("/admin/attendance", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/attendance.html", gin.H{})
	})

	router.GET("/student/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "student/login.html", gin.H{})
	})

	router.GET("/student/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "student/home.html", gin.H{})
	})

	router.GET("/student/profile", func(c *gin.Context) {
		c.HTML(http.StatusOK, "student/profile.html", gin.H{})
	})

	router.GET("/student/assessment1", func(c *gin.Context) {
		c.HTML(http.StatusOK, "student/assessment1.html", gin.H{})
	})

	router.GET("/student/assessment2", func(c *gin.Context) {
		c.HTML(http.StatusOK, "student/assessment2.html", gin.H{})
	})

	router.GET("/student/model", func(c *gin.Context) {
		c.HTML(http.StatusOK, "student/model.html", gin.H{})
	})

	router.GET("/student/attendance", func(c *gin.Context) {
		c.HTML(http.StatusOK, "student/attendance.html", gin.H{})
	})

}
