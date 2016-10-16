package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"message-board/server/user"
	"net/http"
	"os"
	"time"

	"github.com/golang/glog"
	"github.com/urfave/negroni"

	mgo "gopkg.in/mgo.v2"
)

var logln = glog.Infoln

func testMySQL() {
	newUser := user.CreateUser("joint-song", "xuguang1992@gmail.com", "rootpwd", user.Man)
	fmt.Printf("New use: %+v\n", newUser)
}

func testMongo() {
	mongoURI := fmt.Sprintf("%s:%s", os.Getenv("MONGO_LINK"), os.Getenv("MONGO_PORT"))
	logln("mongoURI: ", mongoURI)
	session, err := mgo.DialWithTimeout(mongoURI, 2*time.Second)
	logln("err: ", err)
	if err != nil {
		fmt.Println("panic on mongo")
		panic(err)
	}
	defer session.Close()
}

func testServeApp() {
	mux := http.NewServeMux()
	mux.HandleFunc("/index", webHandler)
	mux.Handle("/", http.FileServer(http.Dir("../client/")))

	// middleare 这里可以先不用管它， 从line 41直接看line 47就行
	middleware := negroni.Classic()
	middleware.UseHandler(mux)

	// log.Fatal(http.ListenAndServe(":8091", mux))
	log.Fatal(http.ListenAndServe(":8091", middleware))
}

func webHandler(w http.ResponseWriter, r *http.Request) {
	t, error := template.ParseFiles("../client/index.html")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if error != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Error: %s", error)
		return
	}
	t.Execute(w, nil)
}

func main() {
	flag.Lookup("logtostderr").Value.Set("true")
	flag.Parse()
	logln("begin...")
	testMySQL()
	testMongo()
	testServeApp()
}
