package ddd

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
)

func HashPassword(password string) string {
	h := sha256.New()
	h.Write([]byte(password))
	b := h.Sum(nil)
	return base64.StdEncoding.EncodeToString(b)
}

type User struct {
	Email     string
	FirstName string
	LastName  string
	// this password field is not necessary, but we're using it as storage in the mock
	// this would be necessary if we verified the pw outside of the repos
	Password string
}

type UserRepository interface {
	Create(UserCreate) (*User, error)
	Login(UserLogin) (*User, error)
	List() ([]*User, error)
	Update(UserUpdate) (*User, error)
}

type UserCreate struct {
	Email     string
	FirstName string
	LastName  string
	Password  string
}

func (uc UserCreate) Validate() error {
	if uc.Email == "" {
		return errors.New("email was empty")
	}

	if uc.FirstName == "" {
		return errors.New("firstName was empty")
	}

	if uc.LastName == "" {
		return errors.New("lastName was empty")
	}

	if uc.Password == "" {
		return errors.New("password was empty")
	}

	return nil
}

type UserLogin struct {
	Email    string
	Password string
}

func (ul UserLogin) Validate() error {
	if ul.Email == "" {
		return errors.New("email was empty")
	}

	if ul.Password == "" {
		return errors.New("password was empty")
	}

	return nil
}

type UserUpdate struct {
	// email is required, or an ID to update the correct account
	Email     string
	FirstName string
	LastName  string
}

func (uu UserUpdate) Validate() error {
	if uu.Email == "" {
		return errors.New("email was empty")
	}

	if uu.FirstName == "" {
		return errors.New("firstName was empty")
	}

	if uu.LastName == "" {
		return errors.New("lastName was empty")
	}

	return nil
}
