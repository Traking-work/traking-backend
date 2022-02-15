package repository

import (
	"context"

	"github.com/Traking-work/traking-backend.git/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StaffRepo struct {
	db *mongo.Collection
}

func NewStaffRepo(db *mongo.Database) *StaffRepo {
	return &StaffRepo{db: db.Collection(accountsCollection)}
}

func (r *StaffRepo) GetAccounts(ctx context.Context, userID primitive.ObjectID) ([]domain.AccountData, error) {
	var accounts []domain.AccountData

	cur, err := r.db.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}

	if err := cur.All(ctx, &accounts); err != nil {
		return nil, err
	}

	return accounts, nil
}
	
func (r *StaffRepo) AddAccount(ctx context.Context, account domain.NewAccount, userID primitive.ObjectID) error {
	_, err := r.db.InsertOne(ctx, bson.M{"name": account.Name, "user_id": userID})
	return err
}