package repository

import (
	"context"

	"github.com/Traking-work/traking-backend.git/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type AdminRepo struct {
	db *mongo.Collection
}

func NewAdminRepo(db *mongo.Database) *AdminRepo {
	return &AdminRepo{db: db.Collection(usersCollection)}
}

func (r *AdminRepo) GetTeamLeads(ctx context.Context) ([]domain.UserData, error) {
	var teamleads []domain.UserData

	cur, err := r.db.Find(ctx, bson.M{"position": "teamlead"})
	if err != nil {
		return nil, err
	}

	if err := cur.All(ctx, &teamleads); err != nil {
		return nil, err
	}

	return teamleads, nil
}

func (r *AdminRepo) AddUser(ctx context.Context, inp domain.UserData) error {
	var err error

	if inp.Position == "staff" {
		_, err = r.db.InsertOne(ctx, inp)
	} else {
		_, err = r.db.InsertOne(ctx, bson.M{"name": inp.Name, "username": inp.Username, "password": inp.Password, "position": inp.Position})
	}

	return err
}