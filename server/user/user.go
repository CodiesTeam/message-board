package user

import (
	"fmt"
	"message-board/server/mysql"
	"os"
)

var (
	db    = os.Getenv("USERS_DB_NAME")
	table = "user"
)

// Gender for user.
type Gender uint

const (
	// Man Gender for man.
	Man Gender = iota
	// Woman Gender for woman.
	Woman
)

// User Basic user.
type User struct {
	ID     int64
	Name   string
	Email  string
	Pwd    string
	Gender Gender
}

// user table scheme.
// CREATE TABLE `user` (
//   `id` int(5) NOT NULL AUTO_INCREMENT,
//   `name` char(20) NOT NULL,
//   `gender` tinyint(1) NOT NULL DEFAULT '0',
//   `email` char(40) NOT NULL,
//   `password` char(50) NOT NULL,
//   PRIMARY KEY (`id`)
// )

// CreateUser will create a new user with name, email and password, return new user created right now.
func CreateUser(name, email, pwd string, gender Gender) *User {
	if mysql.MySQLDB == nil {
		fmt.Println("mysql is nil")
	}
	_, e := mysql.MySQLDB.Exec(fmt.Sprintf("USE %s", db))
	if e != nil {
		panic(fmt.Sprintf("CreateUser error: %s", e))
	}
	stmt, e := mysql.MySQLDB.Prepare(fmt.Sprintf("INSERT %s SET name=?,email=?,password=?,gender=?", table))
	if e != nil {
		panic(fmt.Sprintf("Prepare INSERT USER error: %s", e))
	}
	defer stmt.Close()
	res, e := stmt.Exec(name, email, pwd, gender)
	if e != nil {
		panic(fmt.Sprintf("INSERT USER error: %s", e))
	}
	id, e := res.LastInsertId()
	if e != nil {
		panic(fmt.Sprintf("INSERT USER and get last insert id error: %s", e))
	}
	return &User{
		ID:     id,
		Name:   name,
		Email:  email,
		Pwd:    pwd,
		Gender: gender,
	}
}
