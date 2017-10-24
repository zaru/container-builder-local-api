package main

import (
	"flag"
	"github.com/labstack/echo"
	"github.com/zaru/container-builder-local-api/handler"
)

func main() {

	port := flag.String("port", "1323", "port number")
	flag.Parse()

	e := echo.New()
	e.POST("/v1/projects/:project_id/builds", handler.BuildCreate)
	e.Logger.Fatal(e.Start(":" + *port))
}
