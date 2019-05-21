package middleware

import "testing"

type testRoute struct {
	route     string
	isAllowed bool
}

var testRoutes = []testRoute{
	{"/", true},
	{"/static/css/main.edc01b1d.chunk.css", true},
	{"/static/js/2.0a8cc717.chunk.js", true},
	{"/static/js/main.4ac80d62.chunk.js", true},
	{"/manifest.json", true},
	{"/favicon.ico", true},
	{"/register", false},
	{"/test", false},
	{"/" + AllowedPrefixes[0] + "/test", true},
	{"/" + AllowedPrefixes[1] + "/test", true},
}

func updateTestRoutes() {
	for _, i := range AllowedPrefixes {
		r := testRoute{}
		r.isAllowed = true
		r.route = "/" + i + "/test" // "/api/test"
		testRoutes = append(testRoutes, r)
	}
	for _, i := range AllowedPrefixes {
		r := testRoute{}
		r.isAllowed = false
		r.route = "/" + i + "test" // "/apitest"
		testRoutes = append(testRoutes, r)
	}
}

//TestIsIfAllowedRequest - test for IsIfAllowedRequest func
func TestIsIfAllowedRequest(t *testing.T) {
	updateTestRoutes()
	for _, tr := range testRoutes {
		res := IsIfAllowedRequest(tr.route)
		if res != tr.isAllowed {
			t.Error("For route", tr.route, "test failed")
		}
	}
}
