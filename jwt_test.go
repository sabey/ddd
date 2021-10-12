package ddd

import (
	"fmt"
	"testing"
)

func TestJWT(t *testing.T) {
	tokenString := SignJWTClaims("jackson@juandefu.ca")
	if tokenString == "" {
		t.Errorf("tokenString was empty!")
	}

	fmt.Printf("tokenString: %s\n", tokenString)

	email := ParseJWTClaims(tokenString)
	if email == "" {
		t.Errorf("email was empty!")
	}

	fmt.Printf("email: %s\n", email)

	if email != "jackson@juandefu.ca" {
		t.Errorf("unknown token claim email: %s", email)
	}
}
