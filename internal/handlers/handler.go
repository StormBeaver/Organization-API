package handler

import (
	"context"
	"net/http"
	"orgService/internal/model"
	"time"

	"github.com/rs/zerolog"
)

type Service interface {
	CreateDepartment(ctx context.Context, name string, parentID *int) (model.Department, error)
	CreateEmployee(ctx context.Context, name string, position string, depID int, hiredAt *time.Time) (model.Employee, error)
	GetDepartment(ctx context.Context, id, depth int, employees bool) (model.Department, *[]model.Employee, error)
	PatchDepartment(ctx context.Context, id int, name *string, parentID *int) (model.Department, error)
	DeleteDepartment(ctx context.Context, id int, mode string, reasignDestination *int) error

	LastDepartment(ctx context.Context) (model.Department, error)
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

	return mux
}
