package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Clinic struct {
	LKKEY          string     `bson:"LK_KEY,omitempty" json:"LK_KEY,omitempty"`
	Active         bool       `bson:"active,omitempty" json:"active,omitempty"`
	Addr1          string     `bson:"addr1,omitempty" json:"addr1,omitempty"`
	Addr2          string     `bson:"addr2,omitempty" json:"addr2,omitempty"`
	City           string     `bson:"city,omitempty" json:"city,omitempty"`
	Clinic         string     `bson:"clinic,omitempty" json:"clinic,omitempty"`
	ContactEMail   string     `bson:"contactEMail,omitempty" json:"contactEMail,omitempty"`
	ContactName    string     `bson:"contactName,omitempty" json:"contactName,omitempty"`
	ContactPhone   string     `bson:"contactPhone,omitempty" json:"contactPhone,omitempty"`
	ContactTitle   string     `bson:"contactTitle,omitempty" json:"contactTitle,omitempty"`
	Enable         bool       `bson:"enable,omitempty" json:"enable,omitempty"`
	Fax            string     `bson:"fax,omitempty" json:"fax,omitempty"`
	LocationCode   string     `bson:"locationCode,omitempty" json:"locationCode,omitempty"`
	NewClinic      bool       `bson:"newClinic,omitempty" json:"newClinic,omitempty"`
	Phone          string     `bson:"phone,omitempty" json:"phone,omitempty"`
	PickupDay      *PickupDay `bson:"pickupDay,omitempty" json:"pickupDay,omitempty"`
	PickupParcel   string     `bson:"pickupParcel,omitempty" json:"pickupParcel,omitempty"`
	PickupPoint    string     `bson:"pickupPoint,omitempty" json:"pickupPoint,omitempty"`
	PickupSchedule string     `bson:"pickupSchedule,omitempty" json:"pickupSchedule,omitempty"`
	State          string     `bson:"state,omitempty" json:"state,omitempty"`
	Zip            string     `bson:"zip,omitempty" json:"zip,omitempty"`
}

type PickupDay struct {
	Fri   string `bson:"Fri,omitempty" json:"Fri,omitempty"`
	Mon   string `bson:"Mon,omitempty" json:"Mon,omitempty"`
	Other string `bson:"Other,omitempty" json:"Other,omitempty"`
	Thu   string `bson:"Thu,omitempty" json:"Thu,omitempty"`
	Tue   string `bson:"Tue,omitempty" json:"Tue,omitempty"`
	Wed   string `bson:"Wed,omitempty" json:"Wed,omitempty"`
}

func (c *Clinic) Collection() string {
	return "Clinics"
}

func (c *Clinic) Unique() bson.M {
	return bson.M{"clinic": c.Clinic}
}

func (c *Clinic) Indexes() []mgo.Index {
	index := mgo.Index{
		Unique:   true,
		DropDups: false,
		Key:      []string{"_id"},
	}
	return []mgo.Index{index}
}
