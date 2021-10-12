package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sabey/ddd"
	"github.com/sabey/ddd/mock"
)

func TestListUsers_NoJWT(t *testing.T) {
	ts := httptest.NewServer(
		NewHTTPService(
			mock.NewUserRepository(),
		),
	)
	defer ts.Close()

	client := new(http.Client)

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/users", ts.URL), nil)
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

func TestListUsers_InvalidJWT(t *testing.T) {
	ts := httptest.NewServer(
		NewHTTPService(
			mock.NewUserRepository(),
		),
	)
	defer ts.Close()

	client := new(http.Client)

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/users", ts.URL), nil)
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

func TestListUsers_NoUsers(t *testing.T) {
	ts := httptest.NewServer(
		NewHTTPService(
			mock.NewUserRepository(),
		),
	)
	defer ts.Close()

	client := new(http.Client)

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/users", ts.URL), nil)
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

	if resp.StatusCode != 200 {
		t.Errorf("route failed?")
	}

	if string(body) != `{"users":[]}` {
		t.Errorf("unknown body: `%s`", body)
	}
}

func TestListUsers_Success(t *testing.T) {
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

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/users", ts.URL), nil)
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

	if resp.StatusCode != 200 {
		t.Errorf("route failed?")
	}

	if string(body) != `{"users":[{"email":"jackson@juandefu.ca","firstName":"Jackson","lastName":"Sabey"}]}` {
		t.Errorf("unknown body: `%s`", body)
	}
}

func TestListUsers_SuccessMultiple(t *testing.T) {
	mockUsers := mock.NewUserRepository()
	mockUsers.Accounts["jackson@juandefu.ca"] = ddd.User{
		Email:     "jackson@juandefu.ca",
		FirstName: "Jackson",
		LastName:  "Sabey",
		Password:  "pass",
	}
	mockUsers.Accounts["jackson@sabey.co"] = ddd.User{
		Email:     "jackson@sabey.co",
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

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/users", ts.URL), nil)
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

	if resp.StatusCode != 200 {
		t.Errorf("route failed?")
	}

	if string(body) != `{"users":[{"email":"jackson@sabey.co","firstName":"Jackson","lastName":"Sabey"},{"email":"jackson@juandefu.ca","firstName":"Jackson","lastName":"Sabey"}]}` {
		t.Errorf("unknown body: `%s`", body)
	}
}
