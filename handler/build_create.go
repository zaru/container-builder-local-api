package handler

import (
	"bytes"
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

func BuildCreate(c echo.Context) error {

	build_id := build_id()

	save_err := save_build_id(build_id)
	if save_err != nil {
		return save_err
	}

	if err := os.MkdirAll("/tmp/gcb-local/"+build_id, 0755); err != nil {
		fmt.Println(err)
	}

	body, _ := ioutil.ReadAll(c.Request().Body)
	content := []byte(string(body))
	file_err := ioutil.WriteFile("/tmp/gcb-local/"+build_id+"/cloudbuild.json", content, os.ModePerm)
	if file_err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "System Error")
	}
	fmt.Print(string(body))

	jsonResponse := `{"metadata": {"build": {"id": "` + build_id + `"}}}`
	res := &Response{}
	err := json.Unmarshal([]byte(jsonResponse), res)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "System Error")
	}

	go gcb_build(build_id)

	return c.JSON(http.StatusCreated, res)
}

type Response struct {
	Metadata struct {
		Build struct {
			ID string `json:"id"`
		} `json:"build"`
	} `json:"metadata"`
}

func build_id() string {
	return uuid.NewV4().String()
}

func save_build_id(build_id string) error {
	db, _ := leveldb.OpenFile("./db", nil)
	defer db.Close()

	return db.Put([]byte(build_id), []byte("dummy value"), nil)
}

func gcb_build(build_id string) {

	cmd := exec.Command("container-builder-local",
		"--config=/tmp/gcb-local/"+build_id+"/cloudbuild.json",
		"--dryrun=false", "/tmp/gcb-local/"+build_id)

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}

	//TODO: save pid
	fmt.Println("pid: ", cmd.Process.Pid)

	cmd.Wait()

	log := stdout.String()
	output_log(log, build_id)
	fmt.Println(log)

}

func output_log(log, build_id string) error {
	content := []byte(log)
	return ioutil.WriteFile("/tmp/gcb-local/"+build_id+"/result.txt", content, os.ModePerm)
}
