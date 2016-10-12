package main

import (
	"fmt"
	"html/template"
	"log"
	"message-board/server/user"
	"net/http"
	"os"
	"time"

	mgo "gopkg.in/mgo.v2"
)

func testMySQL() {
	newUser := user.CreateUser("joint-song", "xuguang1992@gmail.com", "rootpwd")
	fmt.Printf("New user name: %s\n", newUser.Name)
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
