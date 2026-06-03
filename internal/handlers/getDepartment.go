package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	appErrors "orgService/internal/errors"
	"orgService/internal/handlers/dto"
	"strconv"
)

func (h *Handler) getDepartment(w http.ResponseWriter, r *http.Request) {
	req := dto.GetDepartmentRequest{Depth: 1, IncludeEmployees: true}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Bad JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("wront type of department"))
		return
	}

	h.logger.Debug().Msgf("department: %d, depth: %d, includeEmployees: %v",
		id, req.Depth, req.IncludeEmployees)

	res, err := h.service.GetDepartmentWithDepth(r.Context(), id, req)
	if err != nil {
		switch {
		case errors.Is(err, appErrors.ErrInvalidDepth):
			http.Error(w, err.Error(), http.StatusBadRequest)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	resJSON, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resJSON)
}
