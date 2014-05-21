package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/", PageHandler)

	r.HandleFunc("/public/{path:.*}", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		path := params["path"]
		fmt.Print(path + "\n")
		http.ServeFile(w, r, "public/"+path)
	})

	r.HandleFunc("/{name}", PageHandler).Methods("GET")

	http.Handle("/", r)

	// Listen and serve requests
	fmt.Printf("Serving requests on port 80...")
	log.Fatal(http.ListenAndServe(":80", nil))
}

func PageHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	body, err := ioutil.ReadFile("public/" + name + ".html")

	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}

	fmt.Fprintf(w, "%s", body)
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	w.Write([]byte("Hello " + name))
}
