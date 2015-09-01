// @SubApi Case Management API [/cases]
package main

import (
	"encoding/json"
	"fmt"
<<<<<<< HEAD
	"log"
=======
>>>>>>> a1e0e468b830fd9f6f6a7d69a13e9b162a07ce21
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
<<<<<<< HEAD
	caseID := getCaseIDVar(r)
=======
	caseID := getCaseIdVar(r)
>>>>>>> a1e0e468b830fd9f6f6a7d69a13e9b162a07ce21
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

<<<<<<< HEAD
// CaseRetrieve retrieves the case
func CaseRetreive(caseID string) *Case {
	cases := &Case{}
	err := Find(&Case{CaseID: caseID}).One(&cases)
	if err == mgo.ErrNotFound {
		log.Fatal(err)
	}
	return cases
}

// CaseMove Move case to removed collection
func CaseMove(cases *Case) {

	err := Move(cases)
	if err == mgo.ErrNotFound {
		log.Fatalln(err)
	}
}

// CaseDelete Delete case detail
func CaseDelete(w http.ResponseWriter, r *http.Request) {
	caseID := getCaseIDVar(r)

	retrievedCase := CaseRetreive(caseID)
	CaseMove(retrievedCase)

	//retrievedCassette := CassetteRetrieve(caseID)
	retCassID := CassetteIDRetrieve(caseID)
	fmt.Println("Cassette IDs", retCassID)
	//CassetteMove(retrievedCassette)
	CassetteDelete(caseID)
=======
// CaseDelete Delete case detail
func CaseDelete(w http.ResponseWriter, r *http.Request) {
	caseID := getCaseIdVar(r)
>>>>>>> a1e0e468b830fd9f6f6a7d69a13e9b162a07ce21

	err := Delete(&Case{CaseID: caseID})
	if err == mgo.ErrNotFound {
		handleError(w, 404)
		return
	}

}

// CaseUpdate Update individual case details
func CaseUpdate(w http.ResponseWriter, r *http.Request) {
	cases := &Case{}
<<<<<<< HEAD
	caseID := getCaseIDVar(r)
=======
	caseID := getCaseIdVar(r)
>>>>>>> a1e0e468b830fd9f6f6a7d69a13e9b162a07ce21
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

<<<<<<< HEAD
// CassetteRetrieve retrieves the cassettes for a case
func CassetteRetrieve(caseID string) []Cassette {
	//cassettes := &Cassettes{}
	cassette := &Cassette{}
	collection := db.C(cassette.Collection())

	var results []Cassette
	//err := Where(cassette, bson.M{"CaseID": caseID}).All(&results)
	err := collection.Find(bson.M{"caseID": caseID}).All(&results)

	if err == mgo.ErrNotFound {
		log.Fatal(err)
	}
	fmt.Println("Cassette Retrieve", caseID)
	fmt.Println("Cassette Retrieve", results)
	return results
}

// CassetteRetrieve retrieves the cassettes for a case
func CassetteIDRetrieve(caseID string) []Cassette {
	//cassettes := &Cassettes{}
	cassette := &Cassette{}
	collection := db.C(cassette.Collection())

	var results []Cassette
	//err := Where(cassette, bson.M{"CaseID": caseID}).All(&results)
	err := collection.Find(bson.M{"caseID": caseID}).Select(bson.M{"QRCode": 1}).All(&results)

	if err == mgo.ErrNotFound {
		log.Fatal(err)
	}
	//fmt.Println("Cassette Retrieve", caseID)
	//fmt.Println("Cassette Retrieve", results)
	return results
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

// CassetteDelete Moves cassette to new collection and deletes old docs
func CassetteDelete(caseID string) {
	retrievedCassette := CassetteRetrieve(caseID)
	CassetteMove(retrievedCassette)
	log.Println(caseID)

	//for _, c := range retrievedCassette {
	//	err := Delete(&Cassette{CaseID: caseID})
	//	if err == mgo.ErrNotFound {
	//		log.Fatalln(err)
	//	}
	//}

	//retrievedCassettes := CassetteRetreive(caseID)
	//log.Println(retrievedCassettes)
	//CaseMove(retrievedCase)

	err := Delete(&Cassette{CaseID: caseID})
	if err == mgo.ErrNotFound {
		log.Fatal(err)
		return
	}
}

// SingleCassetteDelete Delete information for individual cassette
func SingleCassetteDelete(w http.ResponseWriter, r *http.Request) {
=======
// CassetteDelete Delete information for individual cassette
func CassetteDelete(w http.ResponseWriter, r *http.Request) {
>>>>>>> a1e0e468b830fd9f6f6a7d69a13e9b162a07ce21
	qrCode := getQRCodeVar(r)
	err := Delete(&Cassette{QRCode: qrCode})
	if err == mgo.ErrNotFound {
		handleError(w, 404)
		return
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
<<<<<<< HEAD
	caseID := getCaseIDVar(r)
=======
	caseID := getCaseIdVar(r)
>>>>>>> a1e0e468b830fd9f6f6a7d69a13e9b162a07ce21
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

// SpecimenIndex Return all cassette information for specific case
func SpecimenIndex(w http.ResponseWriter, r *http.Request) {
	cases := &Case{}
<<<<<<< HEAD
	caseID := getCaseIDVar(r)
=======
	caseID := getCaseIdVar(r)
>>>>>>> a1e0e468b830fd9f6f6a7d69a13e9b162a07ce21
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

// SlideIndex Return all slide information for specific case
func SlideIndex(w http.ResponseWriter, r *http.Request) {
	slides := &Slide{}
<<<<<<< HEAD
	caseID := getCaseIDVar(r)
=======
	caseID := getCaseIdVar(r)
>>>>>>> a1e0e468b830fd9f6f6a7d69a13e9b162a07ce21
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

	var results []Code

	err := All(codes).Sort(sortFields).Select(fields).All(&results)
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

	//collection := db.C(cases.Collection())

	var results []Case
<<<<<<< HEAD
	iter := Where(cases, filter).Sort(sortFields).Select(fields).Iter()
	err := iter.All(&results)
=======
	err := Where(cases, filter).Sort(sortFields).Select(fields).All(&results)
	//err := iter.All(&results)
>>>>>>> a1e0e468b830fd9f6f6a7d69a13e9b162a07ce21
	if err != nil {
		handleError(w, 404)
		return
	}
	JSON(w, results, prettyPrint, 200)
}
