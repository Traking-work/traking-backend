package service

import (
	"context"

	"github.com/Traking-work/traking-backend.git/pkg/repository"
	"github.com/Traking-work/traking-backend.git/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userData struct {
	AccessToken  string
	Position     string
	RefreshToken string
	UserID       string
}

type Authorization interface {
	Login(ctx context.Context, username, password string) (userData, error)
	Refresh(ctx context.Context, refreshToken string) (userData, error)
	Logout(ctx context.Context, refreshToken string) error
	ParseToken(token string) (string, error)
}

type Admin interface {
	GetTeamLeads(ctx context.Context) ([]domain.UserData, error)
	AddUser(ctx context.Context, inp domain.UserData) error
}

type Teamlead interface {
	GetStaff(ctx context.Context, userID primitive.ObjectID) ([]domain.UserData, error)
}

type Staff interface {

}

type Service struct {
	Authorization
	Admin
	Teamlead
	Staff
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthorizationService(repos.Authorization),
		Admin: 		   NewAdminService(repos.Admin),
		Teamlead: 	   NewTeamleadService(repos.Teamlead),
		Staff: 		   NewStaffService(repos.Staff),
	}
}
