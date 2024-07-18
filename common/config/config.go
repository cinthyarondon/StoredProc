package config

import (
	"os"

	//"github.com/joho/godotenv"
)

/**
* Carga las variables de entorno desde el archivo .env
* @author Cinthya Rondon
*/

/*
func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}
/*

/**
* Obtiene el valor de una variable de entorno especifica, si está no tiene valor,
* retorn el valor por defecto
* @author Cinthya Rondon
*/
func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Puerto por defecto si no se especifica en .env
	}
	return port
}