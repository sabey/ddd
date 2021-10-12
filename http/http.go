package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sabey/ddd"
)

func NewHTTPService(
	userRepo ddd.UserRepository,
) http.Handler {
	return httpService{
		userRepo: userRepo,
	}
}

type httpService struct {
	userRepo ddd.UserRepository
}

func (srv httpService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("ServeHTTP.route: %s %s\n", r.Method, r.URL.Path)

	if r.URL.Path == "/signup" && r.Method == "POST" {
		srv.Signup(w, r)

		return
	} else if r.URL.Path == "/login" && r.Method == "POST" {
		srv.Login(w, r)

		return
	} else if r.URL.Path == "/users" && r.Method == "GET" {
		srv.ListUsers(w, r)

		return
	} else if r.URL.Path == "/users" && r.Method == "PUT" {
		srv.UpdateUser(w, r)

		return
	}
	// 404
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, `{"error":"404"}`)
}
