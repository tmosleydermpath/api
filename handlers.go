// @SubApi Case Management API [/cases]
package main

import (
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func Index(w http.ResponseWriter, r *http.Request) {
	handleError(w, 404)
}

func handleError(w http.ResponseWriter, code int) {
	JSONError(w, Error{codes[code], code}, code)

}

func CaseShow(w http.ResponseWriter, r *http.Request) {
	cases := &Case{}
	prettyPrint := getPrettyPrintValue(r)
	queryFields := getQueryFieldsValue(r)
	caseId := getCaseIdVar(r)
	fields := splitCommaFieldsToMap(queryFields)
	if queryFields == "" {
		fields = nil
	}

	collection := db.C(cases.Collection())

	//result := Case{}

	err := collection.Find(bson.M{"caseID": caseId}).Select(fields).One(&cases)
	if err == mgo.ErrNotFound {
		handleError(w, 404)
		return
	}

	JSON(w, cases, prettyPrint, 200)

}

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

	//result := Case{}

	err := collection.Find(bson.M{"QRCode": qrCode}).Select(fields).One(&cassettes)
	if err == mgo.ErrNotFound {
		handleError(w, 404)
		return
	}

	JSON(w, cassettes, prettyPrint, 200)

}

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

	//result := Case{}

	err := collection.Find(bson.M{"QRCode": qrCode}).Select(fields).One(&slides)
	if err == mgo.ErrNotFound {
		handleError(w, 404)
		return
	}

	JSON(w, slides, prettyPrint, 200)

}
func CassetteIndex(w http.ResponseWriter, r *http.Request) {
	cassettes := &Cassette{}
	caseId := getCaseIdVar(r)
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

	err := collection.Find(bson.M{"caseID": caseId}).Sort(sortFields).Select(fields).All(&results)
	if err != nil {
		fmt.Printf("got an error find cassette for %s\n", err)
		handleError(w, 404)
		return
	}

	JSON(w, results, prettyPrint, 200)
}

func SlideIndex(w http.ResponseWriter, r *http.Request) {
	slides := &Slide{}
	caseId := getCaseIdVar(r)
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

	err := collection.Find(bson.M{"caseID": caseId}).Sort(sortFields).Select(fields).All(&results)
	if err != nil {
		fmt.Printf("got an error finding slide for %s\n", err)
		handleError(w, 404)
		return
	}

	JSON(w, results, prettyPrint, 200)
}
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
