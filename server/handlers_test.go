package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddReqFunc(t *testing.T) {
	hh := HttpHandlers{}
	ts := httptest.NewServer(http.HandlerFunc(hh.AddReqFunc))
	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/%s", ts.URL, "add_request"))
	if err != nil {
		t.Error("Get failed.")
	}

	defer resp.Body.Close()

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error("Error while reading response body")
	}

	if string(res) != "Accepted request" {
		t.Error("Response mismatch")
	}

}

func TestGetStatsFunc(t *testing.T) {

	tests := []struct {
		name        string
		AddReqCount int
		want        []string
	}{
		{"Zero", 0, []string{}},
		{"One", 1, []string{"A : 0", "B : 0", "C : 1"}},
		{"Three", 3, []string{"A : 0", "B : 1", "C : 2"}},
		{"Four", 4, []string{"A : 0", "B : 1", "C : 3"}},
		{"Seven", 7, []string{"A : 1", "B : 1", "C : 5"}},
	}

	for _, tt := range tests {
		hh := HttpHandlers{}
		mux := http.NewServeMux()
		mux.HandleFunc("/add_request", hh.AddReqFunc)
		mux.HandleFunc("/get_stats", hh.GetStatsFunc)
		ts := httptest.NewServer(mux)

		//populate buckets
		fmt.Println("ts.URL: ", ts.URL)
		for i := tt.AddReqCount; i > 0; i-- {
			rBody, err := http.Get(fmt.Sprintf("%s/%s", ts.URL, "add_request"))
			if err != nil {
				t.Error("Get failed.")
			}

			defer rBody.Body.Close()
		}

		rBody, err := http.Get(fmt.Sprintf("%s/%s", ts.URL, "get_stats"))
		if err != nil {
			t.Error("Get failed.")
		}

		defer rBody.Body.Close()
		ts.Close()

		r, err := ioutil.ReadAll(rBody.Body)
		if err != nil {
			t.Error("Error while reading response body")
		}

		resp := string(r)

		for _, want := range tt.want {
			if !strings.Contains(resp, want) {
				t.Errorf("Test: %s, want: %v, got: %s\n", tt.name, tt.want, resp)
			}

		}
	}
}
