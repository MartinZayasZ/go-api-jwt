package types

type Response struct {
	Data  interface{} `json:"data"`
	Token string      `json:"token"`
}

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	password  string
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	CreatedBy int    `json:"created_by"`
	UpdatedAt string `json:"updated_at"`
	UpdatedBy int    `json:"updated_by"`
}

func (u *User) GetPassword() string {
	return u.password
}

func (u *User) SetPassword(password string) {
	u.password = password
}
