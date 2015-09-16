package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// SlideRepository Interface should be used to pull slide info
type SlideRepository interface {
	Store(slides Slide)
	FindById(qrCode string)
}

// Slide Will be used to retrieve slide information
type Slide struct {
	ID        bson.ObjectId `bson:",omitempty" json:",omitempty"`
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

// Cutting Information will be retrieved for cutting
type Cutting struct {
	Account string `bson:"account,omitempty" json:"account,omitempty"`
	Date    int    `bson:"date,omitempty" json:"date,omitempty"`
}

// Digital information will be retrieved about scanning
type Digital struct {
	Account string `bson:"account,omitempty" json:"account,omitempty"`
	Date    int    `bson:"date,omitempty" json:"date,omitempty"`
}

// DigitalID information will be retrieved about DigitalID
type DigitalID struct {
	AppMag       string `bson:"AppMag,omitempty" json:"AppMag,omitempty"`
	Date         string `bson:"date,omitempty" json:"date,omitempty"`
	Desc         string `bson:"desc,omitempty" json:"desc,omitempty"`
	Height       string `bson:"height,omitempty" json:"height,omitempty"`
	SlideImageID string `bson:"slideImageID,omitempty" json:"slideImageID,omitempty"`
	Title        string `bson:"title,omitempty" json:"title,omitempty"`
	Width        string `bson:"width,omitempty" json:"width,omitempty"`
}

// Collection will return collection name
func (s *Slide) Collection() string {
	return "slide"
}

// Move method returns string with collection name of moved documents
func (s *Slide) Move() string {
	return "slide_removed"
}

// Unique method will return unique query result
func (s *Slide) Unique() bson.M {
	return bson.M{"QRCode": s.QRCode}
}

// Indexes method will return index
func (s *Slide) Indexes() []mgo.Index {
	index := mgo.Index{
		Unique:   false,
		DropDups: false,
		Key:      []string{"_id"},
	}
	return []mgo.Index{index}
}

// Slides method will return slice of slides
type Slides []Slide

// SlideIndex Return all slide information for specific case
func SlideIndex(w http.ResponseWriter, r *http.Request) {
	slides := &Slide{}
	caseID := getCaseIDVar(r)
	prettyPrint := getPrettyPrintValue(r)
	queryFields := getQueryFieldsValue(r)
	sortFields := getSortFields(r)
	if sortFields == "" {
		sortFields = " "
	}

	fields := splitCommaFieldsToMap(queryFields)
	if queryFields == "" {
		fields = nil
	}

	collection := db.C(slides.Collection())

	var results []Slide
	err := collection.Find(bson.M{"caseID": caseID}).Sort(sortFields).Select(fields).All(&results)
	if err != nil {
		fmt.Printf("got an error finding slide for %s\n", err)
		handleError(w, 404)
		return
	}

	JSON(w, results, prettyPrint, 200)
}

// SlideInsert Insert new information for individual slide
func SlideInsert(w http.ResponseWriter, r *http.Request) {
	slides := &Slide{}
	prettyPrint := getPrettyPrintValue(r)
	collection := db.C(slides.Collection())

	json.NewDecoder(r.Body).Decode(&slides)

	err := collection.Insert(&slides)
	if err == mgo.ErrNotFound {
		handleError(w, 404)
		return
	}

	JSON(w, slides, prettyPrint, 200)

}

// SlideShow Return slide information for individual slide
func SlideShow(w http.ResponseWriter, r *http.Request) {
	slides := &Slide{}
	prettyPrint := getPrettyPrintValue(r)
	queryFields := getQueryFieldsValue(r)
	qrCode := getQRCodeVar(r)
	fields := splitCommaFieldsToMap(queryFields)
	if queryFields == "" {
		fields = nil
	}

	collection := db.C(slides.Collection())

	err := collection.Find(bson.M{"QRCode": qrCode}).Select(fields).One(&slides)
	if err == mgo.ErrNotFound {
		handleError(w, 404)
		return
	}

	JSON(w, slides, prettyPrint, 200)

}

// SlideUpdate Update information for individual slide
func SlideUpdate(w http.ResponseWriter, r *http.Request) {
	slides := &Slide{}
	qrCode := getQRCodeVar(r)
	prettyPrint := getPrettyPrintValue(r)
	collection := db.C(slides.Collection())

	json.NewDecoder(r.Body).Decode(&slides)
	change := mgo.Change{
		Update:    bson.M{"$set": &slides},
		ReturnNew: true,
	}

	_, err := collection.Find(bson.M{"QRCode": qrCode}).Apply(change, &slides)
	if err == mgo.ErrNotFound {
		handleError(w, 404)
		return
	}

	JSON(w, slides, prettyPrint, 200)

}

// SoftSlideDelete Moves slide to new collection and deletes old docs
func SoftSlideDelete(caseID string) {
	retrievedSlide := SlideRetrieve(caseID)
	SlideMove(retrievedSlide)
	log.Println(retrievedSlide)

	for _, c := range retrievedSlide {
		err := Delete(&c)
		if err == mgo.ErrNotFound {
			log.Fatalln(err)
		}
	}
}

// SlideMove Move slide document to removed collection
func SlideMove(slides []Slide) {

	for _, c := range slides {
		err := Move(&c)
		if err == mgo.ErrNotFound {
			log.Fatalln(err)
		}
	}
}

// SlideDelete Delete information for individual slide
func SlideDelete(w http.ResponseWriter, r *http.Request) {
	qrCode := getQRCodeVar(r)
	err := Delete(&Slide{QRCode: qrCode})
	if err == mgo.ErrNotFound {
		handleError(w, 404)
		return
	}
}

// SlideIDRetrieve retrieves the ID of the slides for a case
func SlideIDRetrieve(caseID string) []Slide {
	slide := &Slide{}
	collection := db.C(slide.Collection())

	var results []Slide
	err := collection.Find(bson.M{"caseID": caseID}).Select(bson.M{"QRCode": 1}).All(&results)

	if err == mgo.ErrNotFound {
		log.Fatal(err)
	}
	return results
}

// SlideRetrieve retrieves the slides for a case
func SlideRetrieve(caseID string) []Slide {
	slide := &Slide{}
	collection := db.C(slide.Collection())

	var results []Slide
	err := collection.Find(bson.M{"caseID": caseID}).All(&results)

	if err == mgo.ErrNotFound {
		log.Fatal(err)
	}
	return results
}
