package main

import (
	"github.com/iftdt/server/common"
	"github.com/iftdt/server/routers"
)

func main() {
	app := routers.SetupRouters()
	app.Run(common.ENV.ServerPort)
}
