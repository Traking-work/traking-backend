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

func (s *StaffService) DeletePack(ctx context.Context, packID primitive.ObjectID) error {
	err := s.repo.DeletePack(ctx, packID)
	return err
}

func (s *StaffService) DeleteAccount(ctx context.Context, accountID primitive.ObjectID) error {
	err := s.repo.DeleteAccount(ctx, accountID)
	return err
}

func (s *StaffService) GetParamsMainStaff(ctx context.Context, userID primitive.ObjectID) (float32, float32, error) {
	incomeAll := float32(0)
	incomeAdmin := float32(0)

	accounts, err := s.repo.GetAllAccounts(ctx, userID)
	if err != nil {
		return 0, 0, err
	}

	for _, account := range accounts {
		packsAccount, err := s.repo.GetPacksAccount(ctx, account.ID, "")
		if err != nil {
			return 0, 0, err
		}

		for _, pack := range packsAccount {
			incomeAll += float32(pack.CountTask) * float32(pack.Payment)
		}
		incomeAdmin = incomeAll * float32(account.Percent)
	}

	return incomeAll, incomeAdmin, nil
}

func (s *StaffService) GetParamsDateStaff(ctx context.Context, userID primitive.ObjectID, date string) (float32, float32, error) {
	incomeAll := float32(0)
	incomeAdmin := float32(0)

	accounts, err := s.repo.GetAllAccounts(ctx, userID)
	if err != nil {
		return 0, 0, err
	}

	for _, account := range accounts {
		packsAccount, err := s.repo.GetPacksAccount(ctx, account.ID, date)
		if err != nil {
			return 0, 0, err
		}

		for _, pack := range packsAccount {
			incomeAll += float32(pack.CountTask) * float32(pack.Payment)
		}
		incomeAdmin = incomeAll * float32(account.Percent)
	}

	return incomeAll, incomeAdmin, nil
}

func (s *StaffService) GetParamsMainTeamlead(ctx context.Context, userID primitive.ObjectID) (float32, float32, error) {
	incomeAll := float32(0)
	incomeAdmin := float32(0)

	staff, err := s.repo.GetStaff(ctx, userID)
	if err != nil {
		return 0, 0, err
	}

	for _, st := range staff {
		income_st_all, income_st_admin, err := s.GetParamsMainStaff(ctx, st.ID)
		if err != nil {
			return 0, 0, err
		}
		incomeAll += income_st_all
		incomeAdmin += income_st_admin
	}

	return incomeAll, incomeAdmin, nil
}

func (s *StaffService) GetParamsDateTeamlead(ctx context.Context, userID primitive.ObjectID, date string) (float32, float32, error) {
	incomeAll := float32(0)
	incomeAdmin := float32(0)

	staff, err := s.repo.GetStaff(ctx, userID)
	if err != nil {
		return 0, 0, err
	}

	for _, st := range staff {
		income_st_all, income_st_admin, err := s.GetParamsDateStaff(ctx, st.ID, date)
		if err != nil {
			return 0, 0, err
		}
		incomeAll += income_st_all
		incomeAdmin += income_st_admin
	}

	return incomeAll, incomeAdmin, nil
}

func (s *StaffService) GetParamsMainAdmin(ctx context.Context, userID primitive.ObjectID) (float32, float32, error) {
	incomeAll := float32(0)
	incomeAdmin := float32(0)

	teamleads, err := s.repo.GetTeamLeads(ctx)
	if err != nil {
		return 0, 0, err
	}

	for _, teamlead := range teamleads {
		income_st_all, income_st_admin, err := s.GetParamsMainTeamlead(ctx, teamlead.ID)
		if err != nil {
			return 0, 0, err
		}
		incomeAll += income_st_all
		incomeAdmin += income_st_admin
	}

	staff, err := s.repo.GetStaff(ctx, userID)
	if err != nil {
		return 0, 0, err
	}

	for _, st := range staff {
		income_st_all, income_st_admin, err := s.GetParamsMainStaff(ctx, st.ID)
		if err != nil {
			return 0, 0, err
		}
		incomeAll += income_st_all
		incomeAdmin += income_st_admin
	}

	return incomeAll, incomeAdmin, nil
}

func (s *StaffService) GetParamsDateAdmin(ctx context.Context, userID primitive.ObjectID, date string) (float32, float32, error) {
	incomeAll := float32(0)
	incomeAdmin := float32(0)

	teamleads, err := s.repo.GetTeamLeads(ctx)
	if err != nil {
		return 0, 0, err
	}

	for _, teamlead := range teamleads {
		income_st_all, income_st_admin, err := s.GetParamsDateTeamlead(ctx, teamlead.ID, date)
		if err != nil {
			return 0, 0, err
		}
		incomeAll += income_st_all
		incomeAdmin += income_st_admin
	}

	staff, err := s.repo.GetStaff(ctx, userID)
	if err != nil {
		return 0, 0, err
	}

	for _, st := range staff {
		income_st_all, income_st_admin, err := s.GetParamsDateStaff(ctx, st.ID, date)
		if err != nil {
			return 0, 0, err
		}
		incomeAll += income_st_all
		incomeAdmin += income_st_admin
	}

	return incomeAll, incomeAdmin, nil
}

func (s *StaffService) ChangeTeamlead(ctx context.Context, userID primitive.ObjectID, teamleadID primitive.ObjectID) error {
	err := s.repo.ChangeTeamlead(ctx, userID, teamleadID)
	return err
}