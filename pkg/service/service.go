package service

import (
	"context"

	"github.com/Traking-work/traking-backend.git/internal/domain"
	"github.com/Traking-work/traking-backend.git/pkg/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userData struct {
	AccessToken  string
	RefreshToken string
	UserID       string
}

type Admins interface {
	Login(ctx context.Context, username, password string) (userData, error)
	Refresh(ctx context.Context, refreshToken string) (userData, error)
	Logout(ctx context.Context, refreshToken string) error
	ParseToken(token string) (string, error)
}

type Service struct {
	Admins
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Admins: NewAdminsService(repos.Admins),
	}
}
