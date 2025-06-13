package TUI

import (
	components_chatroom "main/src/components/chatroom"

	"github.com/rivo/tview"
)

func Chatroom() *tview.Flex {
	box := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(components_chatroom.UserList(), 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(components_chatroom.ChatHistory(), 0, 1, false).
			AddItem(components_chatroom.ChatArea(), 0, 1, false),
			0, 5, false)

	return box
}