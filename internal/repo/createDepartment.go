package repo

import (
	"context"
	"fmt"
	"orgService/internal/model"
)

func (r *repo) CreateDepartment(ctx context.Context, department *model.Department) (*model.Department, error) {

	res := r.db.WithContext(ctx).Model(model.Department{}).Create(&department)

	if res.Error != nil {
		return nil, fmt.Errorf("gorm create department: %w", res.Error)
	}

	return department, nil
}
