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
	//TODO: escape colon
	//ref https://forum.labstack.com/t/i-want-to-use-a-colon-in-the-url/394/1
	e.POST("/v1/projects/:project_id/builds/:build_id:cancel", handler.BuildCancel)
	e.Logger.Fatal(e.Start(":" + *port))
}
