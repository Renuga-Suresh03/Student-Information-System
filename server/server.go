package server

import (
	"context"
	"controllers/backend/controllers"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Server struct {
	app                  *fiber.App
	adminController      *controllers.AdminController
	studentController    *controllers.StudentController
	markController       *controllers.MarkController
	attendanceController *controllers.AttendanceController
}

func NewServer() *Server {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("failed to connect to MongoDB: %v", err)
	}

	// Access the database from the client
	db := client.Database("your_database_name")

	return &Server{
		app:                  fiber.New(),
		adminController:      controllers.NewAdminController(db),
		studentController:    controllers.NewStudentController(db),
		markController:       controllers.NewMarkController(db),
		attendanceController: controllers.NewAttendanceController(db),
	}
}
func (s *Server) Start(port string) error {
	s.setupAdminRoutes()
	s.setupStudentRoutes()
	s.setupMarkRoutes()
	s.setupAttendanceRoutes()

	return s.app.Listen(":" + port)
}
