package main

import (
	"log"
	"net/http"

	"github.com/SyedMuzamiM/software_simple/data"
	"github.com/gorilla/mux"
)

type Growers struct {
	l *log.Logger
}

// getProducts returns the products from the data store
func (g *Growers) ProductHandler(rw http.ResponseWriter, r *http.Request) {
	g.l.Println("Handle GET Products")

	// fetch the products from the datastore
	lp := data.GetGrowers()

	// serialize the list to JSON
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func main() {

	// Maybe I am gonna use the Fiber framework
	r := mux.NewRouter()
	r.HandleFunc("/api/grower", ProductHandler)

	log.Fatal(http.ListenAndServe(":8000", r))
}
