package service

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/Traking-work/traking-backend.git/pkg/repository"
	"github.com/Traking-work/traking-backend.git/internal/domain"
	"github.com/Traking-work/traking-backend.git/pkg/logging"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StaffService struct {
	repo repository.Staff
}

func NewStaffService(repo repository.Staff) *StaffService {
	logger := logging.GetLogger()
	if err := godotenv.Load(); err != nil {
		logger.Fatalf("error loading env variables: %s", err.Error())
	}

	return &StaffService{repo: repo}
}

func (s *StaffService) GetAccounts(ctx context.Context, userID primitive.ObjectID) ([]domain.AccountData, error) {
	accounts, err := s.repo.GetAccounts(ctx, userID)
	return accounts, err
}

func (s *StaffService) AddAccount(ctx context.Context, account domain.NewAccount) error {
	err := s.repo.AddAccount(ctx, account)
	return err
}

func (s *StaffService) GetDataAccount(ctx context.Context, accountID primitive.ObjectID) ([]domain.AccountTable, error) {
	dataAccount, err := s.repo.GetDataAccount(ctx, accountID)
	return dataAccount, err
}

func (s *StaffService) AddPack(ctx context.Context, accountID primitive.ObjectID, pack domain.AccountTable) error {
	err := s.repo.AddPack(ctx, accountID, pack)
	return err
}