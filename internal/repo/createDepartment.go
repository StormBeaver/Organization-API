package repo

import (
	"context"
	"fmt"
	"orgService/internal/model"
)

func (r *repo) CreateDepartment(ctx context.Context, name string, parentID *int) (model.Department, error) {
	department := model.Department{
		Name:     &name,
		ParentID: parentID,
	}

	res := r.db.WithContext(ctx).Model(model.Department{}).Create(&department)

	if res.Error != nil {
		return model.Department{}, fmt.Errorf("gorm create department: %w", res.Error)
	}

	return department, nil
}
