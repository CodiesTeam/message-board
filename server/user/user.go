package user

import (
	"fmt"
	"os"
)

var (
	db = os.Getenv("USERS_DB_NAME")
	p  = fmt.Println
)

// User Basic user.
type User struct {
	ID    string
	Name  string
	Email string
	Pwd   string
}

// CreateUser will create a new user with name, email and password, return new user created right now.
func CreateUser(name, email, pwd string) *User {
	return &User{
		ID:    "lsdknx0",
		Name:  name,
		Email: email,
		Pwd:   pwd,
	}
}
