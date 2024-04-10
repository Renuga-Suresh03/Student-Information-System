// server.go

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
	app             *fiber.App
	adminController *controllers.AdminController
}

func NewServer() *Server {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017/STU"))
	if err != nil {
		log.Fatalf("failed to connect to MongoDB: %v", err)
	} 

	db := client.Database("STU")

	return &Server{
		app:             fiber.New(),
		adminController: controllers.NewAdminController(db),
	}
}

func (s *Server) Start(port string) error {
	s.setupAdminRoutes()

	return s.app.Listen(":" + port)
}
