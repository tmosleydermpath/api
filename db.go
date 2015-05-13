package main

import "gopkg.in/mgo.v2"

const MongoURI = "tmosley:BackendTest@ds061518.mongolab.com:61518/dlcs_test"

const dbname = "dlcs_test"

var db *mgo.Database

func init() {
	// Connect to Mongo instance
	session, err := mgo.Dial(MongoURI)

	session.SetBatch(100)
	session.SetPrefetch(0.25)

	// Check for connection issues
	if err != nil {
		panic(err)
	}

	db = session.DB(dbname)
}
