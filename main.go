package main

import (
	"fmt"

	"github.com/amrebada/go-template/core"
	"github.com/amrebada/go-template/modules"
)

var config *core.Config = core.NewConfig()

func main() {

	app := core.NewServer()
	app.MainModule = modules.NewAppModule()
	app.RegisterMainModule()
	err := app.Start()
	if err != nil {
		fmt.Println(err)
	}

}
