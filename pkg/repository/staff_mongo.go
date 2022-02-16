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
	
func (r *StaffRepo) AddAccount(ctx context.Context, account domain.NewAccount) error {
	_, err := r.db.InsertOne(ctx, account)
	return err
}

func (r *StaffRepo) GetDataAccount(ctx context.Context, accountID primitive.ObjectID) ([]domain.AccountTable, error) {
	var dataAccount []domain.AccountTable

	cur, err := r.db.Database().Collection(packAccountsCollection).Find(ctx, bson.M{"account_id": accountID})
	if err != nil {
		return nil, err
	}

	if err := cur.All(ctx, &dataAccount); err != nil {
		return nil, err
	}

	return dataAccount, nil
}