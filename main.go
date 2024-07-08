package main

import (
	"log"
	"net/http"

	"main/common/config"
	"main/feat/validate/application"
	"main/feat/validate/controller"
	"main/feat/validate/dao"
)

func main() {
	err := config.LoadEnv()
	if err != nil {
		log.Fatalf("Error cargando el archivo .env: %v", err)
	}

	db := dao.NewOracleDB()
	err = db.Connect()
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}

	procedureService := application.NewProcedureService(db)
	handler := controller.NewProcedureController(procedureService)

	http.HandleFunc("/execute-procedure", handler.HandleExecuteProcedure)

	port := config.GetPort()
	log.Printf("Servidor escuchando en el puerto %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}