package service

import (
	"context"
	"crypto/sha1"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/Traking-work/traking-backend.git/pkg/repository"
	"github.com/Traking-work/traking-backend.git/internal/domain"
	"github.com/Traking-work/traking-backend.git/pkg/logging"
)

type AdminService struct {
	repo repository.Admin
}

func NewAdminService(repo repository.Admin) *AdminService {
	logger := logging.GetLogger()
	if err := godotenv.Load(); err != nil {
		logger.Fatalf("error loading env variables: %s", err.Error())
	}

	return &AdminService{repo: repo}
}

func (s *AdminService) GetTeamLeads(ctx context.Context) ([]domain.UserData, []domain.UserSelect, error) {
	teamleads, err := s.repo.GetTeamLeads(ctx)
	if err != nil {
		return nil, nil, err
	}

	teamleadsCreate := []domain.UserSelect{}

	for _, teamlead := range teamleads {
		teamleadsCreate = append(teamleadsCreate, domain.UserSelect{teamlead.ID, teamlead.Name})
	}

	return teamleads, teamleadsCreate, nil
}

func (s *AdminService) AddUser(ctx context.Context, inp domain.UserData) error {
	inp.Password = generateHash(inp.Password)

	err := s.repo.AddUser(ctx, inp)
	return err
}

func generateHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("SALT"))))
}