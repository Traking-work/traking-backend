package service

import (
	"context"
	"time"

	"github.com/Traking-work/traking-backend.git/internal/domain"
	"github.com/Traking-work/traking-backend.git/pkg/repository"
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
	GetTeamLeads(ctx context.Context) ([]domain.UserDataAccount, []domain.UserSelect, error)
	GetCountWorkers(ctx context.Context, userID primitive.ObjectID) (int, error)
	GetWorkers(ctx context.Context, userID primitive.ObjectID) ([]domain.UserDataAccount, error)
	AddUser(ctx context.Context, inp domain.UserData) error
	DeleteUser(ctx context.Context, userID primitive.ObjectID) error
}

type Teamlead interface {
	GetStaff(ctx context.Context, userID primitive.ObjectID) ([]domain.UserDataAccount, error)
}

type Staff interface {
	GetDataUser(ctx context.Context, userID primitive.ObjectID) (domain.UserDataAccount, error)
	GetAccounts(ctx context.Context, userID primitive.ObjectID, date time.Time) ([]domain.AccountData, error)
	AddAccount(ctx context.Context, account domain.AccountData) error
	GetPacksAccount(ctx context.Context, accountID primitive.ObjectID, date string) ([]domain.AccountPack, error)
	GetDataAccount(ctx context.Context, accountID primitive.ObjectID) (domain.AccountData, error)
	AddPack(ctx context.Context, accountID primitive.ObjectID, pack domain.AccountPack) error
	UpgradePack(ctx context.Context, packID primitive.ObjectID, pack domain.AccountPack) error
	ApprovePack(ctx context.Context, packID primitive.ObjectID) error
	DeleteAccount(ctx context.Context, accountID primitive.ObjectID) error
	GetParamsMainStaff(ctx context.Context, userID primitive.ObjectID) (float64, error)
	GetParamsDateStaff(ctx context.Context, userID primitive.ObjectID, date string) (float64, error)
	GetParamsMainTeamlead(ctx context.Context, userID primitive.ObjectID) (float64, error)
	GetParamsDateTeamlead(ctx context.Context, userID primitive.ObjectID, date string) (float64, error)
	GetParamsMainAdmin(ctx context.Context) (float64, error)
	GetParamsDateAdmin(ctx context.Context, date string) (float64, error)
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
		Admin:         NewAdminService(repos.Admin),
		Teamlead:      NewTeamleadService(repos.Teamlead),
		Staff:         NewStaffService(repos.Staff),
	}
}
