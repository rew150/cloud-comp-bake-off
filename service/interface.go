package service

import (
	"fmt"
	"math/rand"
	"time"
)

type DbService interface {
	SaveNewAccount(username, password string) error
	FindHashedPassword(username string) (password string, found bool, err error)
}

type StubDb struct{}

type StubDbData struct {
	Password string
}

var StubData StubDbData

func (s StubDb) SaveNewAccount(username, password string) error {
	fmt.Println("Stub Save: username =", username, "password =", password)
	return nil
}

func (s StubDb) FindHashedPassword(username string) (password string, found bool, err error) {
	fmt.Println("Stub Find: username =", username)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	n := r1.Float64()
	if n <= 0.2 {
		fmt.Println("Stub Find: not found")
		return "", false, nil
	}

	return StubData.Password, true, nil
}
