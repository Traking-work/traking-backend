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

func (r *StaffRepo) AddPack(ctx context.Context, accountID primitive.ObjectID, pack domain.AccountTable) error {
	_, err := r.db.Database().Collection(packAccountsCollection).InsertOne(ctx, bson.M{"account_id": accountID, "name": pack.Name, "count_task": pack.CountTask, "approved": pack.Approved, "date": pack.Date})
	return err
}

func (r *StaffRepo) UpgradePack(ctx context.Context, packID primitive.ObjectID, pack domain.AccountTable) error {
	_, err := r.db.Database().Collection(packAccountsCollection).UpdateOne(ctx, bson.M{"_id": packID}, bson.M{"$set": bson.M{"name": pack.Name, "count_task": pack.CountTask}})
	return err
}

func (r *StaffRepo) DeleteAccount(ctx context.Context, accountID primitive.ObjectID) error {
	_, err := r.db.DeleteOne(ctx, bson.M{"_id": accountID})
	if err != nil {
		return err
	}
	_, err = r.db.Database().Collection(packAccountsCollection).DeleteOne(ctx, bson.M{"account_id": accountID})
	return err
}