package model

import "time"

type Department struct {
	ID        int       `db:"id"`
	Name      *string   `db:"name"`
	ParentID  *int      `db:"parent_id" json:"parent_id"`
	CreatedAt time.Time `db:"created_at"`
}

func (d Department) TableName() string {
	return "department"
}

type Employee struct {
	ID           int        `db:"id"`
	DepartmentID int        `db:"department_id"`
	FullName     string     `db:"full_name" json:"full_name"`
	Position     string     `db:"position"`
	HiredAt      *time.Time `db:"hired_at" json:"hired_at"`
	CreatedAt    time.Time  `db:"created_at"`
}

func (e Employee) TableName() string {
	return "employee"
}
