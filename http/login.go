package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sabey/ddd"
)

/*
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"email": "jackson@juandefu.ca","password": "pass"}' \
  http://localhost:8080/login
*/

func (srv httpService) Login(w http.ResponseWriter, r *http.Request) {
	request := &LoginRequest{}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		// 400
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error":"invalid request"}`)

		return
	}

	if err := request.Validate(); err != nil {
		// 400
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error":"%s"}`, err)

		return
	}

	user, err := srv.userRepo.Login(
		ddd.UserLogin{
			Email:    request.Email,
			Password: ddd.HashPassword(request.Password),
		},
	)
	if err != nil {
		// 400 - we should be able to override this status code but we will use this for now
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error":"%s"}`, err)

		return
	}

	jwt := ddd.SignJWTClaims(user.Email)

	fmt.Fprintf(w, `{"token":"%s"}`, jwt)
}
