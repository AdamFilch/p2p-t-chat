package components_chatroom

import "github.com/rivo/tview"


func UserList() *tview.Box {
	
	box := tview.NewBox().SetBorder(true).SetTitle("User List")

	return box
}