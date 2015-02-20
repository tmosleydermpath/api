package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GetFields(s string) (m map[string]interface{}) {
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

func JSON(w http.ResponseWriter, v interface{}, s string, c int) {
	b, err := JsonIndent(v, s)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(c)
	w.Write(b)
}

func JSONError(w http.ResponseWriter, v interface{}, c int) {
	b, _ := json.MarshalIndent(v, "", "    ")
	w.WriteHeader(c)
	w.Write(b)
}

func JsonIndent(v interface{}, s string) (rj []byte, err error) {
	if len(s) != 0 && s == "false" {
		rj, err := json.Marshal(v)
		return rj, err
	} else {
		rj, err := json.MarshalIndent(v, "", "    ")
		return rj, err
	}
	return rj, err
}
