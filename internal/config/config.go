package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type DatabaseConfig = struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

var Database DatabaseConfig

func Init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error cargando el archivo .env:", err)
		return
	}

	Database = DatabaseConfig{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
	}

}
