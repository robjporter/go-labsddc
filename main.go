package main

import (
	"./app"
)

var (
	App *app.Application
)

func main() {
	App = app.New()
	App.Run()
}
