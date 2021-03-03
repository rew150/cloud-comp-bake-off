package handler

import (
	"lilandfriends/bakeoff/config"
	"lilandfriends/bakeoff/service"
	"log"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var body AccountDto
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	//TODO: Change Stub out
	{
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)

		n := r1.Float64()
		var p string
		if n <= 0.4 {
			p = "1234abcd"
		} else {
			p = body.Password
		}

		hp, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalln("Stub: unexpected error")
		}

		service.StubData.Password = string(hp)
	}

	hashedPassword, found, err := config.Database.FindHashedPassword(body.Username)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !found {
		c.String(404, "username not found")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(body.Password))
	if err != nil {
		c.String(401, "password not correct")
		return
	}

	c.String(200, "OK")
}
