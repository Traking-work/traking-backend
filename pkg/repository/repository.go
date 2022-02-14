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

type Repository struct {
	Authorization
	Admin
}

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{
		Authorization: NewAuthorizationRepo(db.Database(viper.GetString("mongo.databaseName"))),
		Admin: 		   NewAdminRepo(db.Database(viper.GetString("mongo.databaseName"))),
	}
}
