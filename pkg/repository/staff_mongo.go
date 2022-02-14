package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type StaffRepo struct {
	db *mongo.Collection
}

func NewStaffRepo(db *mongo.Database) *StaffRepo {
	return &StaffRepo{db: db.Collection(usersCollection)}
}