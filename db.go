package main

import "gopkg.in/mgo.v2"

// const MongoURI URL for the MongoDB instance
const MongoURI = "api:apitest@ds035713.mongolab.com:35713/dlcs"

//const MongoURI = "10.30.177.203:27017/DLCS"

//const MongoURI = "10.30.43.104:27017/DLCS"

//const MongoURI = "tmosley:BackendTest@ds061518.mongolab.com:61518/dlcs_test"

//const MongoURI = "tmosley:BackendTest@ds043927.mongolab.com:43927/dlcs"
//const dbname = "dlcs_test"

const dbname = "dlcs"

//const dbname = "DLCS"

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
