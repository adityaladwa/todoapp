package main

import (
	"github.com/adityaladwa/todoapp/app"
)

func main() {
	app := &app.App{}
	app.Init()
	defer app.DB.Close()
}
