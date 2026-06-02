package service

import (
	"context"
	"orgService/internal/model"
	"time"
)

func (s *Service) CreateEmployee(ctx context.Context, name string, position string, depID int, hiredAt *time.Time) (model.Employee, error) {
	err := s.validateCreateEmployee(ctx, name, position, depID, hiredAt)
	if err != nil {
		return model.Employee{}, err
	}

	return s.repo.CreateEmployee(ctx, name, position, depID, hiredAt)
}
