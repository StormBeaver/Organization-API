package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	appErrors "orgService/internal/errors"
	"orgService/internal/model"
	"strconv"
)

func (h *Handler) createEmployee(w http.ResponseWriter, r *http.Request) {
	var emp model.Employee

	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		http.Error(w, "Bad JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	department := r.PathValue("id")

	depID, err := strconv.Atoi(department)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("wront type of department"))
		return
	}

	emp, err = h.service.CreateEmployee(r.Context(), emp.FullName, emp.Position, depID, emp.HiredAt)
	if err != nil {
		switch {
		case errors.Is(err, appErrors.ErrInvalidDepartmentNumber):
			http.Error(w, err.Error(), http.StatusNotFound)
		case errors.Is(err, appErrors.ErrInvalidFieldLength):
			http.Error(w, err.Error(), http.StatusBadRequest)
		case errors.Is(err, appErrors.ErrInvalidTime):
			http.Error(w, err.Error(), http.StatusBadRequest)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	empJSON, err := json.Marshal(emp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(empJSON)
}
