package model

import "time"

type Department struct {
	Id   int     `gorm:"primary_key;column:id;type:bigserial;not null" json:"id"`
	Name *string `gorm:"column:name;type:text;not null" json:"name"`

	ParentID  *int         `json:"parent_id"`
	Parent    *Department  `gorm:"foreignKey:ParentID;references:Id" json:"-"`
	Children  []Department `gorm:"foreignKey:ParentID;references:Id" json:"children,omitempty"`
	Employees []Employee   `gorm:"foreignKey:DepartmentID;references:id" json:"employees,omitempty"`

	CreatedAt time.Time `gorm:"column:created_at;type:datetime;not null;default:now()" json:"created_at"`
}

func (d Department) TableName() string {
	return "department"
}

type Employee struct {
	Id int `gorm:"primary_key;column:id;type:bigserial;not null" json:"id"`

	DepartmentID int         `json:"department_id"`
	Department   *Department `gorm:"foreignKey:DepartmentID;references:id" json:"-"`

	FullName  string     `gorm:"column:full_name;type:text;not null" json:"full_name"`
	Position  string     `gorm:"column:position;type:text" json:"position"`
	HiredAt   *time.Time `gorm:"column:hired_at;type:date" json:"hired_at"`
	CreatedAt time.Time  `gorm:"column:created_at;type:datetime;not null;default:now()" json:"created_at"`
}

func (e Employee) TableName() string {
	return "employee"
}

type GetParams struct {
	Depth            int
	IncludeEmployees bool
}
