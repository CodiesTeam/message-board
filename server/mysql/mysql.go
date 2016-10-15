package mysql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	link     = os.Getenv("MYSQL_LINK")
	port     = os.Getenv("MYSQL_PORT")
	user     = "root"
	password = "root"
)

var P = fmt.Println

func NewDB() *sql.DB {
	uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/users_test", user, password, link, port)
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
