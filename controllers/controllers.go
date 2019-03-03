package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type appContext struct {
	db *mgo.Database
}
type Tea struct {
	Id       bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string        `json:"name"`
	Category string        `json:"category"`
}

type TeaResource struct {
	Data Tea `json:"data"`
}

type TeaRepo struct {
	coll *mgo.Collection
}

func (r *TeaRepo) Find(id string) (TeaResource, error) {
	result := TeaResource{}
	err := r.coll.FindId(bson.ObjectIdHex(id)).One(&result.Data)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (c *appContext) TeaHandler(w http.ResponseWriter, r *http.Request) {
	params := context.Get(r, "params").(httprouter.Params)
	repo := TeaRepo{c.db.C("teas")}
	trobos, err := repo.Find(params.ByName("id"))
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(trobos)

}
