package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type NewAccount struct {
	Name   string `json:"name" binding:"required"`
	UserID primitive.ObjectID `json:"user_id" bson:"user_id"`
}

type AccountData struct {
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string `json:"name" binding:"required"`
}

type AccountTable struct {
	Name      string `json:"name"`
	CountTask int `json:"count_task" bson:"count_task"`
	Approved  bool `json:"approved"`
}