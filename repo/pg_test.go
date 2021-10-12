package repo

import (
	"testing"

	"github.com/sabey/ddd"
)

var (
	repoOpts = RepositoryOpts{
		Addr:     "192.168.2.214:5432",
		User:     "postgres",
		Password: "postgres",
		Database: "postgres",
		Drop:     true,
	}
)

func TestNewRepository(t *testing.T) {
	repo, err := NewRepository(
		repoOpts,
	)
	if err != nil {
		t.Errorf("failed to connect to postgres: %s", err)
	}

	defer repo.Close()
}

func TestCreate(t *testing.T) {
	repo, err := NewRepository(
		repoOpts,
	)
	if err != nil {
		t.Errorf("failed to connect to postgres: %s", err)
	}

	defer repo.Close()

	user, err := repo.Create(ddd.UserCreate{
		Email:     "jackson@juandefu.ca",
		FirstName: "Jackson",
		LastName:  "Sabey",
		Password:  "pass",
	})
	if err != nil {
		t.Errorf("failed to create user: %s", err)
	}

	if user.Email != "jackson@juandefu.ca" {
		t.Errorf("unknown user found: %s", user.Email)
	}
}

func TestCreate_AlreadyExists(t *testing.T) {
	repo, err := NewRepository(
		repoOpts,
	)
	if err != nil {
		t.Errorf("failed to connect to postgres: %s", err)
	}

	defer repo.Close()

	user, err := repo.Create(ddd.UserCreate{
		Email:     "jackson@sabey.co",
		FirstName: "Jackson",
		LastName:  "Sabey",
		Password:  "pass",
	})
	if err != nil {
		t.Errorf("failed to create user: %s", err)
	}

	if user.Email != "jackson@sabey.co" {
		t.Errorf("unknown user found: %s", user.Email)
	}

	_, err = repo.Create(ddd.UserCreate{
		Email:     "jackson@sabey.co",
		FirstName: "Jackson",
		LastName:  "Sabey",
		Password:  "pass",
	})
	if err == nil {
		t.Errorf("created a duplicate user?")
	}
}

func TestLogin_NotFound(t *testing.T) {
	repo, err := NewRepository(
		repoOpts,
	)
	if err != nil {
		t.Errorf("failed to connect to postgres: %s", err)
	}

	defer repo.Close()

	_, err = repo.Login(
		ddd.UserLogin{
			Email:    "jackson+notfound@juandefu.ca",
			Password: "pass",
		},
	)
	if err == nil {
		t.Errorf("user should not exist?")
	}
}

func TestLogin(t *testing.T) {
	repo, err := NewRepository(
		repoOpts,
	)
	if err != nil {
		t.Errorf("failed to connect to postgres: %s", err)
	}

	defer repo.Close()

	_, err = repo.Create(ddd.UserCreate{
		Email:     "jackson@juandefu.ca",
		FirstName: "Jackson",
		LastName:  "Sabey",
		Password:  "pass",
	})
	if err != nil {
		t.Errorf("failed to create user: %s", err)
	}

	user, err := repo.Login(
		ddd.UserLogin{
			Email:    "jackson@juandefu.ca",
			Password: "pass",
		},
	)
	if err != nil {
		t.Errorf("failed to login: %s", err)
	}

	if user.Email != "jackson@juandefu.ca" {
		t.Errorf("unknown user found: %s", user.Email)
	}
}

func TestList(t *testing.T) {
	repo, err := NewRepository(
		repoOpts,
	)
	if err != nil {
		t.Errorf("failed to connect to postgres: %s", err)
	}

	defer repo.Close()

	_, err = repo.Create(ddd.UserCreate{
		Email:     "jackson@juandefu.ca",
		FirstName: "Jackson",
		LastName:  "Sabey",
		Password:  "pass",
	})
	if err != nil {
		t.Errorf("failed to create user: %s", err)
	}

	_, err = repo.Create(ddd.UserCreate{
		Email:     "jackson@sabey.co",
		FirstName: "Jackson",
		LastName:  "Sabey",
		Password:  "pass",
	})
	if err != nil {
		t.Errorf("failed to create user: %s", err)
	}

	users, err := repo.List()
	if err != nil {
		t.Errorf("failed to login: %s", err)
	}

	if len(users) != 2 {
		t.Errorf("invalid amount of users: %d", len(users))
	}

	if users[0].Email != "jackson@juandefu.ca" {
		t.Errorf("1. unknown email / user sort order: %s", users[1].Email)
	}

	if users[1].Email != "jackson@sabey.co" {
		t.Errorf("0. unknown email / user sort order: %s", users[0].Email)
	}
}

func TestUpdate(t *testing.T) {
	repo, err := NewRepository(
		repoOpts,
	)
	if err != nil {
		t.Errorf("failed to connect to postgres: %s", err)
	}

	defer repo.Close()

	_, err = repo.Create(ddd.UserCreate{
		Email:     "jackson@juandefu.ca",
		FirstName: "Jackson",
		LastName:  "Sabey",
		Password:  "pass",
	})
	if err != nil {
		t.Errorf("failed to create user: %s", err)
	}

	user, err := repo.Update(
		ddd.UserUpdate{
			Email:     "jackson@juandefu.ca",
			FirstName: "JACKSON",
			LastName:  "SABEY",
		},
	)
	if err != nil {
		t.Errorf("failed to login: %s", err)
	}

	if user.FirstName != "JACKSON" {
		t.Errorf("firstname not updated: %s", user.FirstName)
	}

	if user.LastName != "SABEY" {
		t.Errorf("lastname not updated: %s", user.LastName)
	}
}
