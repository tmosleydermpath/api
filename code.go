package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Code represents information diagnosis, note and micro codes
type Code struct {
	Code   string `bson:"code,omitempty" json:"code,omitempty"`
	Desc   string `bson:"desc,omitempty" json:"desc,omitempty"`
	Status string `bson:"status,omitempty" json:"status,omitempty"`
	Text   string `bson:"text,omitempty" json:"text,omitempty"`
	Type   string `bson:"type,omitempty" josn:"type,omitempty"`
}

// Collection method returns collection name
func (c *Code) Collection() string {
	return "Code"
}

// Move method returns collection name for moved documents
func (c *Code) Move() string {
	return "Code_removed"
}

// Unique method returns a unique document from MongoDB
func (c *Code) Unique() bson.M {
	return bson.M{"code": c.Code}
}

// Indexes ensures that indexes are in place
func (c *Code) Indexes() []mgo.Index {
	index := mgo.Index{
		Unique:   true,
		DropDups: false,
		Key:      []string{"_id"},
	}
	return []mgo.Index{index}
}
