package service

import (
	appErrors "orgService/internal/errors"
)

func validateFieldLength(field string) error {
	if len(field) < 1 || len(field) > 200 {
		return appErrors.ErrInvalidFieldLength
	}
	return nil
}
