package http

import "errors"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (lr LoginRequest) Validate() error {
	if lr.Email == "" {
		return errors.New("email was empty")
	}

	if lr.Password == "" {
		return errors.New("password was empty")
	}

	return nil
}

type LoginResponse struct {
	Token string `json:"token"`
}

type SignupRequest struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
}

func (sr SignupRequest) Validate() error {
	if sr.Email == "" {
		return errors.New("email was empty")
	}

	if sr.FirstName == "" {
		return errors.New("firstName was empty")
	}

	if sr.LastName == "" {
		return errors.New("lastName was empty")
	}

	if sr.Password == "" {
		return errors.New("password was empty")
	}

	return nil
}

type SignupResponse struct {
	Token string `json:"token"`
}

type UsersResponse struct {
	Users []UserResponse `json:"users"`
}

type UserResponse struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type UserRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (ur UserRequest) Validate() error {
	if ur.FirstName == "" {
		return errors.New("firstName was empty")
	}

	if ur.LastName == "" {
		return errors.New("lastName was empty")
	}

	return nil
}
