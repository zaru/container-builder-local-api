package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"os"
	"encoding/json"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.POST("/users", save)
	e.Logger.Fatal(e.Start(":1323"))
}

//MEMO: これは使わない…かな…中身をパースすることはなさそう
type BuildParams struct {
	Steps []struct {
		Name       string   `json:"name"`
		Env        []string `json:"env"`
		Args       []string `json:"args"`
		WaitFor    []string `json:"waitFor"`
		Entrypoint string   `json:"entrypoint"`
	} `json:"steps"`
	LogsBucket string `json:"logsBucket"`
}

type Response struct {
	Metadata struct {
		Build struct {
			ID string `json:"id"`
		} `json:"build"`
	} `json:"metadata"`
}

func save(c echo.Context) error {
	body, _ := ioutil.ReadAll(c.Request().Body)
	content := []byte(string(body))
	//TODO: cloudbuild.yamlファイルをワーキングディレクトリに出力する
	ioutil.WriteFile("/tmp/go-file", content, os.ModePerm)
	fmt.Print(string(body))

	//TODO: 動的にユニークなビルドIDを適当に返す
	build_id := "build-1234"
	jsonResponse := `{"metadata": {"build": {"id": "` + build_id + `"}}}`
	res := &Response{}
	err := json.Unmarshal([]byte(jsonResponse), res)
	if(err!=nil) {
		return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")
	}

	//TODO: container-builder-localでビルドする

	return c.JSON(http.StatusCreated, res)
}