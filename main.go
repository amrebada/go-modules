package main

import (
	// "flag"
	"encoding/json"
	"fmt"

	"github.com/amrebada/go-modules/core"
	// "github.com/amrebada/go-modules/modules"
	"github.com/amrebada/go-modules/modules/user"
)

func main() {
	RegisterDTO := user.RegisterDto{}
	StringDTO := ""
	ArrayRegisterDTO := []user.RegisterDto{}
	schemas := core.Schemas{}
	schemas = core.ConvertToSchema(RegisterDTO, schemas)
	schemas = core.ConvertToSchema(StringDTO, schemas)
	schemas = core.ConvertToSchema(ArrayRegisterDTO, schemas)
	data, err := json.MarshalIndent(schemas, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

	// isMigrate := false
	// isSwagger := false
	// env := "dev"

	// flag.BoolVar(&isMigrate, "m", false, "auto migrate database")
	// flag.BoolVar(&isSwagger, "sw", false, "generate swagger")
	// flag.StringVar(&env, "env", "dev", "identify which environment to load from {.env} file [.env.prod, .env.dev, .env.test, .env]")
	// flag.Parse()
	// core.NewConfig(core.Stage(env)).
	// 	SetMigrate(isMigrate).
	// 	SetSwagger(isSwagger)

	// app := core.NewServer()
	// app.MainModule = modules.NewAppModule()
	// app.RegisterMainModule()
	// err := app.Start()
	// if err != nil {
	// 	fmt.Println(err)
	// }

}
