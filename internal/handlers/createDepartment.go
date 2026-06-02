package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	appErrors "orgService/internal/errors"
	"orgService/internal/model"
)

func (h *Handler) createDepartment(w http.ResponseWriter, r *http.Request) {
	var dep model.Department

	err := json.NewDecoder(r.Body).Decode(&dep)
	if err != nil {
		http.Error(w, "Bad JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	h.logger.Debug().Msgf("name: %s, parent_id: %v", dep.Name, dep.ID)

	res, err := h.service.CreateDepartment(r.Context(), dep.Name, dep.ParentID)
	if err != nil {
		switch {
		case errors.Is(err, appErrors.ErrInvalidDepartmentNumber):
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		case errors.Is(err, appErrors.ErrInvalidFieldLength):
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
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
	w.WriteHeader(http.StatusOK)
	w.Write(resJSON)
}
