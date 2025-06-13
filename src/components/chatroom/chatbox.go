package components_chatroom

import "github.com/rivo/tview"

func ChatArea() *tview.Box {
	
	box := tview.NewBox().SetBorder(true).SetTitle("Chat Area")

	return box
}