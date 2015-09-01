package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Model represents information that defines a model
type Model interface {
	Unique() bson.M
	Collection() string
	Move() string
}

// Indexed represents information regarding indexed models
type Indexed interface {
	Indexes() []mgo.Index
	Model
}

// All function returns all documents from a query
func All(m Model) *mgo.Query {
	return Where(m, nil)
}

// Find fuction returns unique documents
func Find(m Model) *mgo.Query {
	return Where(m, m.Unique()).Limit(1)
}

// Where function returns query results
func Where(m Model, q interface{}) *mgo.Query {
	return db.C(m.Collection()).Find(q)
}

// Update function updates query documents
func Update(m Model) (*mgo.ChangeInfo, error) {
	return Find(m).Apply(mgo.Change{
		ReturnNew: true,
		Update: bson.M{
			"$set": m,
		},
	}, m)
}

// Insert creates a new document in the collection
func Insert(m Model) error {
	return db.C(m.Collection()).Insert(m)
}

// Delete removes a document from a collection
func Delete(m Model) error {
	return db.C(m.Collection()).Remove(m.Unique())
}

// Move creates a document in the removed collection
func Move(m Model) error {
	return db.C(m.Move()).Insert(m)
}

// ensureIndexes verifies indexes are in place
func ensureIndexes(m Indexed) {
	coll := db.C(m.Collection())
	for _, i := range m.Indexes() {
		coll.EnsureIndex(i)
	}
}
