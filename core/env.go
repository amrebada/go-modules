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
		err := dotEnv.Load(".env.dev")
		if err != nil {
			dotEnv.Load()
		}
	} else if env == Production {
		err := dotEnv.Load(".env.prod")
		if err != nil {
			dotEnv.Load()
		}
	} else if env == Test {
		err := dotEnv.Load("../.env.test")
		if err != nil {
			dotEnv.Load()
		}
	} else {
		dotEnv.Load(".env")
	}
}
