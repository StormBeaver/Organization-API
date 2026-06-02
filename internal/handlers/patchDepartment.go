package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	appErrors "orgService/internal/errors"
	"orgService/internal/model"
	"strconv"
)

func (h *Handler) patchDepartment(w http.ResponseWriter, r *http.Request) {
	var dep model.Department

	err := json.NewDecoder(r.Body).Decode(&dep)
	if err != nil {
		http.Error(w, "Bad JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	department := r.PathValue("id")

	depID, err := strconv.Atoi(department)
	if err != nil {
		http.Error(w, "wrong type of department", http.StatusBadRequest)
	}

	res, err := h.service.PatchDepartment(r.Context(), depID, dep.Name, dep.ParentID)
	if err != nil {
		switch {
		case errors.Is(err, appErrors.ErrInvalidDepartmentNumber) ||
			errors.Is(err, appErrors.ErrInvalidFieldLength) ||
			errors.Is(err, appErrors.ErrInvalidArguments):
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
	w.WriteHeader(http.StatusOK)
	w.Write(resJSON)
}
