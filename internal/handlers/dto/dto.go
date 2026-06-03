package dto

import (
	appErrors "orgService/internal/errors"
	"time"
)

type DeleteMode int

const (
	ModeDeleteCascade DeleteMode = iota
	ModeDeleteReassign
)

type CreateDepartmentRequest struct {
	Name     string `json:"name"`
	ParentID *int   `json:"parent_id"`
}

type CreateEmployeeRequest struct {
	DepartmentID int
	FullName     string     `json:"full_name"`
	Position     string     `json:"position"`
	HiredAt      *time.Time `json:"hired_at"`
}

type PatchDepartmentRequest struct {
	Id       int
	Name     *string `json:"name"`
	ParentID *int    `json:"parent_id"`
}

type DeleteDepartmentRequest struct {
	Id           int
	Mode         DeleteMode
	ToDepartment *int
}

func ParseDeleteMode(mode string) (DeleteMode, error) {
	switch mode {
	case "cascade":
		return ModeDeleteCascade, nil
	case "reassign":
		return ModeDeleteReassign, nil
	default:
		return 0, appErrors.ErrInvalidMode
	}
}
