package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

type PostData struct {
	key		string
	value	string
}


var theTests = []struct {
	name string
	url string
	methods string
	params []PostData
	expectedStatusCode int
}{
	{"user/Show", "/api/session", "GET", []PostData{}, http.StatusOK},
}
func TestControllers(t *testing.T) {
	routes := Routes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()
    var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
	resp, err := ts.Client().Post(ts.URL + "/api/login", "application/json", bytes.NewBuffer(jsonStr))
	for _, cookie := range resp.Cookies() {
		t.Log("Found a cookie named:", cookie.Name)
	}
	t.Log(err)

	// for _, e := range theTests {
	// 	if e.methods == "GET" {
	// 		resp, err := ts.Client().Get(ts.URL + e.url);if err != nil {
	// 			t.Log(err)
	// 			t.Fatal((err))
	// 		}

	// 		t.Log(resp.Header)

	// 		if resp.StatusCode != e.expectedStatusCode {
	// 			t.Errorf("for %s, expected %d but but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
	// 		}
	// 	}
	// }
}

func TestLogin(t *testing.T) {

}