package handler

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	appErrors "orgService/internal/errors"
	"orgService/internal/handlers/dto"
	"orgService/internal/model"
	"strings"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

type serviceMok struct {
	err error
}

func toPtr[T any](v T) *T {
	return &v
}

func TestHandler_createDepartment(t *testing.T) {
	tests := []struct {
		name     string
		service  Service
		logger   *zerolog.Logger
		w        *httptest.ResponseRecorder
		r        *http.Request
		wantCode int
	}{
		{
			name:     "ok",
			service:  &serviceMok{},
			logger:   &zerolog.Logger{},
			w:        httptest.NewRecorder(),
			r:        httptest.NewRequest("POST", "/departments", strings.NewReader("{\"name\":\"departmentName\"}")),
			wantCode: 200,
		},
		{
			name:     "bad JSON",
			service:  &serviceMok{},
			logger:   &zerolog.Logger{},
			w:        httptest.NewRecorder(),
			r:        httptest.NewRequest("POST", "/departments", strings.NewReader("{\"name\":\"}")),
			wantCode: 400,
		},
		{
			name:     "bad business",
			service:  &serviceMok{err: appErrors.ErrInvalidDepartmentNumber},
			logger:   &zerolog.Logger{},
			w:        httptest.NewRecorder(),
			r:        httptest.NewRequest("POST", "/departments", strings.NewReader("{\"name\":\"}")),
			wantCode: 400,
		},
		{
			name:     "internalServerError",
			service:  &serviceMok{err: errors.New("unexpected error")},
			logger:   &zerolog.Logger{},
			w:        httptest.NewRecorder(),
			r:        httptest.NewRequest("POST", "/departments", strings.NewReader("{\"name\":\"departmentName\"}")),
			wantCode: 500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHandler(tt.service, tt.logger)
			h.createDepartment(tt.w, tt.r)
			assert.Equal(t, tt.wantCode, tt.w.Code)
		})
	}
}

// CreateDepartment implements [Service].
func (s *serviceMok) CreateDepartment(ctx context.Context, request dto.CreateDepartmentRequest) (*model.Department, error) {
	if s.err != nil {
		return nil, s.err
	}

	return &model.Department{
		Id:       1,
		Name:     toPtr(request.Name),
		ParentID: request.ParentID,
	}, nil
}

// CreateEmployee implements [Service].
func (s *serviceMok) CreateEmployee(ctx context.Context, request dto.CreateEmployeeRequest) (*model.Employee, error) {
	panic("unimplemented")
}

// DeleteDepartment implements [Service].
func (s *serviceMok) DeleteDepartment(ctx context.Context, request dto.DeleteDepartmentRequest) error {
	panic("unimplemented")
}

// GetDepartmentWithDepth implements [Service].
func (s *serviceMok) GetDepartmentWithDepth(ctx context.Context, id int, request dto.GetDepartmentRequest) (*model.Department, error) {
	panic("unimplemented")
}

// PatchDepartment implements [Service].
func (s *serviceMok) PatchDepartment(ctx context.Context, request dto.PatchDepartmentRequest) (*model.Department, error) {
	panic("unimplemented")
}

// var _ Service = &serviceMok{}
