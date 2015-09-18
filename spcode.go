package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// SPCode represents information diagnosis, note and micro codes
type SPCode struct {
	SPCode  string `bson:"code,omitempty" json:"code,omitempty"`
	Account string `bson:"account,omitempty" json:"account,omitempty"`
	Status  string `bson:"status,omitempty" json:"status,omitempty"`
	Text    string `bson:"text,omitempty" json:"text,omitempty"`
	Type    string `bson:"type,omitempty" josn:"type,omitempty"`
}

// Collection method returns collection name
func (c *SPCode) Collection() string {
	return "SlidePrepCode"
}

// Move method returns collection name for moved documents
func (c *SPCode) Move() string {
	return "SPCode_removed"
}

// Unique method returns a unique document from MongoDB
func (c *SPCode) Unique() bson.M {
	return bson.M{"code": c.SPCode}
}

// Indexes ensures that indexes are in place
func (c *SPCode) Indexes() []mgo.Index {
	index := mgo.Index{
		Unique:   true,
		DropDups: false,
		Key:      []string{"_id"},
	}
	return []mgo.Index{index}
}
