package main

import (
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/util/application"
)

func main() {

	app, _ := application.Create(application.Options{
		Title:  "Hello G3N",
		Width:  800,
		Height: 600,
	})

	w1 := gui.NewWindow(100,100)
	w1.SetPosition(100,100)
	w1.SetResizable(true)
	app.Gui().Add(w1)

	app.Run()
}
