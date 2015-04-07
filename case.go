package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type CaseRepository interface {
	Store(cases Case)
	FindById(caseID string) Case
}

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
type Case struct {
	Id                       bson.ObjectId `bson:",omitempty" json:",omitempty"`
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

type Insurance []struct {
	Type string `bson:"type,omitempty" json:"type,omitempty"`
}

type SlidePrepDiagno struct {
	Date  int    `bson:"date,omitempty" json:"date,omitempty"`
	Diag  string `bson:"diag,omitempty" json:"diag,omitempty"`
	Macro string `bson:"macro,omitempty" json:"macro,omitempty"`
	Micro string `bson:"micro,omitempty" json:"micro,omitempty"`
	Note  string `bson:"note,omitempty" json:"note,omitempty"`
}

type DisplayType struct {
	R string `bson:"R,omitempty" json:"R,omitempty"`
	S struct {
		A string `bson:"A,omitempty" json:"A,omitempty"`
		B string `bson:"B,omitempty" json:"B,omitempty"`
		C string `bson:"C,omitempty" json:"C,omitempty"`
		//Extra bson.M `bson:",inline"`
	}
}
type Gross struct {
	Account   string `bson:"account" json:"account"`
	CutMethod string `bson:"cutMethod" json:"cutMethod"`
	Date      int    `bson:"date" json:"date"`
	Descript  string `bson:"descript" json:"descript"`
	Dimension string `bson:"dimension" json:"dimension"`
	Term      string `bson:"term" json:"term"`
}

type Collect struct {
	Account string `bson:"account,omitempty" json:"account,omitempty"`
	Date    int    `bson:"date,omitempty" json:"date,omitempty"`
}

func (c *Case) Collection() string {
	return "DLCSCase"
}

func (c *Case) Unique() bson.M {
	return bson.M{"caseID": c.CaseID}
}

func (c *Case) Indexes() []mgo.Index {
	index := mgo.Index{
		Unique:   false,
		DropDups: false,
		Key:      []string{"_id"},
	}
	return []mgo.Index{index}
}

type Cases []Case
