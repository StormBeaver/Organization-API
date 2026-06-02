package repo

import (
	"context"
	"orgService/internal/model"
)

// я хз как это сделать...
func (r *repo) GetDepartment(ctx context.Context, id, depth int, employees bool) (model.Department, *[]model.Employee, error) {
	var department model.Department

	res := r.db.WithContext(ctx).Model(model.Department{}).Where("id = ?", id).Take(&department)
	return department, nil, res.Error
}
