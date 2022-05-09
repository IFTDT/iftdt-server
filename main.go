package main

import "github.com/iftdt/server/routers"

func main() {
	app := routers.SetupRouters()
	app.Run()
}
