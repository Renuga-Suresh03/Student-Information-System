// server/server.go

package server

import (
	"controllers/backend/controllers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupServer(db *mongo.Database) *gin.Engine {
	router := gin.Default()

	// Initialize controllers
	adminController := controllers.NewAdminController(db)
	studentController := controllers.NewStudentController(db)
	markController := controllers.NewMarkController(db)
	attendanceController := controllers.NewAttendanceController(db)

	// Set up routes
	SetupAdminRoutes(router, adminController)
	SetupStudentRoutes(router, studentController)
	SetupMarkRoutes(router, markController)
	SetupAttendanceRoutes(router, attendanceController)
	SetupPageRoutes(router) // Add page routes

	return router
}
