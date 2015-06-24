package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Account struct {
	DragonDetection    bool                   `bson:"dragonDetection,omitempty" json:"dragonDetection,omitempty"`
	NPI                string                 `bson:"NPI,omitempty" json:"NPI,omitempty"`
	Account            string                 `bson:"account,omitempty" json:"account,omitempty"`
	Addr1              string                 `bson:"addr1,omitempty" json:"addr1,omitempty"`
	Addr2              string                 `bson:"addr2,omitempty" json:"addr2,omitempty"`
	Aetna              string                 `bson:"aetna,omitempty" json:"aetna,omitempty"`
	AssignLimit        int                    `bson:"assignLimit,omitempty" json:"assignLimit,omitempty"`
	Billing            string                 `bson:"billing,omitempty" json:"billing,omitempty"`
	City               string                 `bson:"city,omitempty" json:"city,omitempty"`
	CliniPrefer        map[string]interface{} `bson:"cliniPrefer,omitempty" json:"cliniPrefer,omitempty"`
	Code               string                 `bson:"code,omitempty" json:"code,omitempty"`
	Comments           string                 `bson:"comments,omitempty" json:"comments,omitempty"`
	Contact            string                 `bson:"contact,omitempty" json:"contact,omitempty"`
	Degree             string                 `bson:"degree,omitempty" json:"degree,omitempty"`
	Delay              bool                   `bson:"delay,omitempty" json:"delay,omitempty"`
	DelayDate          int                    `bson:"delayDate,omitempty" json:"delayDate,omitempty"`
	DiagnoType         string                 `bson:"diagnoType,omitempty" json:"diagnoType,omitempty"`
	Diagnosis          *Diagnosis             `bson:"diagnosis,omitempty" json:"diagnosis,omitempty"`
	DigitalLimit       int                    `bson:"digitalLimit,omitempty" json:"digitalLimit,omitempty"`
	DupSlide           bool                   `bson:"dupSlide,omitempty" json:"dupSlide,omitempty"`
	Email              string                 `bson:"email,omitempty" json:"email,omitempty"`
	EMR                string                 `bson:"EMR,omitempty" json:"EMR,omitempty"`
	Enable             bool                   `bson:"enable,omitempty" json:"enable,omitempty"`
	First              string                 `bson:"first,omitempty" json:"first,omitempty"`
	FirstAcc           int                    `bson:"firstAcc,omitempty" json:"firstAcc,omitempty"`
	GetComplete        int                    `bson:"getComplete,omitempty" json:"getComplete,omitempty"`
	GetConsultComplete int                    `bson:"getConsultComplete,omitempty" json:"getConsultComplete,omitempty"`
	GlassLimit         int                    `bson:"glassLimit,omitempty" json:"glassLimit,omitempty"`
	Last               string                 `bson:"last,omitempty" json:"last,omitempty"`
	LastAcc            int                    `bson:"lastAcc,omitempty" json:"lastAcc,omitempty"`
	LastAssign         int                    `bson:"lastAssign,omitempty" json:"lastAssign,omitempty"`
	LastNotify         int                    `bson:"lastNotify,omitempty" json:"lastNotify,omitempty"`
	Location           map[string]interface{} `bson:"location,omitempty" json:"location,empty"`
	Lock               bool                   `bson:"lock,omitempty" json:"lock,omitempty"`
	LoginCount         int                    `bson:"loginCount,omitempty" json:"loginCount,omitempty"`
	MediaCare          bool                   `bson:"mediaCare,omitempty" json:"mediaCare,omitempty"`
	Middle             string                 `bson:"middle,omitempty" json:"middle,omitempty"`
	Mobile             string                 `bson:"mobile,omitempty" json:"mobile,omitempty"`
	Olsen              bool                   `bson:"olsen,omitempty" json:"olsen,omitempty"`
	Password           string                 `bson:"password,omitempty" json:"password,omitempty"`
	Phone              string                 `bson:"phone,omitempty" json:"phone,omitempty"`
	PhysiPref          map[string]interface{} `bson:"physiPref,omitempty" json:"physiPref,omitempty"`
	Pic                string                 `bson:"pic,omitempty" json:"pic,omitempty"`
	PracticeName       string                 `bson:"practiceName,omitempty" json:"practiceName,omitempty"`
	PracticeWebsite    string                 `bson:"practiceWebsite,omitempty" json:"practiceWebsite,omitempty"`
	Reports            []string               `bson:"reports,omitempty" json:"reports,omitempty"`
	RePwDate           int                    `bson:"rePwDate,omitempty" json:"rePwDate,omitempty"`
	Requests           string                 `bson:"requests,omitempty" json:"requests,omitempty"`
	ResetPW            bool                   `bson:"resetPW,omitempty" json:"resetPW,omitempty"`
	ShowReport         *ShowReport            `bson:"showReport,omitempty" json:"showReport,omitempty"`
	Specialty          string                 `bson:"specialty,omitempty" json:"specialty,omitempty"`
	State              string                 `bson:"state,omitempty" json:"state,omitempty"`
	Territory          string                 `bson:"territory,omitempty" json:"territory,omitempty"`
	Title              string                 `bson:"title,omitempty" json:"title,omitempty"`
	Type               string                 `bson:"type,omitempty" json:"type,omitempty"`
	UserCode           string                 `bson:"userCode,omitempty" json:"userCode,omitempty"`
	Website            string                 `bson:"website,omitempty" json:"website,omitempty"`
	Zip                string                 `bson:"zip,omitempty" json:"zip,omitempty"`
}

type ShowReport struct {
	Addr      *bool `bson:"Addr,omitempty" json:"Addr,omitempty"`
	CPT       *bool `bson:"CPT,omitempty" json:"CPT,omitemtpy"`
	Deplicate *bool `bson:"Deplicate,omitempty" json:"Deplicate,omitempty"`
	History   *bool `bson:"History,omitempty" json:"History,omitempty"`
	ICD_9     *bool `bson:"ICD-9,omitempty" json:"ICD-9,omitempty"`
	Phone     *bool `bson:"Phone,omitempty" json:"Phone,omitempty"`
	SSN       *bool `bson:"SSN,omitempty" jsoon:"SSN,omitempty"`
}

type Diagnosis []struct {
	Desc  string `bson:"desc,omitempty" json:"desc,omitempty"`
	Diag  string `bson:"diag,omitempty" json:"diag,omitempty"`
	ID    string `bson:"id,omitempty" json:"id,omitempty"`
	Micro string `bson:"micro,omitempty" json:"micro,omitempty"`
}

func (a *Account) Collection() string {
	return "Accounts"
}

func (a *Account) Unique() bson.M {
	return bson.M{"account": a.Account}
}

func (a *Account) Indexes() []mgo.Index {
	index := mgo.Index{
		Unique:   true,
		DropDups: false,
		Key:      []string{"_id"},
	}
	return []mgo.Index{index}
}
