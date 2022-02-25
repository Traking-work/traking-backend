package repository

import (
	"context"
	"time"

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
	GetTeamLeads(ctx context.Context) ([]domain.UserDataAccount, error)
	GetCountWorkers(ctx context.Context, userID primitive.ObjectID) (int, error)
	GetWorkers(ctx context.Context, userID primitive.ObjectID) ([]domain.UserDataAccount, error)
	AddUser(ctx context.Context, inp domain.UserData) error
	DeleteUser(ctx context.Context, userID primitive.ObjectID) error
}

type Teamlead interface {
	GetData(ctx context.Context, userID primitive.ObjectID) (domain.UserData, error)
	GetStaff(ctx context.Context, userID primitive.ObjectID) ([]domain.UserDataAccount, error)
}

type Staff interface {
	GetDataUser(ctx context.Context, userID primitive.ObjectID) (domain.UserDataAccount, error)
	GetAccounts(ctx context.Context, userID primitive.ObjectID, date time.Time) ([]domain.AccountData, error)
	GetAllAccounts(ctx context.Context, userID primitive.ObjectID) ([]domain.AccountData, error)
	AddAccount(ctx context.Context, account domain.AccountData) error
	GetPacksAccount(ctx context.Context, accountID primitive.ObjectID, date string) ([]domain.AccountPack, error)
	GetDataAccount(ctx context.Context, accountID primitive.ObjectID) (domain.AccountData, error)
	AddPack(ctx context.Context, accountID primitive.ObjectID, pack domain.AccountPack) error
	UpgradePack(ctx context.Context, packID primitive.ObjectID, pack domain.AccountPack) error
	ApprovePack(ctx context.Context, packID primitive.ObjectID) error
	DeleteAccount(ctx context.Context, accountID primitive.ObjectID) error
	GetStaff(ctx context.Context, userID primitive.ObjectID) ([]domain.UserDataAccount, error)
	GetTeamLeads(ctx context.Context) ([]domain.UserDataAccount, error)
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
		Admin:         NewAdminRepo(db.Database(viper.GetString("mongo.databaseName"))),
		Teamlead:      NewTeamleadRepo(db.Database(viper.GetString("mongo.databaseName"))),
		Staff:         NewStaffRepo(db.Database(viper.GetString("mongo.databaseName"))),
	}
}
