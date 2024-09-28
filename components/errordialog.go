package components

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type ErrorDialog struct {
	Dialog     *dialog.CustomDialog
	DialogText *widget.Label
}

func NewErrorDialog(app *fyne.App, window *fyne.Window) *ErrorDialog {
	dialogIntro := widget.NewLabel("An unexpected error ocurred, the application will be closing. \n\nError:")
	dialogIntro.Wrapping = fyne.TextWrapWord
	dialogText := widget.NewLabel("Error message")
	dialogText.Wrapping = fyne.TextWrapWord
	dialogTextContainer := container.NewScroll(dialogText)
	dialogTextContainer.SetMinSize(fyne.NewSize(200, 200))
	dialogButton := widget.NewButton("Close", func() {
		(*app).Quit()
	})
	dialogContent := container.NewVBox(
		dialogIntro,
		dialogTextContainer,
		dialogButton,
	)
	dialog := dialog.NewCustomWithoutButtons("An error ocurred:", dialogContent, *window)
	dialog.Resize(fyne.NewSize(300, 400))
	return &ErrorDialog{
		Dialog:     dialog,
		DialogText: dialogText,
	}
}

func (d *ErrorDialog) Show(err error) {
	if err == nil {
		log.Print("nil error submitted to ErrorDialog.Show, doing nothing")
		return
	}
	d.DialogText.SetText(err.Error())
	d.Dialog.Show()
}
