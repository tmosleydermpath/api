package main

type Cassette struct {
	QRCode   string    `bson:"QRCode,omitempty" json:"QRCode,omitempty"`
	SN       string    `bson:"SN,omitempty" json:"SN,omitempty"`
	CaseID   string    `bson:"caseID,omitempty" json:"caseID,omitempty"`
	Gross    *CasGross `bson:"gross,omitempty" json:"gross,omitempty"`
	Pieces   string    `bson:"pieces,omitempty" json:"pieces,omitempty"`
	Specimen string    `bson:"specimen,omitempty" json:"specimen,omitempty"`
	Tissue   *Tissue   `bson:"tissue,omitempty" json:"tissue,omitempty"`
}

type CasGross struct {
	Account string `bson:"account,omitempty" json:"account,omitempty"`
	Date    int    `bson:"date,omitempty" json:"date,omitempty"`
}

type Tissue struct {
	Account string `bson:"account,omitempty" json:"account,omitempty"`
	Date    int    `bson:"date,omitempty" json:"date,omitempty"`
}

type Cassettes []Cassette
