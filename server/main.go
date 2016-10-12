package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	mgo "gopkg.in/mgo.v2"

	_ "github.com/go-sql-driver/mysql"
)

func testMySQL() {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", "root", "root", os.Getenv("MYSQL_LINK"), os.Getenv("MYSQL_PORT_3306_TCP_PORT"), os.Getenv("USERS_DB_NAME"))
	fmt.Println(dbURI)
	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		panic("Open error")
	}
	defer db.Close()
}

func testMongo() {
	session, err := mgo.DialWithTimeout(fmt.Sprintf("%s:%s", os.Getenv("MONGO_LINK"), "27017"), 5*time.Second)
	if err != nil {
		fmt.Println("panic on mongo")
		panic(err)
	}
	defer session.Close()
}

func testServeApp() {
	http.HandleFunc("/index", webHandler)
	http.Handle("/", http.FileServer(http.Dir("./client/")))
	log.Fatal(http.ListenAndServe(":8091", nil))
}

func webHandler(w http.ResponseWriter, r *http.Request) {
	t, error := template.ParseFiles("./client/index.html")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if error != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Error: %s", error)
		return
	}
	t.Execute(w, nil)
}

func main() {
	testMySQL()
	testMongo()
	testServeApp()
}
