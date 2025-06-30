package components_chatroom

import "github.com/rivo/tview"

func UserList() *tview.List {

	// box := tview.NewBox().SetBorder(true).SetTitle("User List")

	box := tview.NewList().
		AddItem("Adam Filchoir", "Hey bro hows...", 'q', nil).
		AddItem("Bagas Irfan", "Did you see th...", 'w', nil).
		AddItem("Farhan Nurdiansyah", "Hey bro hows...", 'e', nil).
		AddItem("Donald trump", "Hey bro hows...", 'r', nil).
		AddItem("Barrack Obama", "Hey bro hows...", 't', nil).
		AddItem("Sabrina Carpenter", "Hey bro hows...", 'y', nil)

	box.SetBorder(true).SetTitle("Alt + Q")

	return box
}
