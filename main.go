package main

import (
	"fmt"
	"github.com/gorilla/mux"
	//"io/ioutil"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/{name}", PageHandler)
	r.HandleFunc("/public/{path:.*}", PublicHandler)
	r.HandleFunc("/", PageHandler)
	http.Handle("/", r)

	// Listen and serve requests
	fmt.Printf("Serving requests on port 80...")
	log.Fatal(http.ListenAndServe(":80", nil))
}

func PageHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]

	fmt.Print("In PageHandler. name: " + name + "\n")

	if name == "" {
		name = "index"
	}

	http.ServeFile(w, r, "public/"+name+".html")
}

func PublicHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	path := params["path"]
	fmt.Print("In PublicHandler. path: " + path + "\n")
	http.ServeFile(w, r, "public/"+path)
}
