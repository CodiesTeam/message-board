package user

import (
	"fmt"
	"message-board/server/mysql"
	"os"
)

var (
	db = os.Getenv("USERS_DB_NAME")
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
	_, error := mysql.MySQLDB.Exec(fmt.Sprintf("USE %s", db))
	if error != nil {
		panic(fmt.Sprintf("CreateUser error: %s", error))
	}
	// TODO: return a temporary User
	return &User{
		ID:    "lsdknx0",
		Name:  name,
		Email: email,
		Pwd:   pwd,
	}
}
