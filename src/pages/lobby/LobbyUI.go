package TUI

import (
	components_lobby "main/src/components/lobby"

	"github.com/rivo/tview"
)

var sample string = "In Go, backticks denote raw string literals. Within these literals, all characters, including newlines and other special characters, are interpreted literally, except for the backtick character itself. Therefore, a raw string literal cannot contain a backtick."

func Lobby() *tview.Flex {

	box := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(components_lobby.LobbyLogo(), 0, 2, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(tview.NewBox(), 0, 3, false).
			AddItem(components_lobby.LobbyForm(), 0, 3, true).
			AddItem(tview.NewBox(), 0, 3, false),
			0, 1, true).		
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(tview.NewBox(), 0, 2, false).
			AddItem(tview.NewTextView().SetText(sample), 0, 3, true).
			AddItem(tview.NewBox(), 0, 2, false),
			0, 1, true)
		box.SetBorder(true)
	return box
}
