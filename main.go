package main

import (
	TUI "main/src/pages/lobby"

	"github.com/rivo/tview"
)

func main() {


	app := tview.NewApplication()
	box := TUI.Lobby()
	if err := app.SetRoot(box, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

	
	
	TUI.Lobby()
}

