package main

import (
	"fmt"

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
		container.NewBorder(
			nil, // TOP of the container

			// this will be a the BOTTOM of the container
			widget.NewButton("Add", func() { fmt.Println("Add was clicked!") }),

			nil, // Right
			nil, // Left

			// the rest will take all the rest of the space
			container.NewCenter(
				widget.NewLabel(t.String()),
			),
		),
	)

	w.ShowAndRun()
}
