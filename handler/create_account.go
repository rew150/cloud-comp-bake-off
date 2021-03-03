package handler

import (
	"lilandfriends/bakeoff/config"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateAccount(c *gin.Context) {
	var body AccountDto
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	bHashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = config.Database.SaveNewAccount(body.Username, string(bHashedPassword))
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.String(201, "OK")
}
