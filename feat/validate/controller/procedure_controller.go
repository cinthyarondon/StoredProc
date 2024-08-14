package controller

import (
    "encoding/json"
    "net/http"
	"strconv"
    "os"
    "log"

    "main/common/dto"
    "main/feat/validate/application"
    "main/feat/validate/model"
)

type ProcedureController struct {
    service application.ProcedureServiceInterface
}

/**
* Constructor que inicializa una instancia del controlador 
*/
func NewProcedureController(service application.ProcedureServiceInterface) *ProcedureController {
    return &ProcedureController{service: service}
}

func (h *ProcedureController) HandleExecuteProcedure(w http.ResponseWriter, r *http.Request) {

    /* Si es método Post
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    */

    // Verificar que el método de la solicitud sea GET
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // Nuevo código

	// Leer las variables de entorno del objeto del negocio
	procedureName := os.Getenv("PROCEDURE_NAME")
	inputsJSON := os.Getenv("INPUTS")
	outputsStr := os.Getenv("OUTPUTS")

    // Verificar variables de entorno
    if procedureName == "" || inputsJSON == "" || outputsStr == "" {
        http.Error(w, "Faltan variables de entorno requeridas para StoredProcedureCall", http.StatusInternalServerError)
        return
    }

    var call model.StoredProcedureCall
    call.ProcedureName = procedureName

    // Log for debugging
    log.Println("Procedimiento:", procedureName)
    log.Println("Inputs JSON:", inputsJSON)
    log.Println("Outputs:", outputsStr)

    // Deserialización (convertir datos en un objeto de memoria) del campo Inputs
    // Convertir la cadena JSON en un slice de interfaces ([]interface{})
    err := json.Unmarshal([]byte(inputsJSON), &call.Inputs)
    if err != nil {
        log.Printf("Error al decodificar Inputs: %v\nInputs JSON: %s\n", err, inputsJSON)
        http.Error(w, "Error al decodificar Inputs", http.StatusBadRequest)
        return
    }

    outputs, err := strconv.Atoi(outputsStr)
    if err != nil {
        log.Printf("Error al convertir Outputs a entero: %v\n", err)
        http.Error(w, "Error al convertir Outputs a entero", http.StatusBadRequest)
        return
    }
    call.Outputs = outputs

    /*
    var call model.StoredProcedureCall
    err := json.NewDecoder(r.Body).Decode(&call)
    if err != nil {
        http.Error(w, "Error al decodificar la solicitud JSON", http.StatusBadRequest)
        return
    }
    */
    
    // Verificar que el nombre del procedimiento no esté vacío
    if call.ProcedureName == "" {
        http.Error(w, "Debe proporcionar el nombre del procedimiento", http.StatusBadRequest)
        return
    }

    // Ejecutar el procedimiento almacenado usando el servicio
    result, err := h.service.ExecuteStoredProcedure(call)
    if err != nil {
        log.Printf("Error ejecutando procedimiento: %v\n", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    response := dto.Response{
        Data: result,
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}
