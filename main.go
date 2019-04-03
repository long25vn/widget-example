package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"

	"github.com/kataras/iris"
)

type layout struct {
	Type       int32
	Code       []template.HTML
	CodeString template.HTML
}

func main() {
	app := iris.New()

	jsonFile, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var layoutArr []layout
	json.Unmarshal([]byte([]byte(byteValue)), &layoutArr)

	app.RegisterView(iris.HTML("./templates", ".html"))
	app.Get("/", func(ctx iris.Context) {

		ctx.ViewData("Layout", layoutArr)
		ctx.View("index.html")

	})

	app.Run(iris.Addr(":8080"))
}
