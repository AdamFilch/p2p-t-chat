package components_lobby

import "github.com/rivo/tview"

func LobbyForm() *tview.Form {
	
	userNameInput := tview.NewInputField().SetLabel(`Your Username`)
	roomIdInput := tview.NewInputField().SetLabel(`Room ID (Send to friend to start chatting)`)

	form := tview.NewForm().
		AddFormItem(userNameInput).
		AddFormItem(roomIdInput).
		// AddTextArea("Address", "", 40, 0, 0, nil).
		AddButton("Start Chatting", nil)


	return form
}