package main

import (
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// CaseRepository interface should allow cases to be stored and retrieved
type CaseRepository interface {
	Store(cases Case)
	FindById(caseID string) Case
}

// DepartList represents information regarding workflow stations
type DepartList struct {
	Collection    bool `bson:"Collection" json:"Collection"`
	Cutting       bool `bson:"Cutting" json:"Cutting"`
	Diagnosis     bool `bson:"Diagnosis" json:"Diagnosis"`
	DigitalImage  bool `bson:"DigitalImage" json:"DigitalImage"`
	Embedding     bool `bson:"Embedding" json:"Embedding"`
	Grossing      bool `bson:"Grossing" json:"Grossing"`
	Slideprep     bool `bson:"SlidePrep" json:"SlidePrep"`
	Slidetrans    bool `bson:"SlideTrans" json:"SlideTrans"`
	Tissue        bool `bson:"Tissue" json:"Tissue"`
	Transcription bool `bson:"Transcription" json:"Transcription"`
}

// Patient represents information about patient data
type Patient struct {
	MRN        *string `bson:"MRN,omitempty" json:"MRN,omitempty"`
	SSN        *string `bson:"SSN,omitempty" json:"SSN,omitempty"`
	AccountNum *string `bson:"accountNum,omitempty" json:"accountNum,omitempty"`
	Birthdate  *string `bson:"birthdate,omitempty" json:"birthdate,omitempty"`
	First      *string `bson:"first,omitempty" json:"first,omitempty"`
	Last       *string `bson:"last,omitempty" json:"last,omitempty"`
	Middle     *string `bson:"middle,omitempty" json:"middle,omitempty"`
	Phone      *string `bson:"phone,omitempty" json:"phone,omitempty"`
	Race       *string `bson:"race,omitempty" json:"race,omitempty"`
	Sex        *string `bson:"sex,omitempty" json:"sex,omitempty"`
}

// Case represents information about patient cases
type Case struct {
	ID                       bson.ObjectId `bson:",omitempty" json:",omitempty"`
	MOHS                     bool          `bson:"MOHS,omitempty" json:"MOHS,omitempty"`
	NY                       bool          `bson:"NY,omitempty" json:"NY,omitempty"`
	AccessioningCompleteTime int           `bson:"accessioningCompleteTime,omitempty" json:"accessioningCompleteTime,omitempty"`
	Account                  string        `bson:"account,omitempty" json:"account,omitempty"`
	AssignDate               string        `bson:"assignDate,omitempty" json:"assignDate,omitempty"`
	AssignUser               string        `bson:"assignUser,omitempty" json:"assignUser,omitempty"`
	Billing                  string        `bson:"billing,omitempty" json:"billing,omitempty"`
	BillingComplete          bool          `bson:"billingComplete,omitempty" json:"billingComplete,omitempty"`
	BillingPending           bool          `bson:"billingPending,omitempty" json:"billingPending,omitempty"`
	BiopsyDate               string        `bson:"biopsyDate,omitempty" json:"biopsyDate,omitempty"`
	CaseID                   string        `bson:"caseID,omitempty" json:"caseID,omitempty"`
	CaseNote                 string        `bson:"caseNote,omitempty" json:"caseNote,omitempty"`
	CaseType                 string        `bson:"caseType,omitempty" json:"caseType,omitempty"`
	Clinic                   string        `bson:"clinic,omitempty" json:"clinic,omitempty"`
	ClinicAccount            string        `bson:"clinicAccount,omitempty" json:"clinicAccount,omitempty"`
	ClinicName               string        `bson:"clinicName,omitempty" json:"clinicName,omitempty"`
	CompleDay                int           `bson:"compleDay,omitempty" json:"compleDay,omitempty"`
	CompleDigi               int           `bson:"compleDigi,omitempty" json:"compleDigi,omitempty"`
	CreateFrom               string        `bson:"createFrom,omitempty" json:"createFrom,omitempty"`
	Date                     int           `bson:"date,omitempty" json:"date,omitempty"`
	DepartList               *DepartList   `bson:"departList,omitempty" json:"departList,omitempty"`
	DiagnoType               string        `bson:"diagnoType,omitempty" json:"diagnoType,omitempty"`
	DisplayType              *DisplayType  `bson:"displayType,omitempty" json:"displayType,omitempty"`
	DiagnosisFirst           bool          `bson:"diagnosisFirst,omitempty" json:"diagnosisFirst,omitempty"`
	DupSlide                 bool          `bson:"dupSlide,omitempty" json:"dupSlide,omitempty"`
	EndDEPT                  string        `bson:"endDEPT,omitempty" json:"endDEPT,omitempty"`
	ExecuteDate              int           `bson:"executeDate,omitempty" json:"executeDate,omitempty"`
	HaveRelate               bool          `bson:"haveRelate,omitempty" json:"haveRelate,omitempty"`
	OrderFirst               bool          `bson:"orderFirst,omitempty" json:"orderFirst,omitempty"`
	Patient                  *Patient      `bson:"patient,omitempty" json:"patient,omitempty"`
	PhysicianLock            bool          `bson:"physicianLock,omitempty" json:"physicianLock,omitempty"`
	Problem                  bool          `bson:"problem,omitempty" json:"problem,omitempty"`
	RecutProcess             bool          `bson:"recutProcess,omitempty" json:"recutProcess,omitempty"`
	RelateCases              []interface{} `bson:"relateCases,omitempty" json:"relateCases,omitempty"`
	Rush                     bool          `bson:"rush,omitempty" json:"rush,omitempty"`
	RushReq                  string        `bson:"rushReq,omitempty" json:"rushReq,omitempty"`
	SlidePrepCompleDay       int           `bson:"slidePrepCompleDay,omitempty" json:"slidePrepCompleDay,omitempty"`
	Specimens                *Specimens    `bson:"specimens,omitempty" json:"specimens,omitempty"`
	SpecimenSum              int           `bson:"specimenSum,omitempty" json:"specimenSum,omitempty"`
	Status                   string        `bson:"status,omitempty" json:"status,omitempty"`
	Insurance                *Insurance    `bson:"insurance,omitempty" json:"insurance,omitempty"`
	//Extra                    bson.M        `bson:",inline"`
}

// Specimens represents information regarding patient specimens
type Specimens []struct {
	PAS             bool             `bson:"PAS,omitempty" json:"PAS,omitempty"`
	QRCode          string           `bson:"QRCode,omitempty" json:"QRCode,omitempty"`
	SN              string           `bson:"SN,omitempty" json:"SN,omitempty"`
	AdditStain      int              `bson:"additStain,omitempty" json:"additStain,omitempty"`
	AdditStainReq   string           `bson:"additStainReq,omitempty" json:"additStainReq,omitempty"`
	AntomicText     string           `bson:"anatomicText,omitempty" json:"anatomicText,omitempty"`
	Cassette        int              `bson:"cassette,omitempty" json:"cassette,omitempty"`
	Collect         *Collect         `bson:"collect,omitempty" json:"collect,omitempty"`
	CompleDay       int              `bson:"compleDay,omitempty" json:"compleDay,omitempty"`
	DepartList      *DepartList      `bson:"departList,omitempty" json:"departList,omitempty"`
	DiffDiagText    string           `bson:"diffDiagText,omitempty" json:"diffDiagText,omitempty"`
	Gross           *Gross           `bson:"gross,omitempty" json:"gross,omitempty"`
	GrossNote       string           `bson:"grossNote,omitempty" json:"grossNote,omitempty"`
	Hold            bool             `bson:"hold,omitempty" json:"hold,omitempty"`
	Index           int              `bson:"index,omitempty" json:"index,omitempty"`
	Margin          bool             `bson:"margin,omitempty" json:"margin,omitempty"`
	Name            string           `bson:"name,omitempty" json:"name,omitempty"`
	PhistoryText    string           `bson:"phistoryText,omitempty" json:"phistoryText,omitempty"`
	Processing      bool             `bson:"processing,omitempty" json:"processing,omitempty"`
	Recut           bool             `bson:"recut,omitempty" json:"recut,omitempty"`
	Scanned         int              `bson:"scanned,omitempty" json:"scanned,omitempty"`
	Slide           int              `bson:"slide,omitempty" json:"slide,omitempty"`
	SlidePrepDiagno *SlidePrepDiagno `bson:"slidePrepDiagno,omitempty" json:"slidePrepDiagno,omitempty"`
	SlideSum        int              `bson:"slideSum,omitempty" json:"slideSum,omitempty"`
	SourceCode      string           `bson:"sourceCode,omitempty" json:"sourceCode,omitempty"`
	Status          string           `bson:"status,omitempty" json:"status,omitempty"`
	Stop            bool             `bson:"stop,omitempty" json:"stop,omitempty"`
	Type            string           `bson:"type,omitempty" json:"type,omitempty"`
}

// Insurance represents information about patient information for a case
type Insurance []struct {
	Type string `bson:"type,omitempty" json:"type,omitempty"`
}

// SlidePrepDiagno represents information about diagnosis for slide prep cases
type SlidePrepDiagno struct {
	Date  int    `bson:"date,omitempty" json:"date,omitempty"`
	Diag  string `bson:"diag,omitempty" json:"diag,omitempty"`
	Macro string `bson:"macro,omitempty" json:"macro,omitempty"`
	Micro string `bson:"micro,omitempty" json:"micro,omitempty"`
	Note  string `bson:"note,omitempty" json:"note,omitempty"`
}

// DisplayType represents information about types of notifications that are displayed
type DisplayType struct {
	R string `bson:"R,omitempty" json:"R,omitempty"`
	S struct {
		A string `bson:"A,omitempty" json:"A,omitempty"`
		B string `bson:"B,omitempty" json:"B,omitempty"`
		C string `bson:"C,omitempty" json:"C,omitempty"`
		//Extra bson.M `bson:",inline"`
	}
}

// Gross represents gross information in the case
type Gross struct {
	Account   string `bson:"account" json:"account"`
	CutMethod string `bson:"cutMethod" json:"cutMethod"`
	Date      int    `bson:"date" json:"date"`
	Descript  string `bson:"descript" json:"descript"`
	Dimension string `bson:"dimension" json:"dimension"`
	Term      string `bson:"term" json:"term"`
}

// Collect represents information regarding completion time at Accessioning
type Collect struct {
	Account string `bson:"account,omitempty" json:"account,omitempty"`
	Date    int    `bson:"date,omitempty" json:"date,omitempty"`
}

// Collection method returns string with collection name
func (c *Case) Collection() string {
	return "DLCSCase"
}

// Move method returns string with collection name of moved documents
func (c *Case) Move() string {
	return "DLCSCase_removed"
}

// Unique methos returns unique document from MongoDB
func (c *Case) Unique() bson.M {
	return bson.M{"caseID": c.CaseID}
}

//Indexes ensures all indexes are present
func (c *Case) Indexes() []mgo.Index {
	index := mgo.Index{
		Unique:   false,
		DropDups: false,
		Key:      []string{"_id"},
	}
	return []mgo.Index{index}
}

// Cases represents a slice of Case
type Cases []Case

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
	iter := Where(cases, filter).Sort(sortFields).Select(fields).Iter()
	err := iter.All(&results)
	if err != nil {
		handleError(w, 404)
		return
	}
	JSON(w, results, prettyPrint, 200)
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

// CaseShow Return case detail information for individual case
func CaseShow(w http.ResponseWriter, r *http.Request) {
	cases := &Case{}
	prettyPrint := getPrettyPrintValue(r)
	queryFields := getQueryFieldsValue(r)
	caseID := getCaseIDVar(r)
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

// CaseUpdate Update individual case details
func CaseUpdate(w http.ResponseWriter, r *http.Request) {
	cases := &Case{}
	caseID := getCaseIDVar(r)
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

// CaseDelete Delete case detail
func CaseDelete(w http.ResponseWriter, r *http.Request) {
	caseID := getCaseIDVar(r)

	retrievedCase := CaseRetrieve(caseID)
	CaseMove(retrievedCase)

	SoftCassetteDelete(caseID)

	SoftSlideDelete(caseID)

	err := Delete(&Case{CaseID: caseID})
	if err == mgo.ErrNotFound {
		handleError(w, 404)
		return
	}

}

// CaseMove Move case to removed collection
func CaseMove(cases *Case) {

	err := Move(cases)
	if err == mgo.ErrNotFound {
		log.Fatalln(err)
	}
}

// CaseRetrieve retrieves the case
func CaseRetrieve(caseID string) *Case {
	cases := &Case{}
	err := Find(&Case{CaseID: caseID}).One(&cases)
	if err == mgo.ErrNotFound {
		log.Fatal(err)
	}
	return cases
}
