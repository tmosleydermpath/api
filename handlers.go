// @SubApi Case Management API [/cases]
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func Index(w http.ResponseWriter, r *http.Request) {
	handleError(w, 404)
}

func handleError(w http.ResponseWriter, code int) {
	JSONError(w, Error{codes[code], code}, code)

}
func getVars(r *http.Request) map[string]string {
	return mux.Vars(r)
}

func getFields(r *http.Request, f string) string {
	query := r.URL.Query()
	return query.Get(f)
}

func CaseShow(w http.ResponseWriter, r *http.Request) {
	vars := getVars(r)
	prettySelector := getFields(r, "pretty")
	queryFields := getFields(r, "fields")

	fields := sFields(queryFields)
	if queryFields == "" {
		fields = nil
	}

	caseId := vars["caseId"]
	session := getSession()

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	//collection := session.DB("backend").C("DLCSCase")
	collection := session.DB("DLCS").C("DLCSCase")

	result := Case{}

	err := collection.Find(bson.M{"caseID": caseId}).Select(fields).One(&result)
	if err != nil {
		fmt.Printf("got an error finding a doc %s\n", err)
		handleError(w, 404)
		return
	}

	JSON(w, result, prettySelector, 200)

}

func CassetteShow(w http.ResponseWriter, r *http.Request) {
	vars := getVars(r)
	caseId := vars["caseId"]
	prettySelector := getFields(r, "pretty")
	queryFields := getFields(r, "fields")

	fields := sFields(queryFields)
	if queryFields == "" {
		fields = nil
	}

	session := getSession()
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	collection := session.DB("DLCS").C("cassette")

	var results []Cassette

	err := collection.Find(bson.M{"caseID": caseId}).Select(fields).All(&results)
	if err != nil {
		fmt.Printf("got an error find cassette for %s\n", err)
		handleError(w, 404)
		return
	}

	JSON(w, results, prettySelector, 200)
}

func SlideShow(w http.ResponseWriter, r *http.Request) {
	vars := getVars(r)
	caseId := vars["caseId"]
	prettySelector := getFields(r, "pretty")
	queryFields := getFields(r, "fields")

	fields := sFields(queryFields)
	if queryFields == "" {
		fields = nil
	}

	session := getSession()
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	collection := session.DB("DLCS").C("slide")

	var results []Slide

	err := collection.Find(bson.M{"caseID": caseId}).Select(fields).All(&results)
	if err != nil {
		fmt.Printf("got an error finding slide for %s\n", err)
		handleError(w, 404)
		return
	}

	JSON(w, results, prettySelector, 200)
}
func CaseIndex(w http.ResponseWriter, r *http.Request) {
	prettySelector := getFields(r, "pretty")
	queryFields := getFields(r, "fields")
	filterFields := getFields(r, "filter")
	sortFields := getFields(r, "sort")
	if sortFields == "" {
		sortFields = " "
	}

	fields := sFields(queryFields)
	if queryFields == "" {
		fields = nil
	}
	var filter = bson.M{}
	if filterFields == "" {
		filter = bson.M{}
	} else {
		filter = stationSort(filterFields)
	}

	session := getSession()
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	//collection := session.DB("backend").C("DLCSCase")
	collection := session.DB("DLCS").C("DLCSCase")

	var results []Case
	err := collection.Find(filter).Sort(sortFields).Select(fields).All(&results)
	if err != nil {
		handleError(w, 404)
		return
	}
	JSON(w, results, prettySelector, 200)
}
