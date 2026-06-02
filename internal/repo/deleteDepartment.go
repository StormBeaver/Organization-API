package repo

import (
	"context"
	"errors"
)

func (r *repo) DeleteDepartment(ctx context.Context, id int, mode string, reasignDestination *int) error {
	switch mode {
	case "cascade":
	case "reassign":
	default:
		return errors.New("unknown mode for deletion")
	}

	return nil
}
