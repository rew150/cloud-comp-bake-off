package config

import (
	"fmt"
	"lilandfriends/bakeoff/service"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Env struct {
	IsDynamoDB bool
}

var Database service.DbService
var Config Env

func init() {
	d := os.Getenv("IS_DYNAMO_DB")
	if d != "" {
		Config.IsDynamoDB = true
	} else {
		Config.IsDynamoDB = false
	}

	fmt.Println("[CONFIG]")
	fmt.Println(Config)

	// TODO: Add db implementation
	Database = service.StubDb{}
}
