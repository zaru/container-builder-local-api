package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
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
	body, _ := ioutil.ReadAll(c.Request().Body)
	content := []byte(string(body))
	//TODO: cloudbuild.yamlファイルをワーキングディレクトリに出力する
	ioutil.WriteFile("/tmp/cloudbuild.json", content, os.ModePerm)
	fmt.Print(string(body))

	//TODO: 動的にユニークなビルドIDを適当に返す
	build_id := "build-1234"
	jsonResponse := `{"metadata": {"build": {"id": "` + build_id + `"}}}`
	res := &Response{}
	err := json.Unmarshal([]byte(jsonResponse), res)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")
	}

	//TODO: ビルドキャンセルをどうするか検討する
	//TODO: ワーキングディレクトリに対応する
	out, err := exec.Command("container-builder-local", "--config=/tmp/cloudbuild.json", "--dryrun=false", "/tmp").Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(out))

	return c.JSON(http.StatusCreated, res)
}
