package repo

import (
	"context"
	"fmt"
	"orgService/internal/model"
)

// я хз как это сделать...
func (r *repo) GetDepartment(ctx context.Context, id int) (*model.Department, error) {
	var department model.Department

	err := r.db.WithContext(ctx).Model(model.Department{}).Preload("Parent").Where("id = ?", id).Take(&department).Error
	if err != nil {
		return nil, fmt.Errorf("gorm get department :%w", err)
	}

	return &department, nil

}
