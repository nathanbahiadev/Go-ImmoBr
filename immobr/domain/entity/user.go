package entity

import (
	"errors"
)

type UserRepositoryInterface interface {
	Create(user User) User
	GetUserByEmail(email Email) User
}

type User struct {
	ID       uint
	Name     string
	Email    Email
	Password Password
}

func (u *User) Validate(repository UserRepositoryInterface) error {
	if err := u.Email.Validate(); err != nil {
		return err
	}

	if err := u.Password.Validate(); err != nil {
		return err
	}

	userWithSameEmail := repository.GetUserByEmail(u.Email)
	if userWithSameEmail.ID != 0 && userWithSameEmail.ID != u.ID {
		return errors.New("email already taken")
	}

	return nil
}

func (u User) Create(repository UserRepositoryInterface, hasher PasswordHasher) (User, error) {
	if err := u.Validate(repository); err != nil {
		return User{}, err
	}

	if err := u.Password.Hash(hasher); err != nil {
		return User{}, nil
	}

	userInDb := repository.Create(u)
	return userInDb, nil
}
