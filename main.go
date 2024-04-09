/*package main

import (
	"context"
	"controllers/backend/controllers"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	mongoURI := "mongodb://localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("Error creating MongoDB client:", err)
	}
	defer client.Disconnect(context.Background())
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}
	adminController := controllers.NewAdminController(client.Database("STU"))
	adminID := "admin123"
	password := "secretpassword"
	admin, err := adminController.AuthenticateAdmin(adminID, password)
	if err != nil {
		fmt.Printf("Authentication failed: %v\n", err)
		return
	}
	fmt.Printf("Authentication successful! Admin Details: %+v\n", admin)
	adminProfile, err := adminController.GetAdminProfile(adminID)
	if err != nil {
		fmt.Printf("Error retrieving admin profile: %v\n", err)
		return
	}
	fmt.Printf("Admin Profile: %+v\n", adminProfile)
}
*/