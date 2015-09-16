package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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

// CassetteIndex Return all cassette information for specific case
func CassetteIndex(w http.ResponseWriter, r *http.Request) {
	cassettes := &Cassette{}
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
	collection := db.C(cassettes.Collection())

	var results []Cassette

	err := collection.Find(bson.M{"caseID": caseID}).Sort(sortFields).Select(fields).All(&results)

	if err != nil {
		fmt.Printf("got an error find cassette for %s\n", err)
		handleError(w, 404)
		return
	}

	JSON(w, results, prettyPrint, 200)
}

// CassetteInsert Insert new cassette information for individual cassette
func CassetteInsert(w http.ResponseWriter, r *http.Request) {
	cassettes := &Cassette{}
	prettyPrint := getPrettyPrintValue(r)
	collection := db.C(cassettes.Collection())

	json.NewDecoder(r.Body).Decode(&cassettes)

	err := collection.Insert(&cassettes)
	if err == mgo.ErrNotFound {
		handleError(w, 404)
		return
	}

	JSON(w, cassettes, prettyPrint, 200)

}

// CassetteShow Return information for individual cassette
func CassetteShow(w http.ResponseWriter, r *http.Request) {
	cassettes := &Cassette{}
	prettyPrint := getPrettyPrintValue(r)
	queryFields := getQueryFieldsValue(r)
	qrCode := getQRCodeVar(r)
	fields := splitCommaFieldsToMap(queryFields)
	if queryFields == "" {
		fields = nil
	}

	collection := db.C(cassettes.Collection())

	err := collection.Find(bson.M{"QRCode": qrCode}).Select(fields).One(&cassettes)
	if err == mgo.ErrNotFound {
		handleError(w, 404)
		return
	}

	JSON(w, cassettes, prettyPrint, 200)

}

// CassetteUpdate Update cassette information for individual cassette
func CassetteUpdate(w http.ResponseWriter, r *http.Request) {
	cassettes := &Cassette{}
	qrCode := getQRCodeVar(r)
	prettyPrint := getPrettyPrintValue(r)
	collection := db.C(cassettes.Collection())

	json.NewDecoder(r.Body).Decode(&cassettes)
	change := mgo.Change{
		Update:    bson.M{"$set": &cassettes},
		ReturnNew: true,
	}

	_, err := collection.Find(bson.M{"QRCode": qrCode}).Apply(change, &cassettes)
	if err == mgo.ErrNotFound {
		handleError(w, 404)
		return
	}

	JSON(w, cassettes, prettyPrint, 200)

}

// SoftCassetteDelete Moves cassette to new collection and deletes old docs
func SoftCassetteDelete(caseID string) {
	retrievedCassette := CassetteRetrieve(caseID)
	CassetteMove(retrievedCassette)
	log.Println(retrievedCassette)

	for _, c := range retrievedCassette {
		err := Delete(&c)
		if err == mgo.ErrNotFound {
			log.Fatalln(err)
		}
	}
}

// CassetteMove Move cassette document to removed collection
func CassetteMove(cassettes []Cassette) {

	for _, c := range cassettes {
		err := Move(&c)
		if err == mgo.ErrNotFound {
			log.Fatalln(err)
		}
	}
}

// CassetteDelete Delete information for individual cassette
func CassetteDelete(w http.ResponseWriter, r *http.Request) {
	qrCode := getQRCodeVar(r)
	err := Delete(&Cassette{QRCode: qrCode})
	if err == mgo.ErrNotFound {
		handleError(w, 404)
		return
	}
}

// CassetteIDRetrieve retrieves the ID of the cassettes for a case
func CassetteIDRetrieve(caseID string) []Cassette {
	cassette := &Cassette{}
	collection := db.C(cassette.Collection())

	var results []Cassette
	err := collection.Find(bson.M{"caseID": caseID}).Select(bson.M{"QRCode": 1}).All(&results)

	if err == mgo.ErrNotFound {
		log.Fatal(err)
	}
	return results
}

// CassetteRetrieve retrieves the cassettes for a case
func CassetteRetrieve(caseID string) []Cassette {
	cassette := &Cassette{}
	collection := db.C(cassette.Collection())

	var results []Cassette
	err := collection.Find(bson.M{"caseID": caseID}).All(&results)

	if err == mgo.ErrNotFound {
		log.Fatal(err)
	}
	return results
}
