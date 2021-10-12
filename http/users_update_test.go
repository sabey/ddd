package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/sabey/ddd"
	"github.com/sabey/ddd/mock"
)

func TestUpdateUser_NoJWT(t *testing.T) {
	ts := httptest.NewServer(
		NewHTTPService(
			mock.NewUserRepository(),
		),
	)
	defer ts.Close()

	client := new(http.Client)

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/users", ts.URL), nil)
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

	if resp.StatusCode != 400 {
		t.Errorf("route worked?")
	}

	if string(body) != `{"error":"jwt not found"}` {
		t.Errorf("unknown body: `%s`", body)
	}
}

func TestUpdateUser_InvalidJWT(t *testing.T) {
	ts := httptest.NewServer(
		NewHTTPService(
			mock.NewUserRepository(),
		),
	)
	defer ts.Close()

	client := new(http.Client)

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/users", ts.URL), nil)
	if err != nil {
		t.Errorf("failed to create new http request: %s", err)
	}

	req.Header.Add("X-Authentication-Token", "jwt-token")

	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("failed to make http request: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("failed to read body: %s", err)
	}

	if resp.StatusCode != 400 {
		t.Errorf("route failed?")
	}

	if string(body) != `{"error":"invalid jwt"}` {
		t.Errorf("unknown body: `%s`", body)
	}
}

func TestUpdateUser_InvalidRequest(t *testing.T) {
	ts := httptest.NewServer(
		NewHTTPService(
			mock.NewUserRepository(),
		),
	)
	defer ts.Close()

	client := new(http.Client)

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/users", ts.URL), nil)
	if err != nil {
		t.Errorf("failed to create new http request: %s", err)
	}

	jwt := ddd.SignJWTClaims("jackson@juandefu.ca")

	req.Header.Add("X-Authentication-Token", jwt)

	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("failed to make http request: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("failed to read body: %s", err)
	}

	if resp.StatusCode != 400 {
		t.Errorf("route failed?")
	}

	if string(body) != `{"error":"invalid request"}` {
		t.Errorf("unknown body: `%s`", body)
	}
}

func TestUpdateUser_InvalidRequest_Firstname(t *testing.T) {
	ts := httptest.NewServer(
		NewHTTPService(
			mock.NewUserRepository(),
		),
	)
	defer ts.Close()

	client := new(http.Client)

	reqBody := strings.NewReader(`{}`)

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/users", ts.URL), reqBody)
	if err != nil {
		t.Errorf("failed to create new http request: %s", err)
	}

	jwt := ddd.SignJWTClaims("jackson@juandefu.ca")

	req.Header.Add("X-Authentication-Token", jwt)

	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("failed to make http request: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("failed to read body: %s", err)
	}

	if resp.StatusCode != 400 {
		t.Errorf("route failed?")
	}

	if string(body) != `{"error":"firstName was empty"}` {
		t.Errorf("unknown body: `%s`", body)
	}
}

func TestUpdateUser_InvalidRequest_Lastname(t *testing.T) {
	ts := httptest.NewServer(
		NewHTTPService(
			mock.NewUserRepository(),
		),
	)
	defer ts.Close()

	client := new(http.Client)

	reqBody := strings.NewReader(`{"firstname":"Jackson"}`)

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/users", ts.URL), reqBody)
	if err != nil {
		t.Errorf("failed to create new http request: %s", err)
	}

	jwt := ddd.SignJWTClaims("jackson@juandefu.ca")

	req.Header.Add("X-Authentication-Token", jwt)

	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("failed to make http request: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("failed to read body: %s", err)
	}

	if resp.StatusCode != 400 {
		t.Errorf("route failed?")
	}

	if string(body) != `{"error":"lastName was empty"}` {
		t.Errorf("unknown body: `%s`", body)
	}
}

func TestUpdateUser_Success(t *testing.T) {
	mockUsers := mock.NewUserRepository()
	mockUsers.Accounts["jackson@juandefu.ca"] = ddd.User{
		Email:     "jackson@juandefu.ca",
		FirstName: "Jackson",
		LastName:  "Sabey",
		Password:  "pass",
	}

	ts := httptest.NewServer(
		NewHTTPService(
			mockUsers,
		),
	)
	defer ts.Close()

	client := new(http.Client)

	reqBody := strings.NewReader(`{"firstname":"Jackson","lastname":"Sabey"}`)

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/users", ts.URL), reqBody)
	if err != nil {
		t.Errorf("failed to create new http request: %s", err)
	}

	jwt := ddd.SignJWTClaims("jackson@juandefu.ca")

	req.Header.Add("X-Authentication-Token", jwt)

	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("failed to make http request: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("failed to read body: %s", err)
	}

	if resp.StatusCode != 204 {
		t.Errorf("route failed?")
	}

	if string(body) != `` {
		t.Errorf("unknown body: `%s`", body)
	}

	// verify changes with the list route

	req, err = http.NewRequest("GET", fmt.Sprintf("%s/users", ts.URL), nil)
	if err != nil {
		t.Errorf("failed to create new http request: %s", err)
	}

	req.Header.Add("X-Authentication-Token", jwt)

	resp2, err := client.Do(req)
	if err != nil {
		t.Errorf("failed to make http request: %s", err)
	}
	defer resp2.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("failed to read body: %s", err)
	}

	if resp2.StatusCode != 200 {
		t.Errorf("route failed?")
	}

	if !strings.Contains(`"firstName":"JACKSON"`, string(body)) {
		t.Errorf("firstName was not updated")
	}

	if !strings.Contains(`"lastName":"JACKSON"`, string(body)) {
		t.Errorf("lastName was not updated")
	}
}
