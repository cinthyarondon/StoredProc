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

	/*
	err := config.LoadEnv()
	if err != nil {
		log.Fatalf("Error cargando el archivo .env: %v", err)
	}
	*/

	// Verifica las variables de entorno
	/*
    log.Printf("DB_USER: %s", os.Getenv("DB_USER"))
    log.Printf("DB_PASSWORD: %s", os.Getenv("DB_PASSWORD"))
    log.Printf("DB_HOST: %s", os.Getenv("DB_HOST"))
    log.Printf("DB_PORT: %s", os.Getenv("DB_PORT"))
    log.Printf("DB_SERVICE: %s", os.Getenv("DB_SERVICE"))
	*/

	db := dao.NewOracleDB()
	err := db.Connect()
	// err = db.Connect()
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}

	procedureService := application.NewProcedureService(db)
	handler := controller.NewProcedureController(procedureService)

	http.HandleFunc("/execute-procedure", handler.HandleExecuteProcedure)

	// Esto ya estaba
	port := config.GetPort()
	log.Printf("Servidor escuchando en el puerto %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}