package service

import (
	"errors"

	"github.com/nathanbahiadev/go-immobr/domain/entity"
)

func Authenticate(
	repository entity.UserRepositoryInterface,
	hasher entity.PasswordHasher,
	email entity.Email,
	password entity.Password,
) (*entity.User, error) {
	user := repository.GetUserByEmail(email)

	if user.ID == 0 {
		return nil, errors.New("authenticate: invalid credentials")
	}

	if err := hasher.Verify(user.Password.Value, password.Value); err != nil {
		return nil, errors.New("authenticate: invalid credentials")
	}
	return &user, nil
}
