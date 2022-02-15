package repository

import (
	"context"

	"github.com/Traking-work/traking-backend.git/internal/domain"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Authorization interface {
	GetUser(ctx context.Context, username, password string) (domain.UserLogin, error)
	SetSession(ctx context.Context, userID primitive.ObjectID, session domain.Session) error
	GetByRefreshToken(ctx context.Context, refreshToken string) (domain.UserLogin, error)
	RemoveRefreshToken(ctx context.Context, refreshToken string) error
}

type Admin interface {
	GetTeamLeads(ctx context.Context) ([]domain.UserData, error)
	AddUser(ctx context.Context, inp domain.UserData) error
}

type Teamlead interface{
	GetData(ctx context.Context, userID primitive.ObjectID) (domain.UserData, error)
	GetStaff(ctx context.Context, userID primitive.ObjectID) ([]domain.UserData, error)
}

type Staff interface {
	GetAccounts(ctx context.Context, userID primitive.ObjectID) ([]domain.AccountData, error)
	AddAccount(ctx context.Context, account domain.NewAccount, userID primitive.ObjectID) error
	GetDataAccount(ctx context.Context, accountID primitive.ObjectID) ([]domain.AccountTable, error)
}

type Repository struct {
	Authorization
	Admin
	Teamlead
	Staff
}

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{
		Authorization: NewAuthorizationRepo(db.Database(viper.GetString("mongo.databaseName"))),
		Admin: 		   NewAdminRepo(db.Database(viper.GetString("mongo.databaseName"))),
		Teamlead:      NewTeamleadRepo(db.Database(viper.GetString("mongo.databaseName"))),
		Staff: 		   NewStaffRepo(db.Database(viper.GetString("mongo.databaseName"))),
	}
}
