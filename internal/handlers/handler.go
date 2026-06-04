package handler

import (
	"context"
	"net/http"
	"orgService/internal/handlers/dto"
	"orgService/internal/model"
	"time"

	"github.com/rs/zerolog"
)

type Service interface {
	CreateDepartment(ctx context.Context, request dto.CreateDepartmentRequest) (*model.Department, error)
	CreateEmployee(ctx context.Context, request dto.CreateEmployeeRequest) (*model.Employee, error)
	GetDepartmentWithDepth(ctx context.Context, id int, request dto.GetDepartmentRequest) (*model.Department, error)
	PatchDepartment(ctx context.Context, request dto.PatchDepartmentRequest) (*model.Department, error)
	DeleteDepartment(ctx context.Context, request dto.DeleteDepartmentRequest) error
}

type Handler struct {
	service Service
	logger  *zerolog.Logger
}

func NewHandler(service Service, logger *zerolog.Logger) *Handler {
	return &Handler{service: service, logger: logger}
}

func (h *Handler) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{pong}"))
	})
	mux.HandleFunc("POST /departments", h.createDepartment)
	mux.HandleFunc("POST /departments/{id}/employees", h.createEmployee)
	mux.HandleFunc("GET /departments/{id}", h.getDepartment)
	mux.HandleFunc("PATCH /departments/{id}", h.patchDepartment)
	mux.HandleFunc("DELETE /departments/{id}", h.deleteDepartment)

	return h.loggingMiddleware(mux)
}

func (h *Handler) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		h.logger.Debug().
			Str("Method", r.Method).
			Str("Path", r.URL.Path).
			Dur("Duration", time.Since(start)).Send()
	})
}
