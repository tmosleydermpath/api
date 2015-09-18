// @SubApi Case Management API [/cases]
package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

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

// ClinicShow Return clinic information for individual account
func ClinicShow(w http.ResponseWriter, r *http.Request) {
	clinics := &Clinic{}
	prettyPrint := getPrettyPrintValue(r)
	queryFields := getQueryFieldsValue(r)
	clinicName, err2 := url.QueryUnescape(getClinicVar(r))
	if err2 != nil {
		log.Fatal(err2)
	}
	fields := splitCommaFieldsToMap(queryFields)
	if queryFields == "" {
		fields = nil
	}

	collection := db.C(clinics.Collection())

	err := collection.Find(bson.M{"clinic": clinicName}).Select(fields).One(&clinics)
	if err == mgo.ErrNotFound {
		handleError(w, 404)
		return
	}

	JSON(w, clinics, prettyPrint, 200)

}

// SpecimenIndex Return all cassette information for specific case
func SpecimenIndex(w http.ResponseWriter, r *http.Request) {
	cases := &Case{}
	caseID := getCaseIDVar(r)
	prettyPrint := getPrettyPrintValue(r)
	sortFields := getSortFields(r)
	if sortFields == "" {
		sortFields = " "
	}

	collection := db.C(cases.Collection())

	var results []Case

	//err := All(cases).Sort(sortFields).Select(bson.M{"specimens": 1}).All(&results)

	err := collection.Find(bson.M{"caseID": caseID}).Sort(sortFields).Select(bson.M{"specimens": 1}).All(&results)
	if err != nil {
		fmt.Printf("got an error finding specimen for %s\n", err)
		handleError(w, 404)
		return
	}

	JSON(w, results, prettyPrint, 200)
}

// SpecimenShow Return case detail information for individual case
func SpecimenShow(w http.ResponseWriter, r *http.Request) {
	cases := &Case{}
	prettyPrint := getPrettyPrintValue(r)
	qrCode := getQRCodeVar(r)

	collection := db.C(cases.Collection())

	err := collection.Find(bson.M{"specimens.QRCode": qrCode}).Select(bson.M{"specimens": bson.M{"$elemMatch": bson.M{"QRCode": qrCode}}}).One(&cases)
	if err == mgo.ErrNotFound {
		handleError(w, 404)
		return
	}

	JSON(w, cases, prettyPrint, 200)

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

	var results []Code

	err := All(codes).Sort(sortFields).Select(fields).All(&results)
	if err != nil {
		fmt.Printf("got an error finding code for %s\n", err)
		handleError(w, 404)
		return
	}

	JSON(w, results, prettyPrint, 200)
}

// SPCodeIndex Return all slide information for specific case
func SPCodeIndex(w http.ResponseWriter, r *http.Request) {
	spcodes := &SPCode{}
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

	var results []SPCode

	err := All(spcodes).Sort(sortFields).Select(fields).All(&results)
	if err != nil {
		fmt.Printf("got an error finding slide prep code for %s\n", err)
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

	var results []Account

	err := All(accounts).Sort(sortFields).Select(fields).All(&results)
	if err != nil {
		fmt.Printf("got an error finding account for %s\n", err)
		handleError(w, 404)
		return
	}

	JSON(w, results, prettyPrint, 200)
}

// AccountTypeIndex Return all account information
func AccountTypeIndex(w http.ResponseWriter, r *http.Request) {
	accounts := &Account{}
	prettyPrint := getPrettyPrintValue(r)
	queryFields := getQueryFieldsValue(r)
	accountType := getAccountTypeVar(r)
	sortFields := getSortFields(r)
	if sortFields == "" {
		sortFields = " "
	}

	fields := splitCommaFieldsToMap(queryFields)
	if queryFields == "" {
		fields = nil
	}

	var results []Account

	err := Where(accounts, bson.M{"type": accountType}).Sort(sortFields).Select(fields).All(&results)
	if err != nil {
		fmt.Printf("got an error finding account for %s\n", err)
		handleError(w, 404)
		return
	}

	JSON(w, results, prettyPrint, 200)
}

// ClinicIndex Return all clinic information
func ClinicIndex(w http.ResponseWriter, r *http.Request) {
	clinics := &Clinic{}
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

	var results []Clinic

	err := All(clinics).Sort(sortFields).Select(fields).All(&results)
	if err != nil {
		fmt.Printf("got an error finding clinic for %s\n", err)
		handleError(w, 404)
		return
	}

	JSON(w, results, prettyPrint, 200)
}
