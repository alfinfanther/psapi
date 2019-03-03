package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"goji.io"
	"goji.io/pat"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func ErrorWithJSON(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprintf(w, "{message: %q}", message)
}

func ResponseWithJSON(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(json)
}

type Trobosqua struct {
	Id              bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Ttitle          string        `json:"title"`
	Category        string        `json:"category"`
	Image_url       string        `json:"image_url"`
	Summary_content string        `json:"summary_content"`
	Content         string        `json:"content"`
	Published_date  string        `json:"published_date"`
}

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	ensureIndex(session)

	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/trobosqua"), allTrobos(session))
	mux.HandleFunc(pat.Get("/trobosqua/id/:id"), trobosById(session))
	mux.HandleFunc(pat.Get("/trobosqua/:category"), trobosByCategory(session))
	http.ListenAndServe("localhost:8080", mux)
}

func ensureIndex(s *mgo.Session) {
	session := s.Copy()
	defer session.Close()

	c := session.DB("webscrapper").C("trobosqua")

	index := mgo.Index{
		Key:        []string{"title", "category"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err := c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}

func allTrobos(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()

		c := session.DB("webscrapper").C("trobosqua")

		var trobos []Trobosqua
		err := c.Find(bson.M{}).All(&trobos)
		if err != nil {
			ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("Failed get all trobosqua: ", err)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		enc.Encode(trobos)
	}
}

func trobosByCategory(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()

		category := pat.Param(r, "category")
		log.Println(category)

		c := session.DB("webscrapper").C("trobosqua")

		var trobos []Trobosqua
		err := c.Find(bson.M{"category": category}).All(&trobos)
		log.Println(err)
		if err != nil {
			ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("Failed find trobosqua: ", err)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		enc.Encode(trobos)
	}
}
func trobosById(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()

		id := pat.Param(r, "id")
		log.Println(id)

		c := session.DB("webscrapper").C("trobosqua")

		var trobos Trobosqua
		err := c.FindId(bson.ObjectIdHex(id)).One(&trobos)
		log.Println(err)
		if err != nil {
			ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("Failed find trobosqua: ", err)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		enc.Encode(trobos)
	}
}
