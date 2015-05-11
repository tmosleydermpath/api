package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Account struct {
	DragonDetection bool       `bson:"DragonDetection,omitempty" json:"DragonDetection,omitempty"`
	NPI             string     `bson:"NPI,omitempty" json:"NPI,omitempty"`
	Account         string     `bson:"account,omitempty" json:"account,omitempty"`
	Addr1           string     `bson:"addr1,omitempty" json:"addr1,omitempty"`
	Addr2           string     `bson:"addr2,omitempty" json:"addr2,omitempty"`
	Aetna           string     `bson:"aetna,omitempty" json:"aetna,omitempty"`
	AssignLimit     int        `bson:"assignLimit,omitempty" json:"assignLimit,omitempty"`
	City            string     `bson:"city,omitempty" json:"city,omitempty"`
	Code            string     `bson:"code,omitempty" json:"code,omitempty"`
	Degree          string     `bson:"degree,omitempty" json:"degree,omitempty"`
	Delay           bool       `bson:"delay,omitempty" json:"delay,omitempty"`
	DelayDate       int        `bson:"delayDate,omitempty" json:"delayDate,omitempty"`
	DiagnoType      string     `bson:"diagnoType,omitempty" json:"diagnoType,omitempty"`
	Diagnosis       *Diagnosis `bson:"diagnosis,omitempty" json:"diagnosis,omitempty"`
	DigitalLimit    int        `bson:"digitalLimit,omitempty" json:"digitalLimit,omitempty"`
	Email           string     `bson:"email,omitempty" json:"email,omitempty"`
	Enable          bool       `bson:"enable,omitempty" json:"enable,omitempty"`
	First           string     `bson:"first,omitempty" json:"first,omitempty"`
	FirstAcc        int        `bson:"firstAcc,omitempty" json:"firstAcc,omitempty"`
	GetComplete     int        `bson:"getComplete,omitempty" json:"getComplete,omitempty"`
	GlassLimit      int        `bson:"glassLimit,omitempty" json:"glassLimit,omitempty"`
	Last            string     `bson:"last,omitempty" json:"last,omitempty"`
	LastAcc         int        `bson:"lastAcc,omitempty" json:"lastAcc,omitempty"`
	LastAssign      int        `bson:"lastAssign,omitempty" json:"lastAssign,omitempty"`
	Lock            bool       `bson:"lock,omitempty" json:"lock,omitempty"`
	LoginCount      int        `bson:"loginCount,omitempty" json:"loginCount,omitempty"`
	MediaCare       bool       `bson:"mediaCare,omitempty" json:"mediaCare,omitempty"`
	Middle          string     `bson:"middle,omitempty" json:"middle,omitempty"`
	Mobile          string     `bson:"mobile,omitempty" json:"mobile,omitempty"`
	Olsen           bool       `bson:"olsen,omitempty" json:"olsen,omitempty"`
	Password        string     `bson:"password,omitempty" json:"password,omitempty"`
	Phone           string     `bson:"phone,omitempty" json:"phone,omitempty"`
	Pic             string     `bson:"pic,omitempty" json:"pic,omitempty"`
	PracticeName    string     `bson:"practicename,omitempty" json:"practicename,omitempty"`
	RePwDate        int        `bson:"rePwDate,omitempty" json:"rePwDate,omitempty"`
	ResetPW         bool       `bson:"resetPW,omitempty" json:"resetPW,omitempty"`
	State           string     `bson:"state,omitempty" json:"state,omitempty"`
	Type            string     `bson:"type,omitempty" json:"type,omitempty"`
	UserCode        string     `bson:"userCode,omitempty" json:"userCode,omitempty"`
	Website         string     `bson:"website,omitempty" json:"website,omitempty"`
	Zip             string     `bson:"zip,omitempty" json:"zip,omitempty"`
}

type Diagnosis []struct {
	Desc  string `bson:"desc,omitempty" json:"desc,omitempty"`
	Diag  string `bson:"diag,omitempty" json:"diag,omitemtpy"`
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
