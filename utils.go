package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func stationSort(s string) (m bson.M) {
	switch s {
	case "accessioning":
		return createSortMap("f,f,f,f,f,f,f,f,f")
	case "grossing":
		return createSortMap("t,f,f,f,f,f,f,f,f")
	case "embedding":
		return createSortMap("t,t,t,f,f,f,f,f,f")
	case "cutting":
		return createSortMap("t,t,t,t,f,f,f,f,f")
	case "digitalimage":
		return createSortMap("t,t,t,t,t,f,f,f,f")
	case "transcription":
		return createSortMap("t,t,t,t,t,t,f,f,f")
	case "slidetrans":
		return createSortMap("t,t,t,t,t,t,t,f,f")
	case "microscope":
		return createMicroMap("t,t,t,t,t,t,f,f,f")
	case "slideprep":
		return createSPMap("t,t,t,t,t,t,f,f")

	}
	return m
}
func createSortMap(t string) (m bson.M) {
	sf := splitFields(t)
	c, _ := strconv.ParseBool(sf[0])
	g, _ := strconv.ParseBool(sf[1])
	tis, _ := strconv.ParseBool(sf[2])
	emb, _ := strconv.ParseBool(sf[3])
	cut, _ := strconv.ParseBool(sf[4])
	dig, _ := strconv.ParseBool(sf[5])
	tran, _ := strconv.ParseBool(sf[6])
	stran, _ := strconv.ParseBool(sf[7])
	m = bson.M{
		"departList.Collection":    c,
		"departList.Grossing":      g,
		"departList.Tissue":        tis,
		"departList.Embedding":     emb,
		"departList.Cutting":       cut,
		"departList.DigitalImage":  dig,
		"departList.Transcription": tran,
		"departList.SlideTrans":    stran,
	}
	return m
}

func createSPMap(t string) (m bson.M) {
	sf := splitFields(t)
	c, _ := strconv.ParseBool(sf[0])
	g, _ := strconv.ParseBool(sf[1])
	tis, _ := strconv.ParseBool(sf[2])
	emb, _ := strconv.ParseBool(sf[3])
	cut, _ := strconv.ParseBool(sf[4])
	dig, _ := strconv.ParseBool(sf[5])
	sprep, _ := strconv.ParseBool(sf[6])
	stran, _ := strconv.ParseBool(sf[7])
	m = bson.M{
		"departList.Collection":   c,
		"departList.Grossing":     g,
		"departList.Tissue":       tis,
		"departList.Embedding":    emb,
		"departList.Cutting":      cut,
		"departList.DigitalImage": dig,
		"departList.SlidePrep":    sprep,
		"departList.SlideTrans":   stran,
	}
	return m
}
func createMicroMap(t string) (m bson.M) {
	sf := splitFields(t)
	c, _ := strconv.ParseBool(sf[0])
	g, _ := strconv.ParseBool(sf[1])
	tis, _ := strconv.ParseBool(sf[2])
	emb, _ := strconv.ParseBool(sf[3])
	cut, _ := strconv.ParseBool(sf[4])
	dig, _ := strconv.ParseBool(sf[5])
	tran, _ := strconv.ParseBool(sf[6])
	stran, _ := strconv.ParseBool(sf[7])
	diag, _ := strconv.ParseBool(sf[8])
	m = bson.M{
		"departList.Collection":    c,
		"departList.Grossing":      g,
		"departList.Tissue":        tis,
		"departList.Embedding":     emb,
		"departList.Cutting":       cut,
		"departList.DigitalImage":  dig,
		"departList.Transcription": tran,
		"departList.SlideTrans":    stran,
		"departList.Diagnosis":     bson.M{"$in": []interface{}{diag}},
	}
	return m
}

func splitFields(s string) (sf []string) {
	sf = strings.Split(s, ",")
	return sf
}
func splitCommaFieldsToMap(s string) (m map[string]interface{}) {
	var qf []string
	qf = strings.Split(s, ",")
	if len(qf) != 0 {
		m = make(map[string]interface{})
		for _, v := range qf {
			m[v] = 1
		}
	}
	return m
}

func filteringFields(s string) (m map[string]interface{}) {
	var qf []string
	qf = strings.Split(s, ",")
	if len(qf) != 0 {
		m = make(map[string]interface{})
		for _, v := range qf {
			if qf[0] == "-" {
				m[v] = -1
			} else {
				m[v] = 1
			}
		}
	}
	return m
}

// JSON writes out pretty print JSON
func JSON(w http.ResponseWriter, v interface{}, s string, c int) {
	b, err := JSONIndent(v, s)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(c)
	w.Write(b)
}

// JSONError pretty prints any errors
func JSONError(w http.ResponseWriter, v interface{}, c int) {
	b, _ := json.MarshalIndent(v, "", "    ")
	w.WriteHeader(c)
	w.Write(b)
}

// JSONIndent pretty prints any JSON information
func JSONIndent(v interface{}, s string) (rj []byte, err error) {
	if len(s) != 0 && s == "false" {
		rj, err := json.Marshal(v)
		return rj, err
	}
	rj, err = json.MarshalIndent(v, "", "    ")
	return rj, err

	//return rj, err
}

func getURLVars(r *http.Request) map[string]string {
	return mux.Vars(r)
}

func getFields(r *http.Request, f string) string {
	query := r.URL.Query()
	return query.Get(f)
}

func getVar(r *http.Request, v string) string {
	vars := getURLVars(r)
	return vars[v]
}

func getPrettyPrintValue(r *http.Request) string {
	return getFields(r, "pretty")
}

func getQueryFieldsValue(r *http.Request) string {
	return getFields(r, "fields")
}

func getCaseIDVar(r *http.Request) string {
	return getVar(r, "caseId")
}
func getCodeIDVar(r *http.Request) string {
	return getVar(r, "codeId")
}

func getQRCodeVar(r *http.Request) string {
	return getVar(r, "QRCode")
}
func getAccountVar(r *http.Request) string {
	return getVar(r, "account")
}
func getAccountTypeVar(r *http.Request) string {
	return getVar(r, "accountType")
}
func getFilterFields(r *http.Request) string {
	return getFields(r, "filter")
}
func getSortFields(r *http.Request) string {
	return getFields(r, "sort")
}
