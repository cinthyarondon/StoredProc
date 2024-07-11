package model

/**
* Representa un objecto de negocio
* @author Cinthya Rondon
*/
type StoredProcedureCall struct {
	ProcedureName string        `json:"ProcedureName"`
	Inputs        []interface{} `json:"Inputs"`
	Outputs       int           `json:"Outputs"`
}
