// @SubApi Case Management API [/cases]
package main

import (
	"encoding/json"
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

func CaseDelete(w http.ResponseWriter, r *http.Request) {
	cases := &Case{}
	prettyPrint := getPrettyPrintValue(r)
	caseId := getCaseIdVar(r)

	collection := db.C(cases.Collection())

	err := collection.Remove(bson.M{"caseID": caseId})
	if err == mgo.ErrNotFound {
		handleError(w, 404)
		return
	}

	JSON(w, cases, prettyPrint, 200)

}
func CaseUpdate(w http.ResponseWriter, r *http.Request) {
	cases := &Case{}
	caseId := getCaseIdVar(r)
	prettyPrint := getPrettyPrintValue(r)
	collection := db.C(cases.Collection())

	json.NewDecoder(r.Body).Decode(&cases)
	//changeInfo := bson.M{&cases}
	change := mgo.Change{
		Update:    bson.M{"$set": &cases},
		ReturnNew: true,
	}

	//result := Case{}

	_, err := collection.Find(bson.M{"caseID": caseId}).Apply(change, &cases)
	if err == mgo.ErrNotFound {
		handleError(w, 404)
		return
	}

	JSON(w, cases, prettyPrint, 200)

}

func CaseInsert(w http.ResponseWriter, r *http.Request) {
	cases := &Case{}
	prettyPrint := getPrettyPrintValue(r)
	collection := db.C(cases.Collection())

	json.NewDecoder(r.Body).Decode(&cases)

	err := collection.Insert(&cases)
	if err == mgo.ErrNotFound {
		handleError(w, 404)
		return
	}

	JSON(w, cases, prettyPrint, 200)

}

func CassetteUpdate(w http.ResponseWriter, r *http.Request) {
	cassettes := &Cassette{}
	qrCode := getQRCodeVar(r)
	prettyPrint := getPrettyPrintValue(r)
	collection := db.C(cassettes.Collection())

	json.NewDecoder(r.Body).Decode(&cassettes)
	//changeInfo := bson.M{&cases}
	change := mgo.Change{
		Update:    bson.M{"$set": &cassettes},
		ReturnNew: true,
	}

	//result := Case{}

	_, err := collection.Find(bson.M{"QRCode": qrCode}).Apply(change, &cassettes)
	if err == mgo.ErrNotFound {
		handleError(w, 404)
		return
	}

	JSON(w, cassettes, prettyPrint, 200)

}

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

func SlideUpdate(w http.ResponseWriter, r *http.Request) {
	slides := &Slide{}
	qrCode := getQRCodeVar(r)
	prettyPrint := getPrettyPrintValue(r)
	collection := db.C(slides.Collection())

	json.NewDecoder(r.Body).Decode(&slides)
	//changeInfo := bson.M{&cases}
	change := mgo.Change{
		Update:    bson.M{"$set": &slides},
		ReturnNew: true,
	}

	//result := Case{}

	_, err := collection.Find(bson.M{"QRCode": qrCode}).Apply(change, &slides)
	if err == mgo.ErrNotFound {
		handleError(w, 404)
		return
	}

	JSON(w, slides, prettyPrint, 200)

}

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
