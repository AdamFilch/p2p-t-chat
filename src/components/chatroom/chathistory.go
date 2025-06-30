package components_chatroom

import "github.com/rivo/tview"

func ChatHistory() *tview.Box {
	
	box := tview.NewBox().SetBorder(true).SetTitle("Alt + W")

	return box
}