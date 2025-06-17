package TUI

import (
	"fmt"
	components_chatroom "main/src/components/chatroom"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Chatroom() *tview.Flex {


	input := components_chatroom.ChatArea()
	input.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		fmt.Println("Key pressed in input area")
		return event
	})

	box := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(components_chatroom.UserList(), 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(components_chatroom.ChatHistory(), 0, 2, false).
			AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
				AddItem(components_chatroom.ChatArea(), 0, 6, true).
				AddItem(components_chatroom.SubmitText(), 0, 1, false), 0, 1, true),
			0, 5, true)

	return box
}