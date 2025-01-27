package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/JacobGraver/ToDoList-App.git/todo"
)

func main() {
	a := app.New()
	w := a.NewWindow("TODO App")
	w.Resize(fyne.NewSize(300, 400))

	data := []todo.Todo{
		todo.NewTodo("Some stuff"),
		todo.NewTodo("Some more stuff"),
		todo.NewTodo("Some other things"),
	}

	newtodoDescTxt := widget.NewEntry()
	newtodoDescTxt.PlaceHolder = "New Todo Description..."
	addBtn := widget.NewButton("Add", func() { fmt.Println("Add was clicked!") })
	addBtn.Disable()

	newtodoDescTxt.OnChanged = func(s string) {
		addBtn.Disable()

		if len(s) >= 3 {
			addBtn.Enable()
		}
	}

	// TODO fix window positioning
	w.SetContent(
		container.NewBorder(
			nil, // TOP of the container
			container.NewBorder(
				nil, // TOP
				nil, // BOTTOM
				widget.NewList(
					func() int {
						return len(data)
					},
					func() fyne.CanvasObject {
						return widget.NewLabel("template")
					},
					func(i widget.ListItemID, o fyne.CanvasObject) {
						o.(*widget.Label).SetText(data[i].Description)
					},
				),
				addBtn,
				// take the rest of the space
				newtodoDescTxt,
			),
			nil, // Left
			nil, // Right
		),
	)

	w.ShowAndRun()
}
