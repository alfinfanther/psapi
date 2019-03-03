package structs

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	mgo "gopkg.in/mgo.v2"
)

type Trobosqua struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Category       string             `json:"category,omitempty" bson:"category,omitempty"`
	Title          string             `json:"title,omitempty" bson:"title,omitempty"`
	ImageUrl       string             `json:"image_url,omitempty" bson:"image_url,omitempty"`
	SummaryContent string             `json:"summery_content,omitempty" bson:"summery_content,omitempty"`
	Content        string             `json:"content,omitempty" bson:"content,omitempty"`
	PublishedDate  string             `json:"published_date,omitempty" bson:"published_date,omitempty"`
}

type appContext struct {
	db *mgo.Database
}
