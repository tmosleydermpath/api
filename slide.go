package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type SlideRepository interface {
	Store(slides Slide)
	FindById(qrCode string)
}
type Slide struct {
	Id        bson.ObjectId `bson:",omitempty" json:",omitempty"`
	QRCode    string        `bson:"QRCode,omitempty" json:"QRCode,omitempty"`
	SN        string        `bson:"SN,omitempty" json:"SN,omitempty"`
	CaseID    string        `bson:"caseID,omitempty" json:"caseID,omitempty"`
	Cassette  string        `bson:"cassette,omitempty" json:"cassette,omitempty"`
	Cutting   *Cutting      `bson:"cutting,omitempty" json:"cutting,omitempty"`
	Digital   *Digital      `bson:"digital,omitempty" json:"digital,omitempty"`
	DigitalID *DigitalID    `bson:"digitalID,omitempty" json:"digitalID,omitempty"`
	DupGroup  string        `bson:"dupGroup,omitempty" json:"dupGroup,omitempty"`
	Duplicate bool          `bson:"duplicate,omitempty" json:"duplicate,omitempty"`
	ReCut     bool          `bson:"reCut,omitempty" json:"reCut,omitempty"`
	RecutComm string        `bson:"recutComm,omitempty" json:"reCut,omitEmpty"`
}

type Cutting struct {
	Account string `bson:"account,omitempty" json:"account,omitempty"`
	Date    int    `bson:"date,omitempty" json:"date,omitempty"`
}

type Digital struct {
	Account string `bson:"account,omitempty" json:"account,omitempty"`
	Date    int    `bson:"date,omitempty" json:"date,omitempty"`
}

type DigitalID struct {
	AppMag       string `bson:"AppMag,omitempty" json:"AppMag,omitempty"`
	Date         string `bson:"date,omitempty" json:"date,omitempty"`
	Desc         string `bson:"desc,omitempty" json:"desc,omitempty"`
	Height       string `bson:"height,omitempty" json:"height,omitempty"`
	SlideImageID string `bson:"slideImageID,omitempty" json:"slideImageID,omitempty"`
	Title        string `bson:"title,omitempty" json:"title,omitempty"`
	Width        string `bson:"width,omitempty" json:"width,omitempty"`
}

func (s *Slide) Collection() string {
	return "slide"
}

<<<<<<< HEAD
// Move method returns string with collection name of moved documents
func (s *Slide) Move() string {
	return "slide_removed"
}

=======
>>>>>>> a1e0e468b830fd9f6f6a7d69a13e9b162a07ce21
func (s *Slide) Unique() bson.M {
	return bson.M{"QRCode": s.QRCode}
}

func (s *Slide) Indexes() []mgo.Index {
	index := mgo.Index{
		Unique:   false,
		DropDups: false,
		Key:      []string{"_id"},
	}
	return []mgo.Index{index}
}

type Slides []Slide
