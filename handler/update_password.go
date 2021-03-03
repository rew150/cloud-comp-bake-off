package handler

import (
	"lilandfriends/bakeoff/config"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func UpdatePassword(c *gin.Context) {
	var body AccountDto
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	// assume logged in
	bHashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = config.Database.UpdatePassword(body.Username, string(bHashedPassword))
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.String(200, "OK")
}
