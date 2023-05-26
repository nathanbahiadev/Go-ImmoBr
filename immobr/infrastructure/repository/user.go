package repository

import (
	"github.com/nathanbahiadev/go-immobr/domain/entity"
	"github.com/nathanbahiadev/go-immobr/infrastructure/database"
	"github.com/nathanbahiadev/go-immobr/infrastructure/models"
)

type UserRepository struct{}

func (r UserRepository) Create(user entity.User) entity.User {
	userModel := models.User{
		Name:     user.Name,
		Email:    user.Email.Value,
		Password: user.Password.Value,
	}

	database.DB.Create(&userModel)
	user.ID = userModel.ID

	return user
}

func (r UserRepository) GetUserByEmail(email entity.Email) entity.User {
	var userModel models.User
	database.DB.First(&userModel, "email = ?", email.Value)

	return entity.User{
		ID:       userModel.ID,
		Name:     userModel.Name,
		Email:    entity.Email{Value: userModel.Email},
		Password: entity.Password{Value: userModel.Password},
	}
}
