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

func (s *AdminService) AddUser(ctx context.Context, inp domain.NewUser) error {
	inp.Password = generateHash(inp.Password)

	err := s.repo.AddUser(ctx, inp)
	return err
}

func generateHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("SALT"))))
}