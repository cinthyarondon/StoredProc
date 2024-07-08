package config

import (
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv carga las variables del entorno desde el archivo .env
func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}

// GetPort obtiene el puerto del entorno o devuelve el valor por defecto
func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Puerto por defecto si no se especifica en .env
	}
	return port
}