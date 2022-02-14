package repository

import (
	"context"

	"github.com/Traking-work/traking-backend.git/internal/domain"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Admins interface {
	GetUser(ctx context.Context, username, password string) (domain.User, error)
	SetSession(ctx context.Context, userID primitive.ObjectID, session domain.Session) error
	GetByRefreshToken(ctx context.Context, refreshToken string) (domain.User, error)
	RemoveRefreshToken(ctx context.Context, refreshToken string) error
}

type Repository struct {
	Admins
}

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{
		Admins: NewAdminsRepo(db.Database(viper.GetString("mongo.databaseName"))),
	}
}
