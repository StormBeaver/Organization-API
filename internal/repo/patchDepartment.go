package repo

import (
	"context"
	"fmt"
	"orgService/internal/model"
)

func (r *repo) PatchDepartment(ctx context.Context, id int, name *string, parentID *int) (model.Department, error) {
	// fmt.Println(*parentID, *name, "------------------------------------------------------------")
	var department model.Department

	res := r.db.WithContext(ctx).Model(&model.Department{}).Where("id = ?", id)

	if parentID != nil {
		res = res.Update("parent_id", *parentID)
	}

	if name != nil {
		res = res.Update("name", *name)
	}

	if res.Error != nil {
		return model.Department{}, fmt.Errorf("gorm patch department: %w", res.Error)
	}
	return department, nil
}
