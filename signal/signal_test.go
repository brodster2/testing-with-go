package signal

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
		t.Fatalf("Error creating new request: %s", err)
	}
	Handler(w, r)

	resp := w.Result()

	if resp.StatusCode != 200 {
		t.Fatalf("Handler() response code is %d, wanted 200", resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Handler() content-type = %q; want %q", contentType, "application/json")
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("ioutil.ReadAll(resp.Body) = %s", err)
	}

	var p Person
	err = json.Unmarshal(data, &p)
	if err != nil {
		t.Fatalf("json.Unmarshal(data, &p) = %s", err)
	}

	if p.Age != 21 {
		t.Errorf("p.Age = %d; want %d", p.Age, 21)
	}

	if p.Name != "Brodie Slater" {
		t.Errorf("p.Name = %s; want %s", p.Name, "Brodie Slater")
	}

}
