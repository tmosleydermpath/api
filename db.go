package main

import "gopkg.in/mgo.v2"

//const MongoURI = "10.30.178.203:27017/DLCS"
const MongoURI = "tmosley:BackendTest@ds061518.mongolab.com:61518/dlcs_test"

//const MongoURI = "tmosley:BackendTest@ds043927.mongolab.com:43927/dlcs"
//const MongoURI = "tmosley:BackendTest@ds059887.mongolab.com:59887/backend"
const dbname = "dlcs_test"

//const dbname = "DLCS"

var db *mgo.Database

func init() {
	// Connect to Mongo instance
	session, err := mgo.Dial(MongoURI)

	// Check for connection issues
	if err != nil {
		panic(err)
	}

	db = session.DB(dbname)
}
