package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/JacobGraver/ToDoList-App.git/todo"
)

func main() {
	a := app.New()
	w := a.NewWindow("TODO App")

	//	w.SetContent(widget.NewLabel("TODOs will go here"))

	t := todo.NewTodo("Show this on the window")

	w.SetContent(
		container.NewCenter(
			widget.NewLabel(t.String()),
		),
	)

	w.ShowAndRun()
}
