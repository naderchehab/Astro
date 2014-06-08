package main

import (
	"encoding/json"
	"github.com/go-martini/martini"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"net/http"
	"strconv"
)

type Movie struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Rank    string `json:"rank"`
	Year    string `json:"year"`
	Rating  string `json:"rating"`
	Reviews string `json:"reviews"`
}

func main() {

	// Connect to DB
	session, err := mgo.Dial("localhost:27017")
	session.SetMode(mgo.Strong, true)

	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Serve static files
	m := martini.Classic()

	// GET: List of movies
	m.Get("/movies", func() string {
		c := session.DB("astro").C("movies")
		results := []Movie{}
		err := c.Find(nil).All(&results)
		if err != nil {
			panic(err)
		}
		b, err := json.Marshal(results)
		return string(b)
	})

	// POST: save a movie
	m.Post("/save", func(w http.ResponseWriter, r *http.Request) {
		c := session.DB("astro").C("movies")

		id := r.FormValue("id")
		field := r.FormValue("field")
		value := r.FormValue("value")

		nId, err := strconv.ParseInt(id, 10, 32)

		if err != nil {
			panic(err)
		}

		var movie Movie

		// Update
		err = c.Update(bson.M{"id": nId}, bson.M{"$set": bson.M{field: value}})

		if err != nil {
			panic(err)
		}
	})
	m.Run()
}
