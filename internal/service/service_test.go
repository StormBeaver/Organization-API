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
	err           error
	createCounter int
	tme           time.Time
	parents       []model.Department
}

// CreateDepartment implements [Repo].
func (r *repoMock) CreateDepartment(ctx context.Context, department *model.Department) (*model.Department, error) {
	if r.err != nil {
		return nil, r.err
	}

	r.createCounter++
	return &model.Department{
			Id:        r.createCounter,
			Name:      department.Name,
			ParentID:  department.ParentID,
			Parent:    department.Parent,
			CreatedAt: r.tme,
		},
		nil
}

// CreateEmployee implements [Repo].
func (r *repoMock) CreateEmployee(ctx context.Context, employee *model.Employee) (*model.Employee, error) {
	panic("unimplemented")
}

// DeleteDepartment implements [Repo].
func (r *repoMock) DeleteDepartment(ctx context.Context, tx *gorm.DB, department *model.Department) error {
	panic("unimplemented")
}

// GetDepartment implements [Repo].
func (r *repoMock) GetDepartment(ctx context.Context, id int) (*model.Department, error) {
	if r.err != nil {
		return nil, r.err
	}

	idx := slices.IndexFunc(r.parents, func(a model.Department) bool {
		return a.Id == id
	})

	if idx == -1 {
		return nil, errors.New("not found")
	}

	return &r.parents[idx], nil
}

// GetDepartmentWithDepth implements [Repo].
func (r *repoMock) GetDepartmentWithDepth(ctx context.Context, id int, hint *model.GetParams) (*model.Department, error) {
	panic("unimplemented")
}

// PatchDepartment implements [Repo].
func (r *repoMock) PatchDepartment(ctx context.Context, department *model.Department) (*model.Department, error) {
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

func TestService_CreateDepartment(t *testing.T) {
	logger := log.Level(zerolog.DebugLevel)
	tme := time.Now()

	rep := &repoMock{
		err:           nil,
		createCounter: 0,
		tme:           tme,
		parents: []model.Department{model.Department{
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
					Name:     "    test1    ",
					ParentID: nil,
				},
				want: &model.Department{
					Id:        3,
					Name:      toPtr("test1"),
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
