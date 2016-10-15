package mysql

import (
	"database/sql"
	"fmt"
	"os"
	// for mysql dirver
	_ "github.com/go-sql-driver/mysql"
)

var (
	link     = os.Getenv("MYSQL_LINK")
	port     = os.Getenv("MYSQL_DEFAULT_PORT")
	user     = "root"
	password = "root"
)

// MySQLDB is common used sql.DB.
var MySQLDB *sql.DB

func init() {
	uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/", user, password, link, port)
	var e error
	MySQLDB, e = sql.Open("mysql", uri)
	if e != nil {
		MySQLDB.Close()
		panic(fmt.Sprintf("mysql init error: %s\n", e))
	}

}
