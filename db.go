package main

import "gopkg.in/mgo.v2"

// const MongoURI URL for the MongoDB instance
//const MongoURI = "tmosley:BackendTest@ds035713.mongolab.com:35713/dlcs"

const MongoURI = "10.30.178.203:27017/DLCS_SCRUBBED"

//const MongoURI = "10.30.43.104:27017/DLCS_DEV"

//const MongoURI = "tmosley:BackendTest@ds061518.mongolab.com:61518/dlcs_test"

//const MongoURI = "tmosley:BackendTest@ds043927.mongolab.com:43927/dlcs"
//const dbname = "dlcs_test"

//const dbname = "dlcs"

const dbname = "DLCS_SCRUBBED"

var db *mgo.Database

func init() {
	// Connect to Mongo instance
	session, err := mgo.Dial(MongoURI)

	//session.SetBatch(250)
	//session.SetPrefetch(0.25)

	// Check for connection issues
	if err != nil {
		panic(err)
	}

	db = session.DB(dbname)
}
