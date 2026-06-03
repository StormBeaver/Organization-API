package repo

import (
	"context"
	"fmt"
	"orgService/internal/model"
)

func (r *repo) CreateEmployee(ctx context.Context, employee *model.Employee) (*model.Employee, error) {

	res := r.db.WithContext(ctx).Model(model.Employee{}).Create(&employee)

	if res.Error != nil {
		return nil, fmt.Errorf("gorm create employee: %w", res.Error)
	}

	return employee, nil
}
