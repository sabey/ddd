package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sabey/ddd/mock"
)

func Test404(t *testing.T) {
	ts := httptest.NewServer(
		NewHTTPService(
			mock.NewUserRepository(),
		),
	)
	defer ts.Close()

	client := new(http.Client)

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/404", ts.URL), nil)
	if err != nil {
		t.Errorf("failed to create new http request: %s", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("failed to make http request: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("failed to read body: %s", err)
	}

	if resp.StatusCode != 404 {
		t.Errorf("route existed?")
	}

	if string(body) != `{"error":"404"}` {
		t.Errorf("unknown body: `%s`", body)
	}
}
