package components_lobby

import (
	"github.com/rivo/tview"
)

func LobbyForm() *tview.Form {
	
	userNameInput := tview.NewInputField().SetLabel(`Your Username`)
	roomIdInput := tview.NewInputField().SetLabel(`Room ID`)

	form := tview.NewForm().
		AddFormItem(userNameInput).
		AddFormItem(roomIdInput).
		AddButton("Start Chatting", nil).SetButtonsAlign(tview.AlignCenter)


	return form
}