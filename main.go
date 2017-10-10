package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"github.com/syndtr/goleveldb/leveldb"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	e := echo.New()
	e.POST("/v1/projects/:project_id/builds", build)
	e.Logger.Fatal(e.Start(":1323"))
}

type Response struct {
	Metadata struct {
		Build struct {
			ID string `json:"id"`
		} `json:"build"`
	} `json:"metadata"`
}

func build(c echo.Context) error {

	build_id := build_id()

	save_err := save_build_id(build_id)
	if save_err != nil {
		return save_err
	}

	body, _ := ioutil.ReadAll(c.Request().Body)
	content := []byte(string(body))
	ioutil.WriteFile("/tmp/gcb-local/"+build_id+"/cloudbuild.json", content, os.ModePerm)
	fmt.Print(string(body))

	jsonResponse := `{"metadata": {"build": {"id": "` + build_id + `"}}}`
	res := &Response{}
	err := json.Unmarshal([]byte(jsonResponse), res)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")
	}

	err = exec.Command("container-builder-local",
		"--config=/tmp/gcb-local/"+build_id+"/cloudbuild.json",
		"--dryrun=false", "/tmp/gcb-local/"+build_id).Start()
	if err != nil {
		fmt.Println(err)
	}

	return c.JSON(http.StatusCreated, res)
}

func build_id() string {
	return uuid.NewV4().String()
}

func save_build_id(build_id string) error {
	db, _ := leveldb.OpenFile("./db", nil)
	defer db.Close()

	return db.Put([]byte(build_id), []byte("dummy value"), nil)
}
