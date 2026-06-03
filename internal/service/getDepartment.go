package service

import (
	"context"
	"orgService/internal/model"
)

func (s *Service) GetDepartment(ctx context.Context, id int) (*model.Department, error) {

	return s.repo.GetDepartment(ctx, id)
}
