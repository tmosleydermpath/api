package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

<<<<<<< HEAD
// Code represents information diagnosis, note and micro codes
=======
>>>>>>> a1e0e468b830fd9f6f6a7d69a13e9b162a07ce21
type Code struct {
	Code   string `bson:"code,omitempty" json:"code,omitempty"`
	Desc   string `bson:"desc,omitempty" json:"desc,omitempty"`
	Status string `bson:"status,omitempty" json:"status,omitempty"`
	Text   string `bson:"text,omitempty" json:"text,omitempty"`
	Type   string `bson:"type,omitempty" josn:"type,omitempty"`
}

<<<<<<< HEAD
// Collection method returns collection name
=======
>>>>>>> a1e0e468b830fd9f6f6a7d69a13e9b162a07ce21
func (c *Code) Collection() string {
	return "Code"
}

<<<<<<< HEAD
// Move method returns collection name for moved documents
func (c *Code) Move() string {
	return "Code_removed"
}

// Unique method returns a unique document from MongoDB
=======
>>>>>>> a1e0e468b830fd9f6f6a7d69a13e9b162a07ce21
func (c *Code) Unique() bson.M {
	return bson.M{"code": c.Code}
}

<<<<<<< HEAD
// Indexes ensures that indexes are in place
=======
>>>>>>> a1e0e468b830fd9f6f6a7d69a13e9b162a07ce21
func (c *Code) Indexes() []mgo.Index {
	index := mgo.Index{
		Unique:   true,
		DropDups: false,
		Key:      []string{"_id"},
	}
	return []mgo.Index{index}
}
