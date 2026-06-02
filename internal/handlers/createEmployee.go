package handler

import (
	"encoding/json"
	"net/http"
	"orgService/internal/model"
	"strconv"
	"strings"
)

func (h *Handler) createEmployee(w http.ResponseWriter, r *http.Request) {
	var emp model.Employee

	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		http.Error(w, "Bad JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	department, ok := strings.CutPrefix(r.URL.Path, "/departments/")
	department, ok = strings.CutSuffix(department, "/employees")

	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	depID, err := strconv.Atoi(department)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("wront type of department"))
		return
	}

	h.service.CreateEmployee(r.Context(), emp.FullName, emp.Position, depID, emp.HiredAt)
}
