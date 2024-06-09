package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	password string
}

type Response struct {
	Data  User   `json:"data"`
	Token string `json:"token"`
}

func RegisterHandlers(r *http.ServeMux) {
	r.HandleFunc("/auth", login)
	r.HandleFunc("/logout", logout)
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login")

	res := Response{
		Data: User{
			Id:       1,
			Name:     "Martín Zayas",
			password: "hola mundo",
		},
		Token: "Hola soy un token",
	}

	responseSb, err := json.Marshal(res)

	if err != nil {
		fmt.Println("Ha ocurrido un error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "max-age: 60, must-revalidate")
	w.Write(responseSb)
}

func logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Logout")
}
