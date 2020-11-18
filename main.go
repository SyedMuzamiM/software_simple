package main

import (
	"log"
	"net/http"

	"github.com/SyedMuzamiM/software_simple/data"
	"github.com/gorilla/mux"
)

// getProducts returns the products from the data store
func ProductHandler(rw http.ResponseWriter, r *http.Request) {
	// fetch the products from the datastore
	lg := data.GetGrowers()
	log.Println(lg)
	// serialize the list to JSON
	err := lg.ToJSON(lg)
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
