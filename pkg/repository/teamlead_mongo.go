package repository

import (
	"context"
	"errors"
	"time"

	"github.com/Traking-work/traking-backend.git/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func (r *TeamleadRepo) GetCountAccounts(ctx context.Context, userID primitive.ObjectID) (int, int, error) {
	var accounts []domain.AccountData
	countAccounts := 0
	countAccountsActive := 0

	cur, err := r.db.Database().Collection(accountsCollection).Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return 0, 0, err
	}

	if err := cur.All(ctx, &accounts); err != nil {
		return 0, 0, err
	}

	for _, account := range accounts {
		countAccounts++

		if account.CreateDate.Month() == time.Now().Month() {
			if account.CreateDate.Day() <= time.Now().Day() {
				if account.StatusDelete {
					if account.DeleteDate.Month() == time.Now().Month() {
						if account.DeleteDate.Day() > time.Now().Day() {
							countAccountsActive++
						}
					} else if account.DeleteDate.Month() > time.Now().Month() {
						countAccountsActive++
					}
				} else {
					countAccountsActive++
				}
			}
		} else if account.CreateDate.Month() < time.Now().Month() {
			if account.StatusDelete {
				if account.DeleteDate.Month() == time.Now().Month() {
					if account.DeleteDate.Day() > time.Now().Day() {
						countAccountsActive++
					}
				} else if account.DeleteDate.Month() > time.Now().Month() {
					countAccountsActive++
				}
			} else {
				countAccountsActive++
			}
		}
	}

	return countAccounts, countAccountsActive, nil
}

func (r *TeamleadRepo) GetStaff(ctx context.Context, userID primitive.ObjectID) ([]domain.UserDataAccount, error) {
	var staff []domain.UserDataAccount

	cur, err := r.db.Find(ctx, bson.M{"position": "staff", "teamlead": userID})
	if err != nil {
		return nil, err
	}

	if err := cur.All(ctx, &staff); err != nil {
		return nil, err
	}

	for index, i := range staff {
		countAccounts, countAccountsActive, err := r.GetCountAccounts(ctx, i.ID)
		if err == nil {
			staff[index].CountEmployee = countAccounts
			staff[index].CountEmployeeActive = countAccountsActive
		}
	}

	return staff, nil
}
