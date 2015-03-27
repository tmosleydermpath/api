package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Cases",
		"GET",
		`/cases/{caseId:[a-zA-Z0-9=\-\/]{14}}`,
		CaseShow,
	},
	Route{
		"CaseUpdate",
		"PATCH",
		`/cases/{caseId:[a-zA-Z0-9=\-\/]{14}}`,
		CaseUpdate,
	},
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
		"AccountIndex",
		"GET",
		"/accounts",
		AccountIndex,
	},
	Route{
		"AccountShow",
		"GET",
		`/accounts/{account}`,
		AccountShow,
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
		"CatchAll",
		"GET",
		`/{path:.*}`,
		Index,
	},
}
