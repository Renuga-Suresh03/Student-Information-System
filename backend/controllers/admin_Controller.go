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

	collection := ac.DB.Collection("Admin")
	err := collection.FindOne(context.Background(), bson.M{"admin_id": adminID}).Decode(&admin)
	if err != nil {
		return models.Admin{}, errors.New("admin not found")
	}

	if admin.Password != pass {
		return models.Admin{}, errors.New("invalid password")
	}

	return admin, nil
}

func (ac *AdminController) GetAdminProfile(adminID string) (models.Admin, error) {
	var admin models.Admin
	collection := ac.DB.Collection("Admin")
	err := collection.FindOne(context.Background(), bson.M{"admin_id": adminID}).Decode(&admin)
	if err != nil {
		return models.Admin{}, errors.New("admin not found")
	}
	return admin, nil
}
