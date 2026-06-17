package main

import (
	"dinz-rentbike/internal/bootstrap"
)

func main() {
	app := bootstrap.Init()
	app.Run()
}
