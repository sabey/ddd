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

func TestSignup_InvalidRequest(t *testing.T) {
	ts := httptest.NewServer(
		NewHTTPService(
			mock.NewUserRepository(),
		),
	)
	defer ts.Close()

	client := new(http.Client)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/signup", ts.URL), nil)
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
		t.Errorf("route failed?")
	}

	if string(body) != `{"error":"invalid request"}` {
		t.Errorf("unknown body: `%s`", body)
	}
}

func TestSignup_InvalidRequest_Email(t *testing.T) {
	ts := httptest.NewServer(
		NewHTTPService(
			mock.NewUserRepository(),
		),
	)
	defer ts.Close()

	client := new(http.Client)

	reqBody := strings.NewReader(`{}`)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/signup", ts.URL), reqBody)
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
		t.Errorf("route failed?")
	}

	if string(body) != `{"error":"email was empty"}` {
		t.Errorf("unknown body: `%s`", body)
	}
}

func TestSignup_InvalidRequest_Firstname(t *testing.T) {
	ts := httptest.NewServer(
		NewHTTPService(
			mock.NewUserRepository(),
		),
	)
	defer ts.Close()

	client := new(http.Client)

	reqBody := strings.NewReader(`{"email":"jackson@juandefu.ca"}`)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/signup", ts.URL), reqBody)
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
		t.Errorf("route failed?")
	}

	if string(body) != `{"error":"firstName was empty"}` {
		t.Errorf("unknown body: `%s`", body)
	}
}

func TestSignup_InvalidRequest_Lastname(t *testing.T) {
	ts := httptest.NewServer(
		NewHTTPService(
			mock.NewUserRepository(),
		),
	)
	defer ts.Close()

	client := new(http.Client)

	reqBody := strings.NewReader(`{"email":"jackson@juandefu.ca","firstname":"Jackson"}`)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/signup", ts.URL), reqBody)
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
		t.Errorf("route failed?")
	}

	if string(body) != `{"error":"lastName was empty"}` {
		t.Errorf("unknown body: `%s`", body)
	}
}

func TestSignup_InvalidRequest_Password(t *testing.T) {
	ts := httptest.NewServer(
		NewHTTPService(
			mock.NewUserRepository(),
		),
	)
	defer ts.Close()

	client := new(http.Client)

	reqBody := strings.NewReader(`{"email":"jackson@juandefu.ca","firstname":"Jackson","lastname":"Sabey"}`)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/signup", ts.URL), reqBody)
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
		t.Errorf("route failed?")
	}

	if string(body) != `{"error":"password was empty"}` {
		t.Errorf("unknown body: `%s`", body)
	}
}

func TestSignup_UserAlreadyExists(t *testing.T) {
	mockUsers := mock.NewUserRepository()
	mockUsers.Accounts["jackson@juandefu.ca"] = ddd.User{
		Email:     "jackson@juandefu.ca",
		FirstName: "Jackson",
		LastName:  "Sabey",
		Password:  ddd.HashPassword("pass"),
	}

	ts := httptest.NewServer(
		NewHTTPService(
			mockUsers,
		),
	)
	defer ts.Close()

	client := new(http.Client)

	reqBody := strings.NewReader(`{"email":"jackson@juandefu.ca","firstname":"Jackson","lastname":"Sabey","password":"pass"}`)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/signup", ts.URL), reqBody)
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
		t.Errorf("route failed?")
	}

	if string(body) != `{"error":"user account already exists"}` {
		t.Errorf("unknown body: `%s`", body)
	}
}

func TestSignup_Success(t *testing.T) {
	ts := httptest.NewServer(
		NewHTTPService(
			mock.NewUserRepository(),
		),
	)
	defer ts.Close()

	client := new(http.Client)

	reqBody := strings.NewReader(`{"email":"jackson@juandefu.ca","firstname":"Jackson","lastname":"Sabey","password":"123"}`)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/signup", ts.URL), reqBody)
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

	if resp.StatusCode != 200 {
		t.Errorf("route failed?")
	}

	if !strings.Contains(string(body), `{"token":"`) {
		t.Errorf("unknown body: `%s`", body)
	}

	// parse jwt
	jwt := string(body[10 : len(body)-2])
	email := ddd.ParseJWTClaims(jwt)

	if email != "jackson@juandefu.ca" {
		t.Errorf("unknown jwt email: `%s`", email)
	}
}
