// admincontroller.go

package controllers

import (
	"context"
	"controllers/backend/models"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AdminController struct {
	DB *mongo.Database
}

func NewAdminController(db *mongo.Database) *AdminController {
	return &AdminController{DB: db}
}

func (ac *AdminController) AuthenticateAdmin(adminID, password string) (models.Admin, error) {
	var admin models.Admin

	collection := ac.DB.Collection("Admin")
	err := collection.FindOne(context.Background(), bson.M{"admin_id": adminID}).Decode(&admin)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Printf("Admin with ID %s not found", adminID)
			return models.Admin{}, errors.New("admin not found")
		} else if errors.Is(err, mongo.ErrClientDisconnected) {
			log.Printf("Failed to connect to the database: %v", err)
			return models.Admin{}, errors.New("failed to connect to database")
		} else {
			log.Printf("Error querying admin document: %v", err)
			return models.Admin{}, errors.New("failed to authenticate admin")
		}
	}

	if admin.Password != password {
		log.Printf("Invalid password for admin with ID %s", adminID)
		return models.Admin{}, errors.New("invalid password")
	}

	log.Printf("Admin with ID %s successfully authenticated", adminID)
	return admin, nil
}

func (ac *AdminController) GetAdminProfile(adminID string) (models.Admin, error) {
	var admin models.Admin

	collection := ac.DB.Collection("Admin")
	err := collection.FindOne(context.Background(), bson.M{"admin_id": adminID}).Decode(&admin)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.Admin{}, errors.New("admin not found")
		} else if errors.Is(err, mongo.ErrClientDisconnected) {
			return models.Admin{}, errors.New("failed to connect to database")
		}
		return models.Admin{}, errors.New("failed to get admin profile")
	}

	return admin, nil
}
