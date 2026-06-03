package repo

import (
	"context"
	"fmt"
	"orgService/internal/model"
	"strings"

	"gorm.io/gorm"
)

func (r *repo) GetDepartmentWithDepth(ctx context.Context, id int, hint *model.GetParams) (*model.Department, error) {
	var department model.Department

	query := r.db.WithContext(ctx).Model(model.Department{})

	if hint.IncludeEmployees {
		query = query.Preload("Children.Employees", func(db *gorm.DB) *gorm.DB {
			return db.Order("created_at ASC")
		}).
			Preload("Employees", func(db *gorm.DB) *gorm.DB {
				return db.Order("created_at ASC")
			})

		for i := 1; i <= hint.Depth-1; i++ {
			prefix := strings.Repeat("Children.", i) + "Children" + ".Employees"
			query = query.Preload(prefix, func(db *gorm.DB) *gorm.DB {
				return db.Order("created_at ASC")
			})
		}
	} else {
		query = query.Preload("Children")
		for i := 1; i <= hint.Depth-1; i++ {
			prefix := strings.Repeat("Children.", i) + "Children"
			query = query.Preload(prefix)

		}
	}
	err := query.Where("id = ?", id).Take(&department).Error
	if err != nil {
		return nil, fmt.Errorf("gorm get department :%w", err)
	}

	return &department, nil
}
