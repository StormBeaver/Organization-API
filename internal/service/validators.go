package service

import (
	"context"
	"fmt"
	appErrors "orgService/internal/errors"
	"time"
)

func (s *Service) validateCreateDepartment(ctx context.Context, name string, parent_id *int) error {
	if parent_id == nil && len(name) < 200 && len(name) > 0 {
		return nil
	}

	dep, err := s.LastDepartment(ctx)
	if err != nil {
		s.logger.Err(err).Msg("Error while get LastDepartment from db")
		return err
	}

	s.logger.Debug().Msgf("name length: %v, depaprment_id: %v, parent_id: %v", len(name), dep.ID, *parent_id)

	err = validateFieldLength(name)
	if err != nil {
		s.logger.Err(appErrors.ErrInvalidFieldLength).Msg("name too long")
		return fmt.Errorf("name invalid: %w", appErrors.ErrInvalidFieldLength)
	}

	if dep.ID < *parent_id {
		s.logger.Err(appErrors.ErrInvalidFieldLength).Msg("name too long")
		return appErrors.ErrInvalidDepartmentNumber
	}
	return nil
}

func (s *Service) validateCreateEmployee(ctx context.Context, name, position string, depID int, hiredAt *time.Time) error {
	err := validateFieldLength(name)
	if err != nil {
		s.logger.Err(appErrors.ErrInvalidFieldLength).Msg("invalid name length")
		return fmt.Errorf("name invalid: %w", appErrors.ErrInvalidFieldLength)
	}

	err = validateFieldLength(position)
	if err != nil {
		s.logger.Err(appErrors.ErrInvalidFieldLength).Msg("invalid position length")
		return fmt.Errorf("position invalid: %w", appErrors.ErrInvalidFieldLength)
	}

	if hiredAt != nil && time.Now().Before(*hiredAt) {
		return appErrors.ErrInvalidTime
	}

	_, _, err = s.GetDepartment(ctx, depID, 1, false)
	if err != nil {
		s.logger.Err(appErrors.ErrInvalidDepartmentNumber).Msg("department doesn't exist")
		return fmt.Errorf("department doesn't exist: %w", appErrors.ErrInvalidDepartmentNumber)
	}

	return nil
}

func validateFieldLength(field string) error {
	if len(field) < 1 || len(field) > 200 {
		return appErrors.ErrInvalidFieldLength
	}
	return nil
}
