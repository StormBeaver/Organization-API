package service

import (
	"context"
	"fmt"
	appErrors "orgService/internal/errors"
	"orgService/internal/handlers/dto"
	"orgService/internal/model"
	"time"
)

func (s *Service) CreateEmployee(ctx context.Context, request dto.CreateEmployeeRequest) (*model.Employee, error) {
	err := s.validateCreateEmployee(ctx, request)
	if err != nil {
		return nil, err
	}

	department, err := s.repo.GetDepartment(ctx, request.DepartmentID)
	if err != nil {
		return nil, fmt.Errorf("department not found: %w", appErrors.ErrInvalidDepartmentNumber)
	}

	employee := model.Employee{
		Department: department,
		FullName:   request.FullName,
		Position:   request.Position,
		HiredAt:    request.HiredAt,
	}

	return s.repo.CreateEmployee(ctx, &employee)
}

func (s *Service) validateCreateEmployee(ctx context.Context, request dto.CreateEmployeeRequest) error {
	err := validateFieldLength(request.FullName)
	if err != nil {
		s.logger.Err(appErrors.ErrInvalidFieldLength).Msg("invalid name length")
		return fmt.Errorf("name invalid: %w", appErrors.ErrInvalidFieldLength)
	}

	err = validateFieldLength(request.Position)
	if err != nil {
		s.logger.Err(appErrors.ErrInvalidFieldLength).Msg("invalid position length")
		return fmt.Errorf("position invalid: %w", appErrors.ErrInvalidFieldLength)
	}

	if request.HiredAt != nil && time.Now().Before(*request.HiredAt) {
		return appErrors.ErrInvalidTime
	}

	return nil
}
