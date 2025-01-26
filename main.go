package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("TODO App")

	//	w.SetContent(widget.NewLabel("TODOs will go here"))

	//	t := todo.NewTodo("Show this on the window")

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

	w.SetContent(
		container.NewBorder(
			nil, // TOP of the container

			container.NewBorder(
				nil, // TOP
				nil, // BOTTOM
				nil, // Left
				// RIGHT ↓
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
