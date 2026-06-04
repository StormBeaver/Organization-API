package service

import (
	"context"
	"errors"
	appErrors "orgService/internal/errors"
	"orgService/internal/handlers/dto"
	"orgService/internal/model"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type repoMock struct {
	err         error
	counterID   int
	tme         time.Time
	departments []model.Department
}

func TestService_CreateDepartment(t *testing.T) {
	logger := log.Level(zerolog.DebugLevel)
	tme := time.Now()

	rep := &repoMock{
		err:       nil,
		counterID: 0,
		tme:       tme,
		departments: []model.Department{
			model.Department{
				Id:        1,
				Name:      toPtr("testDepartment1"),
				ParentID:  nil,
				CreatedAt: tme,
			}, model.Department{
				Id:        2,
				Name:      toPtr("testDepartment2"),
				ParentID:  toPtr(1),
				CreatedAt: tme,
			}},
	}

	t.Run("ok", func(t *testing.T) {
		testsOk := []struct {
			name    string
			repo    Repo
			request dto.CreateDepartmentRequest
			want    *model.Department
		}{
			{
				name: "NoParent",
				repo: rep,
				request: dto.CreateDepartmentRequest{
					Name:     "test1",
					ParentID: nil,
				},
				want: &model.Department{
					Id:        1,
					Name:      toPtr("test1"),
					ParentID:  nil,
					CreatedAt: tme,
				},
			},
			{
				name: "HasParent",
				repo: rep,
				request: dto.CreateDepartmentRequest{
					Name:     "test2",
					ParentID: toPtr(1),
				},
				want: &model.Department{
					Id:       2,
					Name:     toPtr("test2"),
					ParentID: nil,
					Parent: &model.Department{
						Id:        1,
						Name:      toPtr("testDepartment1"),
						CreatedAt: tme,
					},
					CreatedAt: tme,
				},
			},
			{
				name: "TrimmedName",
				repo: rep,
				request: dto.CreateDepartmentRequest{
					Name:     "    test3    ",
					ParentID: nil,
				},
				want: &model.Department{
					Id:        3,
					Name:      toPtr("test3"),
					ParentID:  nil,
					CreatedAt: tme,
				},
			},
		}
		for _, tt := range testsOk {
			t.Run(tt.name, func(t *testing.T) {
				s := NewService(tt.repo, &logger)
				got, gotErr := s.CreateDepartment(context.Background(), tt.request)
				assert.NoError(t, gotErr)
				assert.Equal(t, tt.want, got)
			})
		}
	})

	t.Run("errors", func(t *testing.T) {
		testsErr := []struct {
			name    string
			repo    Repo
			request dto.CreateDepartmentRequest
			want    error
		}{
			{
				name:    "shortName",
				repo:    rep,
				request: dto.CreateDepartmentRequest{},
				want:    appErrors.ErrInvalidFieldLength,
			},
			{
				name:    "longName",
				repo:    rep,
				request: dto.CreateDepartmentRequest{Name: strings.Repeat("s", 201)},
				want:    appErrors.ErrInvalidFieldLength,
			},
			{
				name:    "InvalidParentID",
				repo:    rep,
				request: dto.CreateDepartmentRequest{Name: "invalidParent", ParentID: toPtr(3)},
				want:    appErrors.ErrInvalidDepartmentNumber,
			},
		}

		for _, tt := range testsErr {
			t.Run(tt.name, func(t *testing.T) {
				s := NewService(tt.repo, &logger)
				got, gotErr := s.CreateDepartment(context.Background(), tt.request)
				assert.Nil(t, got)
				assert.ErrorIs(t, gotErr, tt.want)
			})
		}
	})
}

// CreateDepartment implements [Repo].
func (r *repoMock) CreateDepartment(ctx context.Context, department *model.Department) (*model.Department, error) {
	if r.err != nil {
		return nil, r.err
	}

	r.counterID++
	return &model.Department{
			Id:        r.counterID,
			Name:      department.Name,
			ParentID:  department.ParentID,
			Parent:    department.Parent,
			CreatedAt: r.tme,
		},
		nil
}

func TestService_CreateEmployee(t *testing.T) {
	logger := log.Level(zerolog.DebugLevel)
	tme := time.Now()

	rep := &repoMock{
		err:       nil,
		counterID: 0,
		tme:       tme,
		departments: []model.Department{
			model.Department{
				Id:        1,
				Name:      toPtr("testDepartment1"),
				ParentID:  nil,
				CreatedAt: tme,
			}, model.Department{
				Id:        2,
				Name:      toPtr("testDepartment2"),
				ParentID:  toPtr(1),
				CreatedAt: tme,
			}},
	}
	t.Run("testsOk", func(t *testing.T) {
		testsOk := []struct {
			name    string
			repo    Repo
			logger  *zerolog.Logger
			request dto.CreateEmployeeRequest
			want    *model.Employee
		}{
			{
				name:   "okWithoutHiredAt",
				repo:   rep,
				logger: &logger,
				request: dto.CreateEmployeeRequest{
					DepartmentID: 1,
					FullName:     "testName1",
					Position:     "testPosition1",
				},
				want: &model.Employee{
					Id: 1,
					Department: &model.Department{
						Id:        1,
						Name:      toPtr("testDepartment1"),
						CreatedAt: tme,
					},
					FullName:  "testName1",
					Position:  "testPosition1",
					CreatedAt: tme,
				},
			},
			{
				name:   "okWithHiredAt",
				repo:   rep,
				logger: &logger,
				request: dto.CreateEmployeeRequest{
					DepartmentID: 1,
					FullName:     "testName2",
					Position:     "testPosition2",
					HiredAt:      toPtr(tme),
				},
				want: &model.Employee{
					Id: 2,
					Department: &model.Department{
						Id:        1,
						Name:      toPtr("testDepartment1"),
						CreatedAt: tme,
					},
					FullName:  "testName2",
					Position:  "testPosition2",
					CreatedAt: tme,
					HiredAt:   toPtr(tme),
				},
			},
		}
		for _, tt := range testsOk {
			t.Run(tt.name, func(t *testing.T) {
				s := NewService(tt.repo, &logger)
				got, gotErr := s.CreateEmployee(context.Background(), tt.request)
				assert.NoError(t, gotErr)
				assert.Equal(t, tt.want, got)
			})
		}
	})

	t.Run("testsErr", func(t *testing.T) {
		testsErr := []struct {
			name    string
			repo    Repo
			logger  *zerolog.Logger
			request dto.CreateEmployeeRequest
			want    error
		}{
			{
				name:   "shortFullName",
				repo:   rep,
				logger: &logger,
				request: dto.CreateEmployeeRequest{
					DepartmentID: 1,
					FullName:     "",
					Position:     "testPosition1",
				},
				want: appErrors.ErrInvalidFieldLength,
			},
			{
				name:   "LongFullName",
				repo:   rep,
				logger: &logger,
				request: dto.CreateEmployeeRequest{
					DepartmentID: 1,
					FullName:     strings.Repeat("s", 201),
					Position:     "testPosition2",
				},
				want: appErrors.ErrInvalidFieldLength,
			},
			{
				name:   "shortPosition",
				repo:   rep,
				logger: &logger,
				request: dto.CreateEmployeeRequest{
					DepartmentID: 1,
					FullName:     "testName3",
					Position:     "",
				},
				want: appErrors.ErrInvalidFieldLength,
			},
			{
				name:   "LongPosition",
				repo:   rep,
				logger: &logger,
				request: dto.CreateEmployeeRequest{
					DepartmentID: 1,
					FullName:     "testName4",
					Position:     strings.Repeat("s", 201),
				},
				want: appErrors.ErrInvalidFieldLength,
			},
			{
				name:   "InvalidDepartment",
				repo:   rep,
				logger: &logger,
				request: dto.CreateEmployeeRequest{
					DepartmentID: 3,
					FullName:     "testName5",
					Position:     "testPosition5",
				},
				want: appErrors.ErrInvalidDepartmentNumber,
			},
			{
				name:   "timeTraveler",
				repo:   rep,
				logger: &logger,
				request: dto.CreateEmployeeRequest{
					DepartmentID: 1,
					FullName:     "testName6",
					Position:     "testPosition6",
					HiredAt:      toPtr(tme.Add(1 * time.Hour)),
				},
				want: appErrors.ErrInvalidTime,
			},
		}
		for _, tt := range testsErr {
			t.Run(tt.name, func(t *testing.T) {
				s := NewService(tt.repo, &logger)
				got, gotErr := s.CreateEmployee(context.Background(), tt.request)
				assert.Nil(t, got)
				assert.ErrorIs(t, gotErr, tt.want)
			})
		}
	})
}

// CreateEmployee implements [Repo].
func (r *repoMock) CreateEmployee(ctx context.Context, employee *model.Employee) (*model.Employee, error) {
	if r.err != nil {
		return nil, r.err
	}

	r.counterID++
	return &model.Employee{
		Id:         r.counterID,
		Department: employee.Department,
		FullName:   employee.FullName,
		Position:   employee.Position,
		HiredAt:    employee.HiredAt,
		CreatedAt:  r.tme,
	}, nil
}

// PatchDepartment implements [Repo].
func (r *repoMock) PatchDepartment(ctx context.Context, department *model.Department) (*model.Department, error) {
	panic("unimplemented")
}

// GetDepartment implements [Repo].
func (r *repoMock) GetDepartment(ctx context.Context, id int) (*model.Department, error) {
	if r.err != nil {
		return nil, r.err
	}

	idx := slices.IndexFunc(r.departments, func(a model.Department) bool {
		return a.Id == id
	})

	if idx == -1 {
		return nil, errors.New("not found")
	}

	return &r.departments[idx], nil
}

// DeleteDepartment implements [Repo].
func (r *repoMock) DeleteDepartment(ctx context.Context, tx *gorm.DB, department *model.Department) error {
	panic("unimplemented")
}

// GetDepartmentWithDepth implements [Repo].
func (r *repoMock) GetDepartmentWithDepth(ctx context.Context, id int, hint *model.GetParams) (*model.Department, error) {
	panic("unimplemented")
}

// ReassignDepartment implements [Repo].
func (r *repoMock) ReassignDepartment(ctx context.Context, tx *gorm.DB, src int, dst int) error {
	panic("unimplemented")
}

// BeginTx implements [Repo].
func (r *repoMock) BeginTx() *gorm.DB {
	panic("unimplemented")
}

func toPtr[T any](v T) *T {
	return &v
}
