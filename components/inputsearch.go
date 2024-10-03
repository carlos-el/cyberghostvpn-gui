package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type InputSearch struct {
	Container *fyne.Container
	Input     *widget.Entry
}

func NewInputSearch() *InputSearch {
	inputSearch := widget.NewEntry()
	inputSearch.SetPlaceHolder("Search...")
	dummySearchIcon := widget.NewButtonWithIcon("", theme.SearchIcon(), nil)
	dummySearchIcon.Disable()

	container := container.New(
		layout.NewFormLayout(),
		dummySearchIcon,
		inputSearch,
	)

	return &InputSearch{
		Container: container,
		Input:     inputSearch,
	}
}

func (is *InputSearch) SetOnChanged(onChanged func(string)) {
	is.Input.OnChanged = onChanged
}

func (is *InputSearch) GetInputText() string {
	return is.Input.Text
}
