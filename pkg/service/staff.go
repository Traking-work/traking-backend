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

func (s *StaffService) GetDataUser(ctx context.Context, userID primitive.ObjectID) (domain.UserDataAccount, error) {
	dataUser, err := s.repo.GetDataUser(ctx, userID)
	return dataUser, err
}

func (s *StaffService) GetAccounts(ctx context.Context, userID primitive.ObjectID) ([]domain.AccountData, error) {
	accounts, err := s.repo.GetAccounts(ctx, userID)
	return accounts, err
}

func (s *StaffService) AddAccount(ctx context.Context, account domain.AccountData) error {
	err := s.repo.AddAccount(ctx, account)
	return err
}

func (s *StaffService) GetPacksAccount(ctx context.Context, accountID primitive.ObjectID, date string) ([]domain.AccountPack, error) {
	packsAccount, err := s.repo.GetPacksAccount(ctx, accountID, date)
	return packsAccount, err
}

func (s *StaffService) GetDataAccount(ctx context.Context, accountID primitive.ObjectID) (domain.AccountData, error) {
	dataAccount, err := s.repo.GetDataAccount(ctx, accountID)
	return dataAccount, err
}

func (s *StaffService) AddPack(ctx context.Context, accountID primitive.ObjectID, pack domain.AccountPack) error {
	err := s.repo.AddPack(ctx, accountID, pack)
	return err
}

func (s *StaffService) UpgradePack(ctx context.Context, packID primitive.ObjectID, pack domain.AccountPack) error {
	err := s.repo.UpgradePack(ctx, packID, pack)
	return err
}

func (s *StaffService) ApprovePack(ctx context.Context, packID primitive.ObjectID) error {
	err := s.repo.ApprovePack(ctx, packID)
	return err
}

func (s *StaffService) DeleteAccount(ctx context.Context, accountID primitive.ObjectID) error {
	err := s.repo.DeleteAccount(ctx, accountID)
	return err
}