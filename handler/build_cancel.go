package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

func BuildCancel(c echo.Context) error {

	build_id := c.Param("build_id")

	res := &CancelResponse{
		ID:     build_id,
		Status: "CANCELLED",
	}
	return c.JSON(http.StatusCreated, res)
}

type CancelResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}
