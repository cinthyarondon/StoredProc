package application

import (
    "main/feat/validate/model"
    "main/feat/validate/dao"
)

type ProcedureServiceInterface interface {
    ExecuteStoredProcedure(call model.StoredProcedureCall) (map[string]interface{}, error)
}

// ProcedureService es la implementación de la interfaz ProcedureServiceInterface
type ProcedureService struct {
    db dao.Database
}

func NewProcedureService(db dao.Database) *ProcedureService {
    return &ProcedureService{db: db}
}

func (s *ProcedureService) ExecuteStoredProcedure(call model.StoredProcedureCall) (map[string]interface{}, error) {
    return s.db.ExecuteStoredProcedure(call)
}
