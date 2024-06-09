// db.go
package db

import (
	"GommerceApi/internal/config"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	//cargamos el archivo de configuración
	config.Init()

	// Obtiene las credenciales de la base de datos desde variables de entorno
	dbUser := config.Database.User
	dbPassword := config.Database.Password
	dbHost := config.Database.Host
	dbPort := config.Database.Port
	dbName := config.Database.Name

	// Cadena de conexión a la base de datos MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Abre la conexión a la base de datos
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}

	// Intenta conectar a la base de datos
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Conexión exitosa a la base de datos MySQL!")
	DB = db
}
