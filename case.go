package main

import "gopkg.in/mgo.v2/bson"

type DepartList struct {
	Collection    bool `bson:"Collection,omitempty" json:"Collection,omitempty"`
	Cutting       bool `bson:"Cutting,omitempty" json:"Cutting,omitempty"`
	Diagnosis     bool `bson:"Diagnosis,omitempty" json:"Diagnosis,omitempty"`
	DigitalImage  bool `bson:"DigitalImage,omitempty" json:"DigitalImage,omitempty"`
	Embedding     bool `bson:"Embedding,omitempty" json:"Embedding,omitempty"`
	Grossing      bool `bson:"Grossing,omitempty" json:"Grossing,omitempty"`
	Slideprep     bool `bson:"SlidePrep,omitempty" json:"SlidePrep,omitempty"`
	Slidetrans    bool `bson:"SlideTrans,omitempty" json:"SlideTrans,omitempty"`
	Tissue        bool `bson:"Tissue,omitempty" json:"Tissue,omitempty"`
	Transcription bool `bson:"Transcription,omitempty" json:"Transcription,omitempty"`
}

type Patient struct {
	MRN        string `bson:"MRN,omitempty" json:"MRN,omitempty"`
	SSN        string `bson:"SSN,omitempty" json:"SSN,omitempty"`
	AccountNum string `bson:"accountNum,omitempty" json:"accountNum,omitempty"`
	Birthdate  string `bson:"birthdate,omitempty" json:"birthdate,omitempty"`
	First      string `bson:"first,omitempty" json:"first,omitempty"`
	Last       string `bson:"last,omitempty" json:"last,omitempty"`
	Middle     string `bson:"middle,omitempty" json:"middle,omitempty"`
	Phone      string `bson:"phone,omitempty" json:"phone,omitempty"`
	Race       string `bson:"race,omitempty" json:"race,omitempty"`
	Sex        string `bson:"sex,omitempty" json:"sex,omitempty"`
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
	PAS           bool        `json:"PAS"`
	QRCode        string      `bson:"QRCode" json:"QRCode"`
	SN            string      `bson:"SN" json:"SN"`
	AdditStain    int         `json:"additStain"`
	AdditStainReq string      `bson:"additStainReq" json:"additStainReq"`
	AntomicText   string      `bson:"anatomicText" json:"anatomicText"`
	Cassette      int         `json:"cassette"`
	Collect       *Collect    `bson:"collect,omitempty" json:"collect,omitempty"`
	CompleDay     int         `json:"compleDay"`
	DepartList    *DepartList `bson:"departList,omitempty" json:"departList,omitempty"`
	DiffDiagText  string      `bson:"diffDiagText" json:"diffDiagText"`
	Gross         *Gross      `bson:"gross,omitempty" json:"gross,omitempty"`
	GrossNote     string      `bson:"grossNote" json:"grossNote"`
	Hold          bool        `json:"hold"`
	Index         int         `json:"index"`
	Margin        bool        `json:"margin"`
	Name          string      `json:"name"`
	PhistoryText  string      `bson:"phistoryText" json:"phistoryText"`
	Processing    bool        `json:"processing"`
	Recut         bool        `json:"recut"`
	Scanned       int         `json:"scanned"`
	Slide         int         `json:"slide"`
	SlideSum      int         `json:"slideSum"`
	SourceCode    string      `bson:"sourceCode" json:"sourceCode"`
	Status        string      `json:"status"`
	Stop          bool        `json:"stop"`
	Type          string      `json:"type"`
}

type Insurance []struct {
	Type string `json:"type"`
}

type DisplayType struct {
	S struct {
		B string `json:"B"`
		C string `json:"C"`
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
	Account string `bson:"account" json:"account"`
	Date    int    `bson:"date" json:"date"`
}

type Cases []Case
