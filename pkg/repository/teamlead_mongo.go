package repository

import (
	"context"
	"errors"

	"github.com/Traking-work/traking-backend.git/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TeamleadRepo struct {
	db *mongo.Collection
}

func NewTeamleadRepo(db *mongo.Database) *TeamleadRepo {
	return &TeamleadRepo{db: db.Collection(usersCollection)}
}

func (r *TeamleadRepo) GetData(ctx context.Context, userID primitive.ObjectID) (domain.UserData, error) {
	var teamlead domain.UserData

	if err := r.db.FindOne(ctx, bson.M{"_id": userID}).Decode(&teamlead); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.UserData{}, domain.ErrUserNotFound
		}
		return domain.UserData{}, err
	}

	return teamlead, nil
}

func (r *TeamleadRepo) GetStaff(ctx context.Context, userID primitive.ObjectID) ([]domain.UserData, error) {
	var staff []domain.UserData

	cur, err := r.db.Find(ctx, bson.M{"position": "staff", "teamlead": userID})
	if err != nil {
		return nil, err
	}

	if err := cur.All(ctx, &staff); err != nil {
		return nil, err
	}

	return staff, nil
}