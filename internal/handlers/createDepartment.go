package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	appErrors "orgService/internal/errors"
	"orgService/internal/handlers/dto"
)

// добавь проверку на уникальность имени в родителе
func (h *Handler) createDepartment(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateDepartmentRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Bad JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	h.logger.Debug().Msgf("name: %s, parent_id: %v", req.Name, req.ParentID)

	res, err := h.service.CreateDepartment(r.Context(), req)
	if err != nil {
		h.logger.Err(err).Send()
		switch {
		case errors.Is(err, appErrors.ErrInvalidDepartmentNumber) ||
			errors.Is(err, appErrors.ErrInvalidFieldLength):
			http.Error(w, err.Error(), http.StatusBadRequest)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	resJSON, err := json.Marshal(res)
	if err != nil {
		h.logger.Err(err).Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resJSON)
}
