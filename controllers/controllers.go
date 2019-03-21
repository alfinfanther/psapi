package controllers

import (
	"log"
	"net/http"
	"psapi/global"
	"psapi/structs"

	"goji.io/pat"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func AllTrobos(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()

		c := session.DB("webscrapper").C("trobosqua")

		var trobos []structs.Trobosqua
		err := c.Find(bson.M{}).All(&trobos)
		if err != nil {
			global.ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("Failed get all trobosqua: ", err)
			return
		}
		global.RespondWithJSON(w, http.StatusOK, trobos)
	}
}

func TrobosByCategory(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()

		category := pat.Param(r, "category")
		log.Println(category)

		c := session.DB("webscrapper").C("trobosqua")

		var trobos []structs.Trobosqua
		err := c.Find(bson.M{"category": category}).All(&trobos)
		log.Println(err)
		if err != nil {
			global.ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("Failed find trobosqua: ", err)
			return
		}

		global.RespondWithJSON(w, http.StatusOK, trobos)
	}
}

func TrobosById(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()

		id := pat.Param(r, "id")
		log.Println(id)

		c := session.DB("webscrapper").C("trobosqua")

		var trobos structs.Trobosqua
		err := c.FindId(bson.ObjectIdHex(id)).One(&trobos)
		log.Println(err)
		if err != nil {
			global.ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("Failed find trobosqua: ", err)
			return
		}

		global.RespondWithJSON(w, http.StatusOK, trobos)
	}
}
