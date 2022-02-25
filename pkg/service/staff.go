package service

import (
	"context"
	"time"

	"github.com/Traking-work/traking-backend.git/internal/domain"
	"github.com/Traking-work/traking-backend.git/pkg/logging"
	"github.com/Traking-work/traking-backend.git/pkg/repository"
	"github.com/joho/godotenv"
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

func (s *StaffService) GetAccounts(ctx context.Context, userID primitive.ObjectID, date time.Time) ([]domain.AccountData, error) {
	accounts, err := s.repo.GetAccounts(ctx, userID, date)
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

func (s *StaffService) GetParamsMainStaff(ctx context.Context, userID primitive.ObjectID) (float64, error) {
	income := 0.0
	accounts, err := s.repo.GetAllAccounts(ctx, userID)
	if err != nil {
		return 0, err
	}

	for _, account := range accounts {
		packsAccount, err := s.repo.GetPacksAccount(ctx, account.ID, "")
		if err != nil {
			return 0, err
		}

		for _, pack := range packsAccount {
			income += float64(pack.CountTask) * float64(pack.Payment)
		}
	}

	return income, nil
}

func (s *StaffService) GetParamsDateStaff(ctx context.Context, userID primitive.ObjectID, date string) (float64, error) {
	income := 0.0
	accounts, err := s.repo.GetAllAccounts(ctx, userID)
	if err != nil {
		return 0, err
	}

	for _, account := range accounts {
		packsAccount, err := s.repo.GetPacksAccount(ctx, account.ID, date)
		if err != nil {
			return 0, err
		}

		for _, pack := range packsAccount {
			income += float64(pack.CountTask) * float64(pack.Payment)
		}
	}

	return income, nil
}

func (s *StaffService) GetParamsMainTeamlead(ctx context.Context, userID primitive.ObjectID) (float64, error) {
	income := 0.0

	staff, err := s.repo.GetStaff(ctx, userID)
	if err != nil {
		return 0, err
	}

	for _, st := range staff {
		income_st, err := s.GetParamsMainStaff(ctx, st.ID)
		if err != nil {
			return 0, err
		}
		income += income_st
	}

	return income, nil
}

func (s *StaffService) GetParamsDateTeamlead(ctx context.Context, userID primitive.ObjectID, date string) (float64, error) {
	income := 0.0

	staff, err := s.repo.GetStaff(ctx, userID)
	if err != nil {
		return 0, err
	}

	for _, st := range staff {
		income_st, err := s.GetParamsDateStaff(ctx, st.ID, date)
		if err != nil {
			return 0, err
		}
		income += income_st
	}

	return income, nil
}

func (s *StaffService) GetParamsMainAdmin(ctx context.Context) (float64, error) {
	income := 0.0

	teamleads, err := s.repo.GetTeamLeads(ctx)
	if err != nil {
		return 0, err
	}

	for _, teamlead := range teamleads {
		income_st, err := s.GetParamsMainTeamlead(ctx, teamlead.ID)
		if err != nil {
			return 0, err
		}
		income += income_st
	}

	return income, nil
}

func (s *StaffService) GetParamsDateAdmin(ctx context.Context, date string) (float64, error) {
	income := 0.0

	teamleads, err := s.repo.GetTeamLeads(ctx)
	if err != nil {
		return 0, err
	}

	for _, teamlead := range teamleads {
		income_st, err := s.GetParamsDateTeamlead(ctx, teamlead.ID, date)
		if err != nil {
			return 0, err
		}
		income += income_st
	}

	return income, nil
}
