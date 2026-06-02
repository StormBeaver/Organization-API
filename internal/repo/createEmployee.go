package repo

import (
	"context"
	"fmt"
	"orgService/internal/model"
	"time"
)

func (r *repo) CreateEmployee(ctx context.Context, name string, position string, depID int, hiredAt *time.Time) (model.Employee, error) {
	employee := model.Employee{
		DepartmentID: depID,
		FullName:     name,
		Position:     position,
		HiredAt:      hiredAt,
	}
	res := r.db.WithContext(ctx).Model(model.Employee{}).Create(&employee)

	if res.Error != nil {
		return model.Employee{}, fmt.Errorf("gorm create employee: %w", res.Error)
	}

	return employee, nil
}
