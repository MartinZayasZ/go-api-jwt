package user_model

import (
	db "GommerceApi/internal/database"
	"GommerceApi/internal/types"
	"fmt"
)

func GetUser() string {
	return "HOLA"
}

func Create(user *types.User) error {
	db.InitDB()
	defer db.DB.Close()

	// Define la consulta de inserción
	query := "INSERT INTO users (first_name, last_name, username, email, password, status, created_by, updated_by) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	// Ejecuta la consulta de inserción
	result, err := db.DB.Exec(query, user.FirstName, user.LastName, user.Username, user.Email, user.Password, user.Status, user.CreatedBy, user.UpdatedBy)
	if err != nil {
		return err
	}

	// Obtiene el ID del último registro insertado
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.Id = int(lastInsertID)

	fmt.Printf("Se insertó un nuevo registro con ID %d\n", lastInsertID)

	return nil

}
