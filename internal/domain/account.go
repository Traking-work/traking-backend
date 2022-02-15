package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type NewAccount struct {
	Name   string `json:"name" binding:"required"`
	UserID string `json:"user_id" binding:"required"`
}

type AccountData struct {
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string `json:"name" binding:"required"`
}