package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sabey/ddd"
)

/*
curl --header "X-Authentication-Token: jwt-token" \
  http://localhost:8080/users
*/

func (srv httpService) ListUsers(w http.ResponseWriter, r *http.Request) {
	jwt := r.Header.Get("X-Authentication-Token")
	if jwt == "" {
		// 400
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error":"jwt not found"}`)

		return
	}

	// validate jwt
	if email := ddd.ParseJWTClaims(jwt); email == "" {
		// 400
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error":"invalid jwt"}`)

		return
	}

	users, err := srv.userRepo.List()
	if err != nil {
		// 400 - we should be able to override this status code but we will use this for now
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error":"%s"}`, err)

		return
	}

	bs, _ := json.Marshal(UsersResponse{
		Users: newUserResponse(users),
	})

	fmt.Fprintf(w, "%s", bs)
}

func newUserResponse(users []*ddd.User) []UserResponse {
	ur := []UserResponse{}

	for _, user := range users {
		ur = append(ur, UserResponse{
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		})
	}

	return ur
}
