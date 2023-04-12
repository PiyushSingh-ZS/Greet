package main

import (
	"awesomeProject/internal/http/bye"
	helloHandler "awesomeProject/internal/http/hello"
	"awesomeProject/internal/http/name"
	helloSvc "awesomeProject/internal/service/hello"
	nameSvc "awesomeProject/internal/service/name"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

func main() {
	app := gofr.New()
	app.Server.ValidateHeaders = false
	byeHTTP := bye.New()
	app.REST("bye", byeHTTP)

	svc := nameSvc.New()
	nameHTTP := name.New(svc)
	app.REST("", nameHTTP)

	hsvc := helloSvc.New()
	helloHTTP := helloHandler.New(hsvc)
	app.REST("hello", helloHTTP)

	app.Start()

}
