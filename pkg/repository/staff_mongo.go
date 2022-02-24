package repository

import (
	"context"
	"time"

	"github.com/Traking-work/traking-backend.git/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type StaffRepo struct {
	db *mongo.Collection
}

func NewStaffRepo(db *mongo.Database) *StaffRepo {
	return &StaffRepo{db: db.Collection(accountsCollection)}
}

func (r *StaffRepo) GetDataUser(ctx context.Context, userID primitive.ObjectID) (domain.UserDataAccount, error) {
	var dataUserAcc domain.UserDataAccount
	var dataUser domain.UserData

	err := r.db.Database().Collection(usersCollection).FindOne(ctx, bson.M{"_id": userID}).Decode(&dataUser)
	if err != nil {
		return domain.UserDataAccount{}, err
	}
	dataUserAcc.ID = dataUser.ID
	dataUserAcc.Name = dataUser.Name
	dataUserAcc.Username = dataUser.Username

	if dataUser.Position == "staff" {
		err = r.db.Database().Collection(usersCollection).FindOne(ctx, bson.M{"_id": dataUser.TeamLead}).Decode(&dataUser)
		if err != nil {
			return domain.UserDataAccount{}, err
		}
		dataUserAcc.Director = dataUser.Name
	}

	return dataUserAcc, nil
}

func (r *StaffRepo) GetAccounts(ctx context.Context, userID primitive.ObjectID, date time.Time) ([]domain.AccountData, error) {
	var accounts []domain.AccountData
	var accountsDate []domain.AccountData

	cur, err := r.db.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}

	if err := cur.All(ctx, &accounts); err != nil {
		return nil, err
	}

	for _, account := range accounts {
		if account.CreateDate.Day() <= date.Day() {
			if account.StatusDelete {
				if account.DeleteDate.Day() > date.Day() {
					accountsDate = append(accountsDate, account)
				}
			} else {
				accountsDate = append(accountsDate, account)
			}
		}
	}

	return accountsDate, nil
}

func (r *StaffRepo) AddAccount(ctx context.Context, account domain.AccountData) error {
	_, err := r.db.InsertOne(ctx, account)
	return err
}

func (r *StaffRepo) GetPacksAccount(ctx context.Context, accountID primitive.ObjectID, date string) ([]domain.AccountPack, error) {
	var packsAccount []domain.AccountPack

	cur, err := r.db.Database().Collection(packAccountsCollection).Find(ctx, bson.M{"account_id": accountID, "date": date})
	if err != nil {
		return nil, err
	}

	if err := cur.All(ctx, &packsAccount); err != nil {
		return nil, err
	}

	return packsAccount, nil
}

func (r *StaffRepo) GetDataAccount(ctx context.Context, accountID primitive.ObjectID) (domain.AccountData, error) {
	var dataAccount domain.AccountData

	err := r.db.FindOne(ctx, bson.M{"_id": accountID}).Decode(&dataAccount)
	return dataAccount, err
}

func (r *StaffRepo) AddPack(ctx context.Context, accountID primitive.ObjectID, pack domain.AccountPack) error {
	_, err := r.db.Database().Collection(packAccountsCollection).InsertOne(ctx, bson.M{"account_id": accountID, "name": pack.Name, "count_task": pack.CountTask, "payment": pack.Payment, "approved": pack.Approved, "date": pack.Date})
	return err
}

func (r *StaffRepo) UpgradePack(ctx context.Context, packID primitive.ObjectID, pack domain.AccountPack) error {
	_, err := r.db.Database().Collection(packAccountsCollection).UpdateOne(ctx, bson.M{"_id": packID}, bson.M{"$set": bson.M{"name": pack.Name, "count_task": pack.CountTask, "payment": pack.Payment}})
	return err
}

func (r *StaffRepo) ApprovePack(ctx context.Context, packID primitive.ObjectID) error {
	_, err := r.db.Database().Collection(packAccountsCollection).UpdateOne(ctx, bson.M{"_id": packID}, bson.M{"$set": bson.M{"approved": true}})
	return err
}

func (r *StaffRepo) DeleteAccount(ctx context.Context, accountID primitive.ObjectID) error {
	//_, err := r.db.DeleteOne(ctx, bson.M{"_id": accountID})
	//if err != nil {
	//	return err
	//}

	//_, err = r.db.Database().Collection(packAccountsCollection).DeleteOne(ctx, bson.M{"account_id": accountID})
	//return err

	_, err := r.db.UpdateOne(ctx, bson.M{"_id": accountID}, bson.M{"$set": bson.M{"status_delete": true, "delete_date": time.Now()}})
	return err
}
