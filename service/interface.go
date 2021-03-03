package service

import "fmt"

type DbService interface {
	SaveNewAccount(username, password string) error
}

type StubDb struct{}

func (s StubDb) SaveNewAccount(username, password string) error {
	fmt.Println("Stub: username =", username, "password =", password)
	return nil
}
