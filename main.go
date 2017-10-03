package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.POST("/users", save)
	e.Logger.Fatal(e.Start(":1323"))
}

type User struct {
	Steps string `json:"steps[]"`
}
type BuildParams struct {
	Name       string   `json:"name"`
	Env        []string `json:"env"`
	Args       []string `json:"args"`
	WaitFor    []string `json:"waitFor"`
	Entrypoint string   `json:"entrypoint"`
}

func save(c echo.Context) error {
	post := new(BuildParams)
	if err := c.Bind(post); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, post)
}