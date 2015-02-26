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
		"CaseIndex",
		"GET",
		"/cases",
		CaseIndex,
	},
	Route{
		"CassetteIndex",
		"GET",
		`/cases/{caseId:[a-zA-Z0-9=\-\/]{14}}/cassettes`,
		CassetteIndex,
	},
	Route{
		"CassetteShow",
		"GET",
		`/cassettes/{QRCode:\w{10,14}}`,
		CassetteShow,
	},
	Route{
		"SlideShow",
		"GET",
		`/slides/{QRCode:\w{10,14}}`,
		SlideShow,
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
