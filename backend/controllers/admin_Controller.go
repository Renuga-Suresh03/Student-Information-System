package controllers

import (
	"context"
	"controllers/backend/models"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AdminController struct {
	DB *mongo.Database
}

func NewAdminController(db *mongo.Database) *AdminController {
	return &AdminController{DB: db}
}

func (ac *AdminController) AuthenticateAdmin(adminID, pass string) (models.Admin, error) {
	var admin models.Admin

	// Fetch admin details from the database based on admin ID
	collection := ac.DB.Collection("Admin")
	err := collection.FindOne(context.Background(), bson.M{"admin_id": adminID}).Decode(&admin)
	if err != nil {
		return models.Admin{}, errors.New("admin not found")
	}

	// Compare the passwords
	if admin.Password != pass {
		return models.Admin{}, errors.New("invalid password")
	}

	// Authentication successful, return admin details
	return admin, nil
}

func (ac *AdminController) GetAdminProfile(adminID string) (models.Admin, error) {
	var admin models.Admin

	// Fetch admin details from the database based on admin ID
	collection := ac.DB.Collection("Admin")
	err := collection.FindOne(context.Background(), bson.M{"admin_id": adminID}).Decode(&admin)
	if err != nil {
		return models.Admin{}, errors.New("admin not found")
	}

	// Admin profile found, return admin details
	return admin, nil
}
