package mysql

import (
	"database/sql"
	"fmt"
	"os"
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
	uri := fmt.Sprintf("%s:%s@tcp(%s:%s)", user, password, link, port)
	MySQLDB, error := sql.Open("mysql", uri)
	if error != nil {
		MySQLDB.Close()
		panic(fmt.Sprintf("mysql init error: %s\n", error))
	}
}
