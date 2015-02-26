package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Cassette struct {
	QRCode   string    `bson:"QRCode,omitempty" json:"QRCode,omitempty"`
	SN       string    `bson:"SN,omitempty" json:"SN,omitempty"`
	CaseID   string    `bson:"caseID,omitempty" json:"caseID,omitempty"`
	Gross    *CasGross `bson:"gross,omitempty" json:"gross,omitempty"`
	Pieces   string    `bson:"pieces,omitempty" json:"pieces,omitempty"`
	Specimen string    `bson:"specimen,omitempty" json:"specimen,omitempty"`
	Tissue   *Tissue   `bson:"tissue,omitempty" json:"tissue,omitempty"`
}

type CasGross struct {
	Account string `bson:"account,omitempty" json:"account,omitempty"`
	Date    int    `bson:"date,omitempty" json:"date,omitempty"`
}

type Tissue struct {
	Account string `bson:"account,omitempty" json:"account,omitempty"`
	Date    int    `bson:"date,omitempty" json:"date,omitempty"`
}

func (b *Cassette) Collection() string {
	return "cassette"
}

func (b *Cassette) Unique() bson.M {
	return bson.M{"QRCode": b.QRCode}
}

func (b *Cassette) Indexes() []mgo.Index {
	index := mgo.Index{
		Unique:   false,
		DropDups: false,
		Key:      []string{"_id"},
	}
	return []mgo.Index{index}
}

type Cassettes []Cassette
