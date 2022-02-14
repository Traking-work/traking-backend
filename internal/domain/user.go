package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username string             `json:"username" binding:"required"`
	Password string             `json:"password" binding:"required"`
	Position string 			`json:"position"`
}

type NewUser struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Position string `json:"position" binding:"required"`
	TeamLead string `json:"teamlead"`
}