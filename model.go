package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

<<<<<<< HEAD
// Model represents information that defines a model
type Model interface {
	Unique() bson.M
	Collection() string
	Move() string
}

// Indexed represents information regarding indexed models
=======
type Model interface {
	Unique() bson.M
	Collection() string
}

>>>>>>> a1e0e468b830fd9f6f6a7d69a13e9b162a07ce21
type Indexed interface {
	Indexes() []mgo.Index
	Model
}

<<<<<<< HEAD
// All function returns all documents from a query
=======
>>>>>>> a1e0e468b830fd9f6f6a7d69a13e9b162a07ce21
func All(m Model) *mgo.Query {
	return Where(m, nil)
}

<<<<<<< HEAD
// Find fuction returns unique documents
=======
>>>>>>> a1e0e468b830fd9f6f6a7d69a13e9b162a07ce21
func Find(m Model) *mgo.Query {
	return Where(m, m.Unique()).Limit(1)
}

<<<<<<< HEAD
// Where function returns query results
=======
>>>>>>> a1e0e468b830fd9f6f6a7d69a13e9b162a07ce21
func Where(m Model, q interface{}) *mgo.Query {
	return db.C(m.Collection()).Find(q)
}

<<<<<<< HEAD
// Update function updates query documents
=======
>>>>>>> a1e0e468b830fd9f6f6a7d69a13e9b162a07ce21
func Update(m Model) (*mgo.ChangeInfo, error) {
	return Find(m).Apply(mgo.Change{
		ReturnNew: true,
		Update: bson.M{
			"$set": m,
		},
	}, m)
}

<<<<<<< HEAD
// Insert creates a new document in the collection
=======
>>>>>>> a1e0e468b830fd9f6f6a7d69a13e9b162a07ce21
func Insert(m Model) error {
	return db.C(m.Collection()).Insert(m)
}

<<<<<<< HEAD
// Delete removes a document from a collection
=======
>>>>>>> a1e0e468b830fd9f6f6a7d69a13e9b162a07ce21
func Delete(m Model) error {
	return db.C(m.Collection()).Remove(m.Unique())
}

<<<<<<< HEAD
// Move creates a document in the removed collection
func Move(m Model) error {
	return db.C(m.Move()).Insert(m)
}

// ensureIndexes verifies indexes are in place
=======
>>>>>>> a1e0e468b830fd9f6f6a7d69a13e9b162a07ce21
func ensureIndexes(m Indexed) {
	coll := db.C(m.Collection())
	for _, i := range m.Indexes() {
		coll.EnsureIndex(i)
	}
}
