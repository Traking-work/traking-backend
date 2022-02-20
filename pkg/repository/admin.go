package repository

import (
	"context"
	"errors"

	"github.com/Traking-work/traking-backend.git/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AdminRepo struct {
	db *mongo.Collection
}

func NewAdminRepo(db *mongo.Database) *AdminRepo {
	return &AdminRepo{db: db.Collection(usersCollection)}
}

func (r *AdminRepo) GetCountWorkers(ctx context.Context, userID primitive.ObjectID) (int, error) {
	var workers []domain.UserData
	countWorkers := 0

	cur, err := r.db.Find(ctx, bson.M{"teamlead": userID})
	if err != nil {
		return 0, err
	}

	if err := cur.All(ctx, &workers); err != nil {
		return 0, err
	}

	for _ = range workers {
		countWorkers++
	}

	return countWorkers, nil
}

func (r *AdminRepo) GetCountStaff(ctx context.Context, userID primitive.ObjectID) (int, error) {
	var teamleads []domain.UserDataAccount
	countWorkers := 0

	cur, err := r.db.Find(ctx, bson.M{"teamlead": userID})
	if err != nil {
		return 0, err
	}

	if err := cur.All(ctx, &teamleads); err != nil {
		return 0, err
	}

	for _ = range teamleads {
		countWorkers++
	}

	return countWorkers, nil
}

func (r *AdminRepo) GetTeamLeads(ctx context.Context) ([]domain.UserDataAccount, error) {
	var teamleads []domain.UserDataAccount

	cur, err := r.db.Find(ctx, bson.M{"position": "teamlead"})
	if err != nil {
		return nil, err
	}

	if err := cur.All(ctx, &teamleads); err != nil {
		return nil, err
	}

	for index, i := range teamleads {
		countWorkers, err := r.GetCountStaff(ctx, i.ID)
		if err == nil {
			teamleads[index].CountEmployee = countWorkers
		}
	}

	return teamleads, nil
}

func (r *AdminRepo) GetCountAccounts(ctx context.Context, userID primitive.ObjectID) (int, error) {
	var accounts []domain.UserDataAccount
	countAccounts := 0

	cur, err := r.db.Database().Collection(accountsCollection).Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return 0, err
	}

	if err := cur.All(ctx, &accounts); err != nil {
		return 0, err
	}

	for _ = range accounts {
		countAccounts++
	}

	return countAccounts, nil
}

func (r *AdminRepo) GetWorkers(ctx context.Context, userID primitive.ObjectID) ([]domain.UserDataAccount, error) {
	var workers []domain.UserDataAccount

	cur, err := r.db.Find(ctx, bson.M{"teamlead": userID})
	if err != nil {
		return nil, err
	}

	if err := cur.All(ctx, &workers); err != nil {
		return nil, err
	}

	for index, i := range workers {
		countAccounts, err := r.GetCountAccounts(ctx, i.ID)
		if err == nil {
			workers[index].CountEmployee = countAccounts
		}
	}

	return workers, nil
}

func (r *AdminRepo) CheckUsername(ctx context.Context, username string) (bool, error) {
	var userData domain.UserDataAccount

	err := r.db.FindOne(ctx, bson.M{"username": username}).Decode(&userData)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return true, nil
		}
		return false, err
	}
	
	return false, nil
}

func (r *AdminRepo) AddUser(ctx context.Context, inp domain.UserData) error {
	var err error

	resultCheck, err := r.CheckUsername(ctx, inp.Username)
	if err != nil {
		return err
	}

	if resultCheck {
		if inp.Position == "staff" {
			_, err = r.db.InsertOne(ctx, inp)
		} else {
			_, err = r.db.InsertOne(ctx, bson.M{"name": inp.Name, "username": inp.Username, "password": inp.Password, "position": inp.Position})
		}
		return err
	} else {
		return domain.ErrReplayUsername
	}
}

func (r *AdminRepo) DeleteUser(ctx context.Context, userID primitive.ObjectID) error {
	_, err := r.db.DeleteOne(ctx, bson.M{"_id": userID})
	if err != nil {
		return err
	}

	_, err = r.db.DeleteOne(ctx, bson.M{"teamlead": userID})
	if err != nil {
		return err
	}

	var accountsDel []domain.AccountData

	cur, err := r.db.Database().Collection(accountsCollection).Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return err
	}

	if err := cur.All(ctx, &accountsDel); err != nil {
		return err
	}

	for _, i := range accountsDel {
		_, err = r.db.Database().Collection(packAccountsCollection).DeleteOne(ctx, bson.M{"account_id": i.ID})
		if err != nil {
			return err
		}

		_, err = r.db.Database().Collection(accountsCollection).DeleteOne(ctx, bson.M{"_id": i.ID})
		if err != nil {
			return err
		}
	}

	return nil
}