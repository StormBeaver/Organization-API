package service

import (
	"context"
	"fmt"
	appErrors "orgService/internal/errors"
	"orgService/internal/model"
)

func (s *Service) PatchDepartment(ctx context.Context, id int, name *string, parentID *int) (model.Department, error) {
	err := s.validatePatchDepartment(ctx, id, parentID, name)
	if err != nil {
		return model.Department{}, err
	}
	return s.repo.PatchDepartment(ctx, id, name, parentID)
}

func (s *Service) validatePatchDepartment(ctx context.Context, id int, parentID *int, name *string) error {

	if parentID == nil && name == nil {
		s.logger.Err(appErrors.ErrInvalidArguments).Msg("invalid arguments")
		return fmt.Errorf("no arguments: %w", appErrors.ErrInvalidArguments)
	}

	if name != nil {
		err := validateFieldLength(*name)

		if err != nil {
			s.logger.Err(appErrors.ErrInvalidFieldLength).Msg("name too long")
			return fmt.Errorf("name invalid: %w", appErrors.ErrInvalidFieldLength)
		}
	}

	if parentID != nil {
		_, _, err := s.GetDepartment(ctx, *parentID, 1, false)
		if err != nil {
			s.logger.Err(appErrors.ErrInvalidDepartmentNumber).Msg("parent department doesn't exist")
			return fmt.Errorf("parent department doesn't exist: %w", appErrors.ErrInvalidDepartmentNumber)
		}
	}

	_, _, err := s.GetDepartment(ctx, id, 1, false)
	if err != nil {
		s.logger.Err(appErrors.ErrInvalidDepartmentNumber).Msg("department doesn't exist")
		return fmt.Errorf("department doesn't exist: %w", appErrors.ErrInvalidDepartmentNumber)
	}

	return nil
}
