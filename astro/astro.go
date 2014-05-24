package main

import (
	"fmt"
	"github.com/gorilla/mux"
	//"io/ioutil"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"net/http"
)

type Rental struct {
	Name string
}

func main() {

	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/rental/{name}", RentalHandler)
	r.HandleFunc("/{name}", PageHandler)
	r.HandleFunc("/public/{path:.*}", PublicHandler)
	r.HandleFunc("/", PageHandler)
	http.Handle("/", r)

	// Listen and serve requests
	fmt.Printf("Serving requests on port 80...")
	log.Fatal(http.ListenAndServe(":80", nil))
}

func RentalHandler(w http.ResponseWriter, r *http.Request) {
	session, err := mgo.Dial("localhost:27017")

	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("astro").C("rentals")

	result := Rental{}
	err = c.Find(bson.M{"name": "Nader"}).One(&result)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, result.Name)
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
