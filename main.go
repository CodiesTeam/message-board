package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// var CREATE_TABLE = `CREATE TABLE `userinfo` (
//     `uid` INT(10) NOT NULL AUTO_INCREMENT,
//     `username` VARCHAR(64) NULL DEFAULT NULL,
//     `departname` VARCHAR(64) NULL DEFAULT NULL,
//     `created` DATE NULL DEFAULT NULL,
//     PRIMARY KEY (`uid`)
// )`

func testMySQL() {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", "root", "root", os.Getenv("MYSQL_PORT_3306_TCP_ADDR"), os.Getenv("MYSQL_PORT_3306_TCP_PORT"), os.Getenv("USERS_DB_NAME"))
	fmt.Println(dbURI)
	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		panic("Open error")
	}
	defer db.Close()
	// dbName := "test"
	// _, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	// if err != nil {
	// 	fmt.Print(err)
	// 	panic("create database error")
	// }
	// _, err = db.Exec("USE " + dbName)
	// if err != nil {
	// 	fmt.Print(err)
	// 	panic("user error")
	// }
	// stmt, err := db.Prepare(`CREATE DATABASE 'test'`)
	// if err != nil {
	// 	fmt.Print(err)
	// 	panic("Crate Database error")
	// }
	// _, err = stmt.Exec()
	// if err != nil {
	// 	panic("Create database error2")
	// }

	// stmt, err := db.Prepare(`CREATE TABLE IF NOT EXISTS userinfo (
	//     uid INT(10) NOT NULL AUTO_INCREMENT,
	//     username VARCHAR(64) NULL DEFAULT NULL,
	//     departname VARCHAR(64) NULL DEFAULT NULL,
	//     created DATE NULL DEFAULT NULL,
	//     PRIMARY KEY (uid)
	// )`)
	// if err != nil {
	// 	fmt.Print(err)
	// 	panic("Create table error")
	// }
	// _, err = stmt.Exec()
	// if err != nil {
	// 	fmt.Print(err)
	// 	panic("Create table error 2")
	// }
}

// func testMongo() {

// 	// dbURI := fmt.Sprintf("%s:%s", os.Getenv("MONGO_PORT_27017_TCP_ADDR"), os.Getenv("MONGO_PORT_27017_TCP_PORT"))
// 	session, err := mgo.DialWithTimeout(os.Getenv("DATABASE_PORT_27017_TCP_ADDR"), 5*time.Second)
// 	if err != nil {
// 		fmt.Printf("MONGO ADDRESS: %s", os.Getenv("DATABASE_PORT_27017_TCP_ADDR"))
// 		fmt.Println("panic on mongo")
// 		panic(err)
// 	}
// 	defer session.Close()
// }

func main() {
	testMySQL()
	for true {

	}
}
