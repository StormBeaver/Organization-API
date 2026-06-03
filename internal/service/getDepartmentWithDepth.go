package service

import (
	"context"
	appErrors "orgService/internal/errors"
	"orgService/internal/handlers/dto"
	"orgService/internal/model"
)

func (s *Service) GetDepartmentWithDepth(ctx context.Context, id int, request dto.GetDepartmentRequest) (*model.Department, error) {
	if request.Depth > 5 || request.Depth < 1 {
		return nil, appErrors.ErrInvalidDepth
	}

	hint := &model.GetParams{Depth: request.Depth, IncludeEmployees: request.IncludeEmployees}

	return s.repo.GetDepartmentWithDepth(ctx, id, hint)
}
