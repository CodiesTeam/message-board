package user

import (
	"database/sql"
	"fmt"
	"message-board/server/mysql"
	"message-board/server/utils"
	"os"
)

var (
	p      = fmt.Println
	dbName = os.Getenv("USERS_DB_NAME")
	db     *sql.DB
	table  = "user"
)

// Gender for user.
type Gender uint

const (
	// Man Gender for man.
	Man Gender = 1 + iota
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

func init() {
	db = mysql.NewDB()
	_, e := db.Exec(fmt.Sprintf("USE %s", dbName))
	utils.CheckError(e, "user.go")
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
func CreateUser(name, email, pwd string, gender Gender) (*User, error) {
	if !checkEmail(email) || !checkUserName(name) {
		return nil, fmt.Errorf("invalid user name or email address")
	}
	stmt, e := db.Prepare(fmt.Sprintf("SELECT id FROM %s WHERE email=?;", table))
	utils.CheckError(e, "user.go")
	defer stmt.Close()
	if GetUser(0, email) != nil {
		return nil, fmt.Errorf("email has registered")
	}
	stmt, e = db.Prepare(fmt.Sprintf("INSERT %s SET name=?,email=?,password=?,gender=?", table))
	utils.CheckError(e, "user.go")
	defer stmt.Close()
	res, e := stmt.Exec(name, email, pwd, gender)
	utils.CheckError(e, "user.go")
	id, e := res.LastInsertId()
	utils.CheckError(e, "user.go")
	return &User{
		ID:     id,
		Name:   name,
		Email:  email,
		Pwd:    pwd,
		Gender: gender,
	}, nil
}

// GetUser get User by id or email(one of them, email preferred), return nil if not exist.
func GetUser(id int64, email string) *User {
	if email != "" {
		return getUserByEmail(email)
	}
	return getUserByEmail(email)
}

func getUserByID(id int64) *User {
	stmt, e := db.Prepare(fmt.Sprintf("SELECT id name email gender FROM %s WHERE id = ?;", table))
	utils.CheckError(e, "user.go")
	defer stmt.Close()
	row := stmt.QueryRow(id)
	user := &User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Gender)
	if err == sql.ErrNoRows {
		return nil
	}
	utils.CheckError(err, "user.go")
	return user
}

func getUserByEmail(email string) *User {
	if !checkEmail(email) {
		return nil
	}
	stmt, e := db.Prepare(fmt.Sprintf("SELECT id, name, email, gender FROM %s WHERE email = ?;", table))
	utils.CheckError(e, "user.go")
	defer stmt.Close()
	row := stmt.QueryRow(email)
	user := &User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Gender)
	if err == sql.ErrNoRows {
		return nil
	}
	utils.CheckError(err, "user.go")
	return user
}
