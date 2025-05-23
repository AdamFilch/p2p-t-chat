package TUI

import "github.com/rivo/tview"


func Lobby() {
	box := tview.NewBox().SetBorder(true).SetTitle("| P2P - T - Chat |")
	if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
		panic(err)
	}
}