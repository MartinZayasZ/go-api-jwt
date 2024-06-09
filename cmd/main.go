package main

import (
	"GommerceApi/internal/handlers/auth"
	"GommerceApi/internal/handlers/users"
	"fmt"
	"log"
	"net/http"
)

func main() {

	r := http.NewServeMux()

	auth.RegisterHandlers(r)
	users.RegisterHandlers(r)

	fmt.Println("Server Listening: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
