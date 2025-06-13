package components_chatroom

import "github.com/rivo/tview"

func ChatHistory() *tview.Box {
	
	box := tview.NewBox().SetBorder(true).SetTitle("Chat with this User")

	return box
}