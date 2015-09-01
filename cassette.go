package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// CassetteRepository represents information for storing/finding cassettes
type CassetteRepository interface {
	Store(cassettes Cassette)
	FindById(qrCode string) Cassette
}

// Cassette represents information regarding cassettes
type Cassette struct {
	ID             bson.ObjectId `bson:",omitempty" json:",omitempty"`
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

// CasGross represents information for grosser completing gross
type CasGross struct {
	Account string `bson:"account,omitempty" json:"account,omitempty"`
	Date    int    `bson:"date,omitempty" json:"date,omitempty"`
}

// Embedding represents information for Embedder completing the embed process
type Embedding struct {
	Account string `bson:"account,omitempty" json:"account,omitempty"`
	Date    int    `bson:"date,omitempty" json:"date,omitempty"`
}

// Tissue represents information for grosser that completed grossing
type Tissue struct {
	Account string `bson:"account,omitempty" json:"account,omitempty"`
	Date    int    `bson:"date,omitempty" json:"date,omitempty"`
}

// Collection method returns MongoDB collection
func (b *Cassette) Collection() string {
	return "cassette"
}

// Move method returns string with collection name of moved documents
func (b *Cassette) Move() string {
	return "cassette_removed"
}

// Unique method returns unique MongoDB document
func (b *Cassette) Unique() bson.M {
	return bson.M{"QRCode": b.QRCode}
}

// Indexes ensures necessary indexes are present
func (b *Cassette) Indexes() []mgo.Index {
	index := mgo.Index{
		Unique:   true,
		DropDups: false,
		Key:      []string{"_id"},
	}
	return []mgo.Index{index}
}

// Cassettes represents a slice of Cassette
type Cassettes []*Cassette
