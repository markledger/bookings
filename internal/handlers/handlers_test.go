package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}
type testStruct struct {
	Name               string
	Url                string
	Method             string
	Params             []postData
	ExpectedStatusCode int
}

var tests = []testStruct{
	{Name: "home", Url: "/", Method: "GET", Params: []postData{}, ExpectedStatusCode: http.StatusOK},
	{Name: "about", Url: "/about", Method: "GET", Params: []postData{}, ExpectedStatusCode: http.StatusOK},
	{Name: "generals quarters", Url: "/generals-quarters", Method: "GET", Params: []postData{}, ExpectedStatusCode: http.StatusOK},
	{Name: "majors suite", Url: "/majors-suite", Method: "GET", Params: []postData{}, ExpectedStatusCode: http.StatusOK},
	{Name: "search availability", Url: "/search-availability", Method: "GET", Params: []postData{}, ExpectedStatusCode: http.StatusOK},
	{Name: "contact", Url: "/contact", Method: "GET", Params: []postData{}, ExpectedStatusCode: http.StatusOK},
	{Name: "make-reservation", Url: "/make-reservation", Method: "GET", Params: []postData{}, ExpectedStatusCode: http.StatusOK},
	{Name: "reservation summary", Url: "/reservation-summary", Method: "GET", Params: []postData{}, ExpectedStatusCode: http.StatusOK},
	{Name: "post search availability", Url: "/search-availability", Method: "POST", Params: []postData{
		{key: "start", value: "2020-01-01"},
		{key: "end", value: "2020-01-02"},
	}, ExpectedStatusCode: http.StatusOK},
	{Name: "post search availability json", Url: "/search-availability-json", Method: "POST", Params: []postData{
		{key: "start", value: "2020-01-01"},
		{key: "end", value: "2020-01-02"},
	}, ExpectedStatusCode: http.StatusOK},
	{Name: "make reservation", Url: "/make-reservation", Method: "POST", Params: []postData{
		{key: "first_name", value: "Jean-Luc"},
		{key: "last_name", value: "Picard"},
		{key: "email", value: "jean-luc@example.com"},
		{key: "phone", value: "555-555-555"},
	}, ExpectedStatusCode: http.StatusOK},
}

func TestHttpHandlers(t *testing.T) {
	routes := getRoutes()
	testServer := httptest.NewTLSServer(routes)

	defer testServer.Close()

	for _, test := range tests {
		switch test.Method {
		case "GET":
			resp, err := testServer.Client().Get(testServer.URL + test.Url)
			handleTestResponse(resp, err, t, test)

		case "POST":
			values := url.Values{}
			for _, param := range test.Params {
				values.Add(param.key, param.value)
			}
			resp, err := testServer.Client().PostForm(testServer.URL+test.Url, values)
			handleTestResponse(resp, err, t, test)
		}
	}
}

func handleTestResponse(resp *http.Response, err error, t *testing.T, test testStruct) {
	if err != nil {
		t.Log(err)
		t.Fatal(err)
	}

	if resp.StatusCode != test.ExpectedStatusCode {

		t.Errorf("for %s handler returned wrong status code: got %d want %d", test.Name, resp.StatusCode, test.ExpectedStatusCode)
	}
}
