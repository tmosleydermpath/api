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

func CaseShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	query := r.URL.Query()
	prettySelector := query.Get("pretty")
	queryFields := query.Get("fields")

	fields := GetFields(queryFields)
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

func CaseIndex(w http.ResponseWriter, r *http.Request) {
	//query := r.URL.Query()
	prettySelector := r.URL.Query().Get("pretty")
	//queryFields := query.Get("fields")

	//fields := GetFields(queryFields)
	//if queryFields == "" {
	//	fields = nil
	//}

	session := getSession()
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	//collection := session.DB("backend").C("DLCSCase")
	collection := session.DB("DLCS").C("DLCSCase")

	var results []Case
	err := collection.Find(bson.M{}).All(&results)
	if err != nil {
		handleError(w, 404)
		return
	}
	JSON(w, results, prettySelector, 200)
}
