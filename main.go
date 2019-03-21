package main

import (
	"log"
	"net/http"
	"os"
	"psapi/controllers"
	"psapi/global"

	ghandlers "github.com/gorilla/handlers"
	"goji.io"
	"goji.io/pat"
	"gopkg.in/mgo.v2"
)

func ResponseWithJSON(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(json)
}

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	ensureIndex(session)

	router := goji.NewMux()
	router.HandleFunc(pat.Get("/trobosqua"), controllers.AllTrobos(session))
	router.HandleFunc(pat.Get("/trobosqua/id/:id"), controllers.TrobosById(session))
	router.HandleFunc(pat.Get("/trobosqua/category/:category"), controllers.TrobosByCategory(session))
	// http.ListenAndServe("localhost:8090", mux)
	log.Fatal(http.ListenAndServe(global.GetEnv("PORT"), ghandlers.LoggingHandler(os.Stdout, router)))
}

func ensureIndex(s *mgo.Session) {
	session := s.Copy()
	defer session.Close()

	c := session.DB("webscrapper").C("trobosqua")

	index := mgo.Index{
		Key:        []string{"id"},
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
