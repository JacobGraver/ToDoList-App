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
	//	w.Resize(fyne.NewSize(300, 400))

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

	w.SetContent(
		container.NewBorder(
			nil, // TOP of the container
			container.NewBorder(
				nil, // TOP
				nil, // BOTTOM
				nil,
				addBtn,
				newtodoDescTxt,
			),
			nil,
			nil,
			widget.NewList(
				// func that returns the number of items in the list
				func() int {
					return len(data)
				},
				// func that returns the component structure of the List Item
				func() fyne.CanvasObject {
					return widget.NewLabel("template")
				},
				// func that is called for each item in the list and allows
				// you to show the content on the previously defined ui structure
				func(i widget.ListItemID, o fyne.CanvasObject) {
					o.(*widget.Label).SetText(data[i].Description)
				}),
		),
	)

	w.ShowAndRun()
}
