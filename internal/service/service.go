package service

import (
	"context"
	"orgService/internal/model"

	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Repo interface {
	CreateDepartment(ctx context.Context, department *model.Department) (*model.Department, error)
	CreateEmployee(ctx context.Context, employee *model.Employee) (*model.Employee, error)
	GetDepartment(ctx context.Context, id int) (*model.Department, error)
	PatchDepartment(ctx context.Context, department *model.Department) (*model.Department, error)
	DeleteDepartment(ctx context.Context, tx *gorm.DB, department *model.Department) error
	ReassignDepartment(ctx context.Context, tx *gorm.DB, src, dst int) error
	BeginTx() *gorm.DB
}

type Service struct {
	repo Repo

	logger *zerolog.Logger
}

func NewService(repo Repo, logger *zerolog.Logger) *Service {
	return &Service{repo: repo, logger: logger}
}

func (s *Service) GetDepartment(ctx context.Context, id int) (*model.Department, error) {
	return s.repo.GetDepartment(ctx, id)
}
