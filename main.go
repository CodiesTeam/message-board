package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"gopkg.in/mgo.v2"
)

// var CREATE_TABLE = `CREATE TABLE `userinfo` (
//     `uid` INT(10) NOT NULL AUTO_INCREMENT,
//     `username` VARCHAR(64) NULL DEFAULT NULL,
//     `departname` VARCHAR(64) NULL DEFAULT NULL,
//     `created` DATE NULL DEFAULT NULL,
//     PRIMARY KEY (`uid`)
// )`
var (
	mysqlLink = "messageboard_mysql_1"
	mysqlPort = "3307"
	mysqlUser = "root"
	mysqlPWD  = "root"

	mongoLink = "messageboard_mongodb_1"
	mongoPort = "27017"
)

func testMySQL() {
	// dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/", "root", "root", os.Getenv("MYSQL_PORT_3306_TCP_ADDR"), os.Getenv("MYSQL_PORT_3306_TCP_PORT"))
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/", mysqlUser, mysqlPWD, mysqlLink, mysqlPort)
	glog.Infoln("mysqlURI: ", mysqlURI)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		panic("Open error")
	}
	glog.Infoln("link to mysql sucecess")
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

func testMongo() {

	// dbURI := fmt.Sprintf("%s:%s", os.Getenv("MONGO_PORT_27017_TCP_ADDR"), os.Getenv("MONGO_PORT_27017_TCP_PORT"))
	session, err := mgo.DialWithTimeout(fmt.Sprintf("%s:%s", mongoLink, mongoPort), 5*time.Second)
	if err != nil {
		fmt.Printf("MONGO ADDRESS: %s", os.Getenv("DATABASE_PORT_27017_TCP_ADDR"))
		fmt.Println("panic on mongo")
		panic(err)
	}
	glog.Infoln("link to mongo success")
	defer session.Close()
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
	flag.Lookup("logtostderr").Value.Set("true")
	flag.Parse()
	testMySQL()
	testMongo()
	for true {
		glog.Infoln("runing...")
		time.Sleep(time.Second * 1)
	}
}

// func main() {
// 	testMySQL()
// 	testMongo()
// 	for true {

// 	}
// }
