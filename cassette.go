package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

<<<<<<< HEAD
// CassetteRepository represents information for storing/finding cassettes
=======
>>>>>>> a1e0e468b830fd9f6f6a7d69a13e9b162a07ce21
type CassetteRepository interface {
	Store(cassettes Cassette)
	FindById(qrCode string) Cassette
}
<<<<<<< HEAD

// Cassette represents information regarding cassettes
type Cassette struct {
	ID             bson.ObjectId `bson:",omitempty" json:",omitempty"`
=======
type Cassette struct {
	Id             bson.ObjectId `bson:",omitempty" json:",omitempty"`
>>>>>>> a1e0e468b830fd9f6f6a7d69a13e9b162a07ce21
	QRCode         string        `bson:"QRCode,omitempty" json:"QRCode,omitempty"`
	SN             string        `bson:"SN,omitempty" json:"SN,omitempty"`
	CaseID         string        `bson:"caseID,omitempty" json:"caseID,omitempty"`
	CuttingProcess bool          `bson:"cuttingProcess,omitempty" json:"cuttingProcess,omitempty"`
	Gross          *CasGross     `bson:"gross,omitempty" json:"gross,omitempty"`
	Embedding      *Embedding    `bson:"embedding,omitempty" json:"embedding,omitempty"`
	Pieces         string        `bson:"pieces,omitempty" json:"pieces,omitempty"`
	Specimen       string        `bson:"specimen,omitempty" json:"specimen,omitempty"`
	Tissue         *Tissue       `bson:"tissue,omitempty" json:"tissue,omitempty"`
}

<<<<<<< HEAD
// CasGross represents information for grosser completing gross
=======
>>>>>>> a1e0e468b830fd9f6f6a7d69a13e9b162a07ce21
type CasGross struct {
	Account string `bson:"account,omitempty" json:"account,omitempty"`
	Date    int    `bson:"date,omitempty" json:"date,omitempty"`
}

<<<<<<< HEAD
// Embedding represents information for Embedder completing the embed process
=======
>>>>>>> a1e0e468b830fd9f6f6a7d69a13e9b162a07ce21
type Embedding struct {
	Account string `bson:"account,omitempty" json:"account,omitempty"`
	Date    int    `bson:"date,omitempty" json:"date,omitempty"`
}
<<<<<<< HEAD

// Tissue represents information for grosser that completed grossing
=======
>>>>>>> a1e0e468b830fd9f6f6a7d69a13e9b162a07ce21
type Tissue struct {
	Account string `bson:"account,omitempty" json:"account,omitempty"`
	Date    int    `bson:"date,omitempty" json:"date,omitempty"`
}

<<<<<<< HEAD
// Collection method returns MongoDB collection
=======
>>>>>>> a1e0e468b830fd9f6f6a7d69a13e9b162a07ce21
func (b *Cassette) Collection() string {
	return "cassette"
}

<<<<<<< HEAD
// Move method returns string with collection name of moved documents
func (b *Cassette) Move() string {
	return "cassette_removed"
}

// Unique method returns unique MongoDB document
=======
>>>>>>> a1e0e468b830fd9f6f6a7d69a13e9b162a07ce21
func (b *Cassette) Unique() bson.M {
	return bson.M{"QRCode": b.QRCode}
}

<<<<<<< HEAD
// Indexes ensures necessary indexes are present
=======
>>>>>>> a1e0e468b830fd9f6f6a7d69a13e9b162a07ce21
func (b *Cassette) Indexes() []mgo.Index {
	index := mgo.Index{
		Unique:   true,
		DropDups: false,
		Key:      []string{"_id"},
	}
	return []mgo.Index{index}
}

<<<<<<< HEAD
// Cassettes represents a slice of Cassette
=======
>>>>>>> a1e0e468b830fd9f6f6a7d69a13e9b162a07ce21
type Cassettes []*Cassette
