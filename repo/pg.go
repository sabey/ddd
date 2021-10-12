package repo

import (
	"errors"

	"github.com/go-pg/pg"
	"github.com/sabey/ddd"
	"github.com/sabey/ddd/repo/models"
)

type RepositoryOpts struct {
	Addr     string
	User     string
	Password string
	Database string
	Drop     bool
}

func NewRepository(
	opts RepositoryOpts,
) (*Repository, error) {
	db := pg.Connect(&pg.Options{
		Addr:     opts.Addr,
		User:     opts.User,
		Password: opts.Password,
		Database: opts.Database,
	})

	if opts.Drop {
		_, err := db.Exec("DROP TABLE users;")
		if err != nil {
			return nil, err
		}
	}

	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email VARCHAR(255) UNIQUE,
		firstname VARCHAR(255),
		lastname VARCHAR(255),
		password VARCHAR(255)
	);`)
	if err != nil {
		return nil, err
	}

	return &Repository{
		db: db,
	}, nil
}

type Repository struct {
	db *pg.DB
}

func (r *Repository) Close() error {
	return r.db.Close()
}

func (r *Repository) Create(
	opts ddd.UserCreate,
) (
	*ddd.User,
	error,
) {
	if err := opts.Validate(); err != nil {
		return nil, err
	}

	user := &models.User{
		Email:     opts.Email,
		Firstname: opts.FirstName,
		Lastname:  opts.LastName,
		Password:  opts.Password,
	}

	_, err := r.db.Model(user).Insert()
	if e, ok := err.(pg.Error); ok && e.IntegrityViolation() {
		return nil, errors.New("user already exists")
	}

	if err != nil {
		return nil, err
	}

	return &ddd.User{
		Email:     user.Email,
		FirstName: user.Firstname,
		LastName:  user.Lastname,
		Password:  user.Password,
	}, nil
}

func (r *Repository) Login(
	opts ddd.UserLogin,
) (
	*ddd.User,
	error,
) {
	if err := opts.Validate(); err != nil {
		return nil, err
	}

	user := &models.User{}

	err := r.db.Model(user).Where("email = ?", opts.Email).Select()
	if err != nil {
		return nil, err
	}

	if user.Password != opts.Password {
		return nil, errors.New("password is invalid")
	}

	return &ddd.User{
		Email:     user.Email,
		FirstName: user.Firstname,
		LastName:  user.Lastname,
		Password:  user.Password,
	}, nil
}

func (r *Repository) List() (
	[]*ddd.User,
	error,
) {
	users := &[]*models.User{}

	err := r.db.Model(&models.User{}).Order("email ASC").Select(users)
	if err != nil {
		return nil, err
	}

	return newUserList(*users), nil
}

func newUserList(users []*models.User) []*ddd.User {
	ur := []*ddd.User{}

	for _, user := range users {
		ur = append(ur, &ddd.User{
			Email:     user.Email,
			FirstName: user.Firstname,
			LastName:  user.Lastname,
			Password:  user.Password,
		})
	}

	return ur
}

func (r *Repository) Update(
	opts ddd.UserUpdate,
) (
	*ddd.User,
	error,
) {
	if err := opts.Validate(); err != nil {
		return nil, err
	}

	user := &models.User{}

	_, err := r.db.Model(user).
		Set("firstname = ?", opts.FirstName).
		Set("lastname = ?", opts.LastName).
		Where("email = ?", opts.Email).Returning("*").
		Update()
	if err != nil {
		return nil, err
	}

	return &ddd.User{
		Email:     user.Email,
		FirstName: user.Firstname,
		LastName:  user.Lastname,
		Password:  user.Password,
	}, nil
}
