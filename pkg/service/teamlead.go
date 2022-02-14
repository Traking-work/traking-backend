package service

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/Traking-work/traking-backend.git/pkg/repository"
	"github.com/Traking-work/traking-backend.git/internal/domain"
	"github.com/Traking-work/traking-backend.git/pkg/logging"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TeamleadService struct {
	repo repository.Teamlead
}

func NewTeamleadService(repo repository.Teamlead) *TeamleadService {
	logger := logging.GetLogger()
	if err := godotenv.Load(); err != nil {
		logger.Fatalf("error loading env variables: %s", err.Error())
	}

	return &TeamleadService{repo: repo}
}

func (s *TeamleadService) GetStaff(ctx context.Context, userID primitive.ObjectID) ([]domain.UserData, error) {
	staff, err := s.repo.GetStaff(ctx, userID)
	return staff, err
}