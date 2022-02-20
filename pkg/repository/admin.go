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

func (r *AdminRepo) GetWorkers(ctx context.Context, userID primitive.ObjectID) ([]domain.UserData, error) {
	var workers []domain.UserData

	cur, err := r.db.Find(ctx, bson.M{"teamlead": userID})
	if err != nil {
		return nil, err
	}

	if err := cur.All(ctx, &workers); err != nil {
		return nil, err
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