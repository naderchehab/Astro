package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"net/http"
)

type Movie struct {
	Id      int
	Title   string
	Rank    string
	Year    string
	Rating  string
	Reviews string
}

type Rental struct {
	Name string
}

func main() {

	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/save", SaveHandler).Methods("POST")
	r.HandleFunc("/movies", MoviesHandler).Methods("GET")
	r.HandleFunc("/{name}", PageHandler).Methods("GET")
	r.HandleFunc("/public/{path:.*}", PublicHandler).Methods("GET")
	r.HandleFunc("/", PageHandler).Methods("GET")
	http.Handle("/", r)

	// Listen and serve requests
	fmt.Printf("Serving requests on port 80...")
	log.Fatal(http.ListenAndServe(":80", nil))
}

func PageHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]

	//fmt.Print("In PageHandler. name: " + name + "\n")

	if name == "" {
		name = "index"
	}

	http.ServeFile(w, r, "public/"+name+".html")
}

func PublicHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	path := params["path"]
	//fmt.Print("In PublicHandler. path: " + path + "\n")
	http.ServeFile(w, r, "public/"+path)
}

// List movie
func MoviesHandler(w http.ResponseWriter, r *http.Request) {

	session, err := mgo.Dial("localhost:27017")

	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("astro").C("movies")

	results := []Movie{}
	err = c.Find(nil).All(&results)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(results); i++ {
		fmt.Fprintf(w, results[i].Title+"\n")
	}
}

// Update a movie
func SaveHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	field := r.FormValue("field")
	value := r.FormValue("value")

	// Open mongo connection
	session, err := mgo.Dial("localhost:27017")

	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("astro").C("movies")

	// Update
	colQuerier := bson.M{"id": id}
	change := bson.M{"$set": bson.M{field: value}}
	err = c.Update(colQuerier, change)
	if err != nil {
		panic(err)
	}
}
