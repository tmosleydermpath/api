package main

import "net/http"

// Route represents information for url routing
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes represents a slice of Route
type Routes []Route

var routes = Routes{
	Route{"Index", "GET", "/", Index},
	Route{"Cases", "GET", `/cases/{caseId:[a-zA-Z0-9=\-\/]{14}}`, CaseShow},
	Route{"SpecimenShow", "GET", `/specimens/{QRCode}`, SpecimenShow},
	Route{"CaseUpdate", "PATCH", `/cases/{caseId:[a-zA-Z0-9=\-\/]{14}}`, CaseUpdate},
	Route{
		"CaseIndex",
		"GET",
		"/cases",
		CaseIndex,
	},
	Route{
		"CodeIndex",
		"GET",
		"/codes",
		CodeIndex,
	},
	Route{
		"SPCodeIndex",
		"GET",
		"/slideprepcodes",
		SPCodeIndex,
	},
	Route{
		"AccountIndex",
		"GET",
		"/accounts",
		AccountIndex,
	},
	Route{
		"AccountTypeIndex",
		"GET",
		"/accounts/type/{accountType}",
		AccountTypeIndex,
	},
	Route{
		"AccountShow",
		"GET",
		`/accounts/{account}`,
		AccountShow,
	},
	Route{
		"ClinicIndex",
		"GET",
		"/clinics",
		ClinicIndex,
	},
	Route{
		"ClinicShow",
		"GET",
		`/clinics/{clinic}`,
		ClinicShow,
	},
	Route{
		"CaseInsert",
		"POST",
		"/cases",
		CaseInsert,
	},
	Route{
		"CaseDelete",
		"DELETE",
		`/cases/{caseId:[a-zA-Z0-9=\-\/]{14}}`,
		CaseDelete,
	},
	Route{
		"CassetteIndex",
		"GET",
		`/cases/{caseId:[a-zA-Z0-9=\-\/]{14}}/cassettes`,
		CassetteIndex,
	},
	Route{
		"CassetteInsert",
		"POST",
		"/cassettes",
		CassetteInsert,
	},
	Route{
		"CassetteShow",
		"GET",
		`/cassettes/{QRCode:\w{10,14}}`,
		CassetteShow,
	},
	Route{
		"CassetteUpdate",
		"PATCH",
		`/cassettes/{QRCode:\w{10,14}}`,
		CassetteUpdate,
	},
	Route{
		"CassetteDelete",
		"DELETE",
		`/cassettes/{QRCode:\w{10,14}}`,
		CassetteDelete,
	},
	Route{
		"SlideInsert",
		"POST",
		"/slides",
		SlideInsert,
	},
	Route{
		"SlideShow",
		"GET",
		`/slides/{QRCode:\w{10,14}}`,
		SlideShow,
	},
	Route{
		"SlideUpdate",
		"PATCH",
		`/slides/{QRCode:\w{10,14}}`,
		SlideUpdate,
	},
	Route{
		"SlideDelete",
		"DELETE",
		`/slides/{QRCode:\w{10,14}}`,
		SlideDelete,
	},
	Route{
		"SlideIndex",
		"GET",
		`/cases/{caseId:[a-zA-Z0-9=\-\/]{14}}/slides`,
		SlideIndex,
	},
	Route{
		"SpecimenIndex",
		"GET",
		`/cases/{caseId:[a-zA-Z0-9=\-\/]{14}}/specimens`,
		SpecimenIndex,
	},
	Route{
		"CatchAll",
		"GET",
		`/{path:.*}`,
		Index,
	},
	Route{
		"CatchAll",
		"GET",
		`/{path:.*}`,
		HealthIndex,
	},
}
