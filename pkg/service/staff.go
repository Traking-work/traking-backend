package service

import (
	"context"
	"strings"
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
	packsAccount, err := s.repo.GetPacksAccount(ctx, accountID, date, "")
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

func (s *StaffService) ChangeTeamlead(ctx context.Context, userID primitive.ObjectID, teamleadID primitive.ObjectID) error {
	err := s.repo.ChangeTeamlead(ctx, userID, teamleadID)
	return err
}

func (s *StaffService) GetIncomeStaff(ctx context.Context, userID primitive.ObjectID, fromDate string, toDate string) (map[string]float32, error) {
	incomeDict := make(map[string]float32)

	accounts, err := s.repo.GetAllAccounts(ctx, userID)
	if err != nil {
		return nil, err
	}

	for _, account := range accounts {
		packsAccount, err := s.repo.GetPacksAccount(ctx, account.ID, fromDate, toDate)
		if err != nil {
			return nil, err
		}

		for _, pack := range packsAccount {
			dateList := strings.Split(pack.Date, "-")
			incomeDict[dateList[1]+"."+dateList[2]] += float32(pack.CountTask) * float32(pack.Payment)
		}
	}

	return incomeDict, nil
}

func (s *StaffService) GetIncomeTeamlead(ctx context.Context, userID primitive.ObjectID, fromDate string, toDate string) (map[string]float32, error) {
	incomeDict := make(map[string]float32)

	staff, err := s.repo.GetStaff(ctx, userID)
	if err != nil {
		return nil, err
	}

	for _, st := range staff {
		incomeStaff, err := s.GetIncomeStaff(ctx, st.ID, fromDate, toDate)
		if err != nil {
			return nil, err
		}

		for data, income := range incomeStaff {
			incomeDict[data] += income
		}
	}

	return incomeDict, nil
}

func (s *StaffService) GetIncomeAdmin(ctx context.Context, userID primitive.ObjectID, fromDate string, toDate string) (map[string]float32, error) {
	incomeDict := make(map[string]float32)

	teamleads, err := s.repo.GetTeamLeads(ctx)
	if err != nil {
		return nil, err
	}

	for _, teamlead := range teamleads {
		incomeTeamlead, err := s.GetIncomeTeamlead(ctx, teamlead.ID, fromDate, toDate)
		if err != nil {
			return nil, err
		}
		for data, income := range incomeTeamlead {
			incomeDict[data] += income
		}
	}

	staff, err := s.repo.GetStaff(ctx, userID)
	if err != nil {
		return nil, err
	}

	for _, st := range staff {
		incomeStaff, err := s.GetIncomeStaff(ctx, st.ID, fromDate, toDate)
		if err != nil {
			return nil, err
		}
		for data, income := range incomeStaff {
			incomeDict[data] += income
		}
	}

	return incomeDict, nil
}
