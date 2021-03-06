package service

import (
	"context"
	"crypto/sha1"
	"fmt"
	"os"

	"github.com/Traking-work/traking-backend.git/internal/domain"
	"github.com/Traking-work/traking-backend.git/pkg/logging"
	"github.com/Traking-work/traking-backend.git/pkg/repository"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (s *AdminService) GetTeamLeads(ctx context.Context) ([]domain.UserDataAccount, []domain.UserSelect, error) {
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

func (s *AdminService) GetCountWorkers(ctx context.Context, userID primitive.ObjectID) (int, error) {
	countWorkers, err := s.repo.GetCountWorkers(ctx, userID)
	return countWorkers, err
}

func (s *AdminService) GetWorkers(ctx context.Context, userID primitive.ObjectID) ([]domain.UserDataAccount, error) {
	workers, err := s.repo.GetWorkers(ctx, userID)
	return workers, err
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

func (s *AdminService) DeleteUser(ctx context.Context, userID primitive.ObjectID, position string) error {
	err := s.repo.DeleteUser(ctx, userID, position)
	return err
}
