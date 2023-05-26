package controllers

import (
	"github.com/gin-gonic/gin"
)

func TestController(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Olá, mundo!"})
}
