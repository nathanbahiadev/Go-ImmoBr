package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nathanbahiadev/go-immobr/domain/dto"
	"github.com/nathanbahiadev/go-immobr/domain/entity"
	"github.com/nathanbahiadev/go-immobr/domain/service"
	"github.com/nathanbahiadev/go-immobr/infrastructure/repository"
	"github.com/nathanbahiadev/go-immobr/infrastructure/utils"
)

var userRepository = repository.UserRepository{}
var passwordHasher = utils.PasswordHasher{}

func CreateUserController(c *gin.Context) {
	var userDto dto.User
	if err := c.BindJSON(&userDto); err != nil {
		c.JSON(422, gin.H{"message": err.Error()})
		return
	}

	user, err := entity.User{
		Name:     userDto.Name,
		Email:    entity.Email{Value: userDto.Email},
		Password: entity.Password{Value: userDto.Password},
	}.Create(
		userRepository,
		passwordHasher,
	)

	if err != nil {
		c.JSON(422, gin.H{"message": err.Error()})
		return
	}

	userResponse := dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email.Value,
	}

	c.JSON(201, userResponse)
}

func LoginUserController(c *gin.Context) {
	var userCredentials dto.UserLogin
	if err := c.ShouldBind(&userCredentials); err != nil {
		c.JSON(401, gin.H{
			"message": "email and password are required",
		})
		return
	}

	user, err := service.Authenticate(
		userRepository,
		passwordHasher,
		entity.Email{Value: userCredentials.Email},
		entity.Password{Value: userCredentials.Password},
	)

	if err != nil {
		c.JSON(401, gin.H{
			"message": err.Error(),
		})
		return
	}

	token, err := utils.CreateToken(user.Email.Value)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Logged successfully",
		"token":   token,
	})
}
