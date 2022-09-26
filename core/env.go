package core

import (
	dotEnv "github.com/joho/godotenv"
)

type Stage string

const (
	Development Stage = "dev"
	Production  Stage = "prod"
	Test        Stage = "test"
)

func LoadEnv(env Stage) {
	if env == Development {
		dotEnv.Load(".env.dev")
	} else if env == Production {
		dotEnv.Load(".env.prod")
	} else if env == Test {
		dotEnv.Load("../.env.test")
	} else {
		dotEnv.Load(".env")
	}
}
