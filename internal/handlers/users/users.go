package users

import (
	"GommerceApi/internal/models/user_model"
	"GommerceApi/internal/types"
	"GommerceApi/pkgs/output"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func RegisterHandlers(r *http.ServeMux) {
	//r.HandleFunc("/users", getUsers)
	r.HandleFunc("POST /users", createUser)
}

func createUser(w http.ResponseWriter, r *http.Request) {

	data, err := validateForm(r)
	if err != nil {
		output.JsonResponse(w, struct{ Message string }{
			Message: err.Error(),
		}, http.StatusBadRequest, 0)
		return
	}
	fmt.Println(data)

	status := r.FormValue("status")

	if status != "active" && status != "inactive" && status != "pending" {
		status = "active"
	}

	createdBy, _ := strconv.Atoi(r.FormValue("created_by"))
	updatedBy, _ := strconv.Atoi(r.FormValue("updated_by"))

	user := types.User{
		FirstName: r.FormValue("first_name"),
		LastName:  r.FormValue("last_name"),
		Username:  r.FormValue("username"),
		Password:  r.FormValue("password"),
		Status:    status,
		CreatedAt: r.FormValue("created_at"),
		CreatedBy: createdBy,
		UpdatedAt: r.FormValue("updated_at"),
		UpdatedBy: updatedBy,
	}

	err = user_model.Create(&user)

	if err != nil {
		output.JsonResponse(w, struct{ Message string }{
			Message: err.Error(),
		}, http.StatusInternalServerError, 0)

		return
	}

	res := types.Response{Data: user,
		Token: "Hola soy un token",
	}

	output.JsonResponse(w, &res, http.StatusCreated, 0)
}

func validateForm(r *http.Request) (types.User, error) {
	var data types.User
	var validationErrors []string

	// Parse the form
	if err := r.ParseForm(); err != nil {
		return data, errors.New("unable to parse form")
	}

	// Validate Username
	data.Username = strings.TrimSpace(r.FormValue("username"))
	if data.Username == "" {
		validationErrors = append(validationErrors, "El Username es requerido")
	} else if len(data.Username) < 3 {
		validationErrors = append(validationErrors, "El username necesita un mínimo de 3 caracteres")
	}

	/*// Validate Email
	data.Email = r.FormValue("email")
	if !strings.Contains(data.Email, "@") {
		validationErrors = append(validationErrors, "Invalid email address")
	}

	// Validate Age
	ageStr := r.FormValue("age")
	age, err := strconv.Atoi(ageStr)
	if err != nil || age <= 0 {
		validationErrors = append(validationErrors, "Invalid age")
	} else {
		data.Age = age
	}
	*/

	// Return errors if any
	if len(validationErrors) > 0 {
		return data, errors.New(strings.Join(validationErrors, ", "))
	}

	return data, nil
}

/*func getUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login")

	data := types.Response{
		Data: types.User{
			Id:       1,
			Name:     "Martín Zayas",
			password: "hola mundo",
		},
		Token: "Hola soy un token",
	}

	output.JsonResponse(w, &data, 1)
}*/
