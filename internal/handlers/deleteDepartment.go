package handler

import (
	"errors"
	"net/http"
	appErrors "orgService/internal/errors"
	"orgService/internal/handlers/dto"
	"strconv"
)

func (h *Handler) deleteDepartment(w http.ResponseWriter, r *http.Request) {
	var req dto.DeleteDepartmentRequest
	var err error

	req.Mode, err = dto.ParseDeleteMode(r.URL.Query().Get("mode"))
	if err != nil {
		http.Error(w, "wrong delete mode", http.StatusBadRequest)
	}

	req.Id, err = strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "wrong type of department", http.StatusBadRequest)
	}

	if req.Mode == dto.ModeDeleteReassign {
		id, err := strconv.Atoi(r.URL.Query().Get("reassign_to_department_id"))
		if err != nil {
			http.Error(w, "wrong type of destination department", http.StatusBadRequest)
		}

		req.ToDepartment = &id
	}

	err = h.service.DeleteDepartment(r.Context(), req)
	if err != nil {
		h.logger.Err(err).Send()
		switch {
		case errors.Is(err, appErrors.ErrInvalidDepartmentNumber):
			http.Error(w, err.Error(), http.StatusBadRequest)
		case errors.Is(err, appErrors.ErrInvalidMode):
			http.Error(w, err.Error(), http.StatusBadRequest)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
