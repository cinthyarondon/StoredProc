package controller

import (
    "encoding/json"
    "net/http"

    "main/common/dto"
    "main/feat/validate/application"
    "main/feat/validate/model"
)

type ProcedureController struct {
    service application.ProcedureServiceInterface
}

func NewProcedureController(service application.ProcedureServiceInterface) *ProcedureController {
    return &ProcedureController{service: service}
}

func (h *ProcedureController) HandleExecuteProcedure(w http.ResponseWriter, r *http.Request) {
    var call model.StoredProcedureCall
    err := json.NewDecoder(r.Body).Decode(&call)
    if err != nil {
        http.Error(w, "Error al decodificar la solicitud JSON", http.StatusBadRequest)
        return
    }

    if call.ProcedureName == "" {
        http.Error(w, "Debe proporcionar el nombre del procedimiento", http.StatusBadRequest)
        return
    }

    result, err := h.service.ExecuteStoredProcedure(call)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    response := dto.Response{
        Data: result,
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}
