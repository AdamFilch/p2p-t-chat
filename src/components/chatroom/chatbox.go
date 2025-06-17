package components_chatroom

import (
	"github.com/rivo/tview"
)

func ChatArea() *tview.TextArea {

	box :=  tview.NewTextArea().SetPlaceholder("Enter your text here...")

	box.SetBorder(true).SetTitle("Chat Area")
	return box
}

func SubmitText() *tview.Button {
	box := tview.NewButton("Submit").SetSelectedFunc(func() {
	})
	box.SetBorder(true)
	return box
}

// func SubmitText(app ) tview.Primitive {
// 	text := tview.NewTextView().
// 		SetText("Click me\nto send your message").
// 		SetTextAlign(tview.AlignCenter).
// 		SetDynamicColors(true).
// 		SetRegions(true).
// 		SetWordWrap(true).
// 		SetBorder(true)

// 	text.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
// 		if event.Key() == tcell.KeyEnter {
// 			// Perform action here
// 			app.Stop()
// 		}
// 		return event
// 	})

// 	return text
// }