package main

import (
	helloHandler "awesomeProject/internal/http/hello"
	helloSvc "awesomeProject/internal/service/hello"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

func main() {
	app := gofr.New()
	app.Server.ValidateHeaders = false

	hsvc := helloSvc.New()
	helloHTTP := helloHandler.New(hsvc)
	app.REST("hello", helloHTTP)

	app.Start()

}
