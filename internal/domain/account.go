package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccountData struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name         string             `json:"name" binding:"required"`
	UserID       primitive.ObjectID `json:"user_id" bson:"user_id"`
	CreateDate   time.Time          `json:"create_date" bson:"create_date"`
	StatusDelete bool               `json:"status_delete" bson:"status_delete"`
	DeleteDate   time.Time          `json:"delete_date" bson:"delete_date"`
}

type AccountPack struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name"`
	CountTask int                `json:"count_task" bson:"count_task"`
	Payment   float32            `json:"payment"`
	Approved  bool               `json:"approved"`
	Date      string             `json:"date"`
}

type Date struct {
	Date time.Time `json:"date"`
}
