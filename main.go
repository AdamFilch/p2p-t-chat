package main

import (
	TUI "main/src/pages/chatroom"

	"github.com/rivo/tview"
)

func main() {


	app := tview.NewApplication()
	box := TUI.Chatroom()
	if err := app.SetRoot(box, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

	
	
	TUI.Chatroom()
}

