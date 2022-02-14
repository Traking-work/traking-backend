package repository

import (
	"context"

	"github.com/Traking-work/traking-backend.git/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type AdminRepo struct {
	db *mongo.Collection
}

func NewAdminRepo(db *mongo.Database) *AdminRepo {
	return &AdminRepo{db: db.Collection(usersCollection)}
}

func (r *AdminRepo) AddUser(ctx context.Context, inp domain.NewUser) error {
	_, err := r.db.InsertOne(ctx, inp)
	return err
}