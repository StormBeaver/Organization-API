package service

import (
	"context"
	"fmt"
	appErrors "orgService/internal/errors"
	"orgService/internal/handlers/dto"
	"orgService/internal/model"
	"strings"
)

func (s *Service) CreateDepartment(ctx context.Context, request dto.CreateDepartmentRequest) (*model.Department, error) {
	err := s.validateCreateDepartment(ctx, request)
	if err != nil {
		return nil, err
	}

	request.Name = strings.Trim(request.Name, " ")

	var parent *model.Department

	if request.ParentID != nil {
		parent, err = s.repo.GetDepartment(ctx, *request.ParentID)
		if err != nil {
			return nil, fmt.Errorf("department not found: %w", appErrors.ErrInvalidDepartmentNumber)
		}
	}

	dep := &model.Department{
		Name:   &request.Name,
		Parent: parent,
	}

	return s.repo.CreateDepartment(ctx, dep)
}

func (s *Service) validateCreateDepartment(ctx context.Context, request dto.CreateDepartmentRequest) error {
	err := validateFieldLength(request.Name)
	if err != nil {
		s.logger.Err(appErrors.ErrInvalidFieldLength).Msg("invalid name length")
		return fmt.Errorf("name invalid: %w", appErrors.ErrInvalidFieldLength)
	}

	return nil
}
