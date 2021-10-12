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
  --data '{"email": "jackson@juandefu.ca","password": "pass","firstName": "Jackson","lastName": "Sabey"}' \
  http://localhost:8080/signup
*/

func (srv httpService) Signup(w http.ResponseWriter, r *http.Request) {
	request := &SignupRequest{}

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

	user, err := srv.userRepo.Create(
		ddd.UserCreate{
			Email:     request.Email,
			FirstName: request.FirstName,
			LastName:  request.LastName,
			Password:  ddd.HashPassword(request.Password),
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
