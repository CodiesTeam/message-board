package mbmux

import (
	"log"
	"net/http"

	"github.com/urfave/negroni"
)

var (
	// SharedMux default serve mux for app.
	SharedMux  *http.ServeMux
	middleware *negroni.Negroni
)

func init() {
	SharedMux = http.NewServeMux()
	middleware = negroni.Classic()
	middleware.UseHandler(SharedMux)
}

// ServeAll startup all serve handlers.
func ServeAll() {
	SharedMux.Handle("/", http.FileServer(http.Dir("../client/")))
	log.Fatal(http.ListenAndServe(":8091", middleware))
}
