package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sabey/ddd"
)

/*
curl --header "X-Authentication-Token: jwt-token" --header "Content-Type: application/json" \
  --request PUT \
  --data '{"firstName": "JACKSON","lastName": "SABEY"}' \
  http://localhost:8080/users
*/

func (srv httpService) UpdateUser(w http.ResponseWriter, r *http.Request) {
	jwt := r.Header.Get("X-Authentication-Token")
	if jwt == "" {
		// 400
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error":"jwt not found"}`)

		return
	}

	// validate jwt
	email := ddd.ParseJWTClaims(jwt)
	if email == "" {
		// 400
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error":"invalid jwt"}`)

		return
	}

	defer r.Body.Close()

	request := &UserRequest{}

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

	_, err := srv.userRepo.Update(
		ddd.UserUpdate{
			Email:     email,
			FirstName: request.FirstName,
			LastName:  request.LastName,
		},
	)
	if err != nil {
		// 400 - we should be able to override this status code but we will use this for now
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error":"%s"}`, err)

		return
	}

	w.WriteHeader(http.StatusNoContent)
}
