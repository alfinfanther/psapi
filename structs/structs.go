package structs

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Trobosqua struct {
	Id              bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Title           string        `json:"title"`
	Category        string        `json:"category"`
	Image_url       string        `json:"image_url"`
	Summary_content string        `json:"summary_content"`
	Content         string        `json:"content"`
	Published_date  string        `json:"published_date"`
}

type appContext struct {
	db *mgo.Database
}
