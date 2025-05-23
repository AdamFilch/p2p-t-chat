package TUI

import (
	components_lobby "main/src/components/lobby"

	"github.com/rivo/tview"
)

func Lobby() *tview.Flex {

	box := tview.NewFlex().
		AddItem(tview.NewBox(), 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(tview.NewBox(), 0, 1, false).
			AddItem(components_lobby.LobbyForm(), 0, 2, true).
			AddItem(tview.NewBox(), 5, 1, false),
			0, 2, true).
		AddItem(tview.NewBox(), 0, 1, false)

	box.SetBorder(true)
	return box
}
