package TUI

import "github.com/rivo/tview"


func Lobby() *tview.Box {
	box := tview.NewBox().SetBorder(true).SetTitle(" Lobby ")
	return box
}