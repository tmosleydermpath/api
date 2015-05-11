package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Code struct {
	Code   string `bson:"code,omitempty" json:"code,omitempty"`
	Desc   string `bson:"desc,omitempty" json:"desc,omitempty"`
	Status string `bson:"status,omitempty" json:"status,omitempty"`
	Text   string `bson:"text,omitempty" json:"text,omitempty"`
	Type   string `bson:"type,omitempty" josn:"type,omitempty"`
}

func (c *Code) Collection() string {
	return "Code"
}

func (c *Code) Unique() bson.M {
	return bson.M{"code": c.Code}
}

func (c *Code) Indexes() []mgo.Index {
	index := mgo.Index{
		Unique:   true,
		DropDups: false,
		Key:      []string{"_id"},
	}
	return []mgo.Index{index}
}
