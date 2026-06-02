package service

import (
	"context"
	"orgService/internal/model"
	"time"

	"github.com/rs/zerolog"
)

type Repo interface {
	CreateDepartment(ctx context.Context, name string, parentID *int) (model.Department, error)
	CreateEmployee(ctx context.Context, name string, position string, depID int, hiredAt *time.Time) (model.Employee, error)
	GetDepartment(ctx context.Context, id, depth int, employees bool) (model.Department, *[]model.Employee, error)
	PatchDepartment(ctx context.Context, id int, name *string, parentID *int) (model.Department, error)
	DeleteDepartment(ctx context.Context, id int, mode string, reasignDestination *int) error

	LastDepartment(ctx context.Context) (model.Department, error)
}

type Service struct {
	repo Repo

	logger *zerolog.Logger
}

func NewService(repo Repo, logger *zerolog.Logger) *Service {
	return &Service{repo: repo, logger: logger}
}

func (s *Service) GetDepartment(ctx context.Context, id, depth int, employees bool) (model.Department, *[]model.Employee, error) {
	return s.repo.GetDepartment(ctx, id, depth, employees)
}

func (s *Service) DeleteDepartment(ctx context.Context, ID int, mode string, reasignDestination *int) error {
	return s.repo.DeleteDepartment(ctx, ID, mode, reasignDestination)
}

func (s *Service) LastDepartment(ctx context.Context) (model.Department, error) {
	return s.repo.LastDepartment(ctx)
}
