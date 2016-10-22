package main

import (
	"flag"
	"fmt"
	"html/template"
	"message-board/server/mbmux"
	"message-board/server/user"
	"net/http"
	"os"
	"time"

	"github.com/golang/glog"

	mgo "gopkg.in/mgo.v2"
)

var logln = glog.Infoln

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

func startupServe() {
	user.Serve()
	mbmux.SharedMux.HandleFunc("/index", webHandler)
	mbmux.ServeAll()
}

func main() {
	flag.Lookup("logtostderr").Value.Set("true")
	flag.Parse()
	logln("begin...")
	testMongo()
	startupServe()
}
