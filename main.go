package main

import (
	"github.com/labstack/echo"
	"github.com/zaru/container-builder-local-api/handler"
)

func main() {
	e := echo.New()
	e.POST("/v1/projects/:project_id/builds", handler.BuildCreate)
	e.Logger.Fatal(e.Start(":1323"))
}
