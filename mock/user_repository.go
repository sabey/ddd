package mock

import (
	"errors"
	"sort"

	"github.com/sabey/ddd"
)

func NewUserRepository() *UserRepository {
	return &UserRepository{
		Accounts: make(map[string]ddd.User),
	}
}

type UserRepository struct {
	// [Email]User
	Accounts map[string]ddd.User
}

func (ur UserRepository) Create(
	opts ddd.UserCreate,
) (
	*ddd.User,
	error,
) {
	if err := opts.Validate(); err != nil {
		return nil, err
	}

	if _, ok := ur.Accounts[opts.Email]; ok {
		return nil, errors.New("user account already exists")
	}

	// account created
	user := ddd.User{
		Email:     opts.Email,
		FirstName: opts.FirstName,
		LastName:  opts.LastName,
		Password:  opts.Password,
	}

	ur.Accounts[opts.Email] = user

	return &user, nil
}

func (ur UserRepository) Login(
	opts ddd.UserLogin,
) (
	*ddd.User,
	error,
) {
	if err := opts.Validate(); err != nil {
		return nil, err
	}

	user, ok := ur.Accounts[opts.Email]

	if !ok {
		return nil, errors.New("user account doesn't exist")
	}

	if user.Password != opts.Password {
		return nil, errors.New("password is invalid")
	}

	return &user, nil
}

func (ur UserRepository) List() (
	[]*ddd.User,
	error,
) {
	keys := make([]string, 0, len(ur.Accounts))
	for k := range ur.Accounts {
		keys = append(keys, k)
	}
	// ORDER BY email ASC
	sort.Sort(sort.Reverse(sort.StringSlice(keys)))

	users := []*ddd.User{}
	for _, key := range keys {
		user := ur.Accounts[key]
		users = append(users, &user)
	}

	return users, nil
}

func (ur UserRepository) Update(
	opts ddd.UserUpdate,
) (
	*ddd.User,
	error,
) {
	if err := opts.Validate(); err != nil {
		return nil, err
	}

	user, ok := ur.Accounts[opts.Email]

	if !ok {
		return nil, errors.New("user account doesn't exist")
	}

	user.FirstName = opts.FirstName
	user.LastName = opts.LastName

	// update account
	ur.Accounts[opts.Email] = user

	return &user, nil
}
