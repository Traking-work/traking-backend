package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserLogin struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username string             `json:"username" binding:"required"`
	Password string             `json:"password" binding:"required"`
	Position string 			`json:"position"`
}

type UserData struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Position string `json:"position" binding:"required"`
	TeamLead primitive.ObjectID `json:"teamlead"`
}

type UserSelect struct {
	Value primitive.ObjectID `json:"value" binding:"required"`
	Name  string `json:"name" binding:"required"`
}