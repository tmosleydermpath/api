// @SubApi Case Management API [/cases]
package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Index Return 404 error when going to host root
func Index(w http.ResponseWriter, r *http.Request) {
	handleError(w, 404)
}

// handleError Receive error codes and display them
func handleError(w http.ResponseWriter, code int) {
	JSONError(w, Error{codes[code], code}, code)

}

// CaseShow Return case detail information for individual case
func CaseShow(w http.ResponseWriter, r *http.Request) {
	cases := &Case{}
	prettyPrint := getPrettyPrintValue(r)
	queryFields := getQueryFieldsValue(r)
	caseID := getCaseIdVar(r)
	fields := splitCommaFieldsToMap(queryFields)
	if queryFields == "" {
		fields = nil
	}

	collection := db.C(cases.Collection())

	err := collection.Find(bson.M{"caseID": caseID}).Select(fields).One(&cases)
	if err == mgo.ErrNotFound {
		handleError(w, 404)
		return
	}

	JSON(w, cases, prettyPrint, 200)

}

// CaseDelete Delete case detail
func CaseDelete(w http.ResponseWriter, r *http.Request) {
	cases := &Case{}
	prettyPrint := getPrettyPrintValue(r)
	caseID := getCaseIdVar(r)

	collection := db.C(cases.Collection())

	err := collection.Remove(bson.M{"caseID": caseID})
	if err == mgo.ErrNotFound {
		handleError(w, 404)
		return
	}

	JSON(w, cases, prettyPrint, 200)

}

// CaseUpdate Update individual case details
func CaseUpdate(w http.ResponseWriter, r *http.Request) {
	cases := &Case{}
	caseID := getCaseIdVar(r)
	prettyPrint := getPrettyPrintValue(r)
	collection := db.C(cases.Collection())

	json.NewDecoder(r.Body).Decode(&cases)
	change := mgo.Change{
		Update:    bson.M{"$set": &cases},
		ReturnNew: true,
	}

	_, err := collection.Find(bson.M{"caseID": caseID}).Apply(change, &cases)
	if err == mgo.ErrNotFound {
		handleError(w, 404)
		return
	}

	JSON(w, cases, prettyPrint, 200)

}

// CaseInsert Insert new case detail for individual case
func CaseInsert(w http.ResponseWriter, r *http.Request) {
	cases := &Case{}
	prettyPrint := getPrettyPrintValue(r)
	collection := db.C(cases.Collection())

	json.NewDecoder(r.Body).Decode(&cases)

	err := collection.Insert(&cases)
	if err == mgo.ErrNotFound {
		handleError(w, 405)
		return
	}

	JSON(w, cases, prettyPrint, 200)

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

// CassetteDelete Delete information for individual cassette
func CassetteDelete(w http.ResponseWriter, r *http.Request) {
	cassettes := &Cassette{}
	qrCode := getQRCodeVar(r)
	prettyPrint := getPrettyPrintValue(r)
	collection := db.C(cassettes.Collection())

	json.NewDecoder(r.Body).Decode(&cassettes)

	err := collection.Remove(bson.M{"QRCode": qrCode})
	if err == mgo.ErrNotFound {
		handleError(w, 404)
		return
	}

	JSON(w, cassettes, prettyPrint, 200)

}

// SlideDelete Delete information for individual slide
func SlideDelete(w http.ResponseWriter, r *http.Request) {
	slides := &Slide{}
	qrCode := getQRCodeVar(r)
	prettyPrint := getPrettyPrintValue(r)
	collection := db.C(slides.Collection())

	json.NewDecoder(r.Body).Decode(&slides)

	err := collection.Remove(bson.M{"QRCode": qrCode})
	if err == mgo.ErrNotFound {
		handleError(w, 404)
		return
	}

	JSON(w, slides, prettyPrint, 200)

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

// AccountShow Return account information for individual account
func AccountShow(w http.ResponseWriter, r *http.Request) {
	accounts := &Account{}
	prettyPrint := getPrettyPrintValue(r)
	queryFields := getQueryFieldsValue(r)
	accountName := getAccountVar(r)
	fields := splitCommaFieldsToMap(queryFields)
	if queryFields == "" {
		fields = nil
	}

	collection := db.C(accounts.Collection())

	err := collection.Find(bson.M{"account": accountName}).Select(fields).One(&accounts)
	if err == mgo.ErrNotFound {
		handleError(w, 404)
		return
	}

	JSON(w, accounts, prettyPrint, 200)

}

// CassetteIndex Return all cassette information for specific case
func CassetteIndex(w http.ResponseWriter, r *http.Request) {
	cassettes := &Cassette{}
	caseID := getCaseIdVar(r)
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

// SlideIndex Return all slide information for specific case
func SlideIndex(w http.ResponseWriter, r *http.Request) {
	slides := &Slide{}
	caseID := getCaseIdVar(r)
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

// CodeIndex Return all slide information for specific case
func CodeIndex(w http.ResponseWriter, r *http.Request) {
	codes := &Code{}
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
	var filter = bson.M{}

	collection := db.C(codes.Collection())

	var results []Code

	err := collection.Find(filter).Sort(sortFields).Select(fields).All(&results)
	if err != nil {
		fmt.Printf("got an error finding code for %s\n", err)
		handleError(w, 404)
		return
	}

	JSON(w, results, prettyPrint, 200)
}

// AccountIndex Return all account information
func AccountIndex(w http.ResponseWriter, r *http.Request) {
	accounts := &Account{}
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
	var filter = bson.M{}

	collection := db.C(accounts.Collection())

	var results []Account

	err := collection.Find(filter).Sort(sortFields).Select(fields).All(&results)
	if err != nil {
		fmt.Printf("got an error finding account for %s\n", err)
		handleError(w, 404)
		return
	}

	JSON(w, results, prettyPrint, 200)
}

// CaseIndex Return case detail information for all cases
func CaseIndex(w http.ResponseWriter, r *http.Request) {
	cases := &Case{}
	prettyPrint := getPrettyPrintValue(r)
	queryFields := getQueryFieldsValue(r)
	filterFields := getFilterFields(r)
	sortFields := getSortFields(r)
	if sortFields == "" {
		sortFields = " "
	}

	fields := splitCommaFieldsToMap(queryFields)
	if queryFields == "" {
		fields = nil
	}
	var filter = bson.M{}
	if filterFields == "" {
		filter = bson.M{}
	} else {
		filter = stationSort(filterFields)
	}

	collection := db.C(cases.Collection())

	var results []Case
	err := collection.Find(filter).Sort(sortFields).Select(fields).All(&results)
	if err != nil {
		handleError(w, 404)
		return
	}
	JSON(w, results, prettyPrint, 200)
}
