package mysql

import (
	"database/sql"
	"fmt"
	"os"
	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

var (
	link     = os.Getenv("MYSQL_LINK")
	port     = os.Getenv("MYSQL_PORT")
	user     = "root"
	password = "root"
)

// P is a convenience for fmt.Println
var P = fmt.Println

// NewDB is a method for create a sql.DB, this method did not specify the database, you should `USE databasename` to use sql.DB created.
func NewDB() *sql.DB {
	uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/", user, password, link, port)
	P("mysql uri: ", uri)
	MySQLDB, err := sql.Open("mysql", uri)
	P("err: ", err)
	if err != nil {
		MySQLDB.Close()
		panic(fmt.Sprintf("mysql init err: %s\n", err))
	}
	err = MySQLDB.Ping()
	P("ping err: ", err)
	return MySQLDB
}
