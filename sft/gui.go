package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/layout"
    "fyne.io/fyne/v2/widget"
)

func Text(text string) *widget.Label {
	return widget.NewLabel(text)
}

func InputText(placeholder, text string, callback func(string)) *widget.Entry {
	input := widget.NewEntry()
	input.SetPlaceHolder(placeholder)
	input.SetText(text)
	input.SetMinRowsVisible(len(text))
	input.OnChanged = callback
	return input
}

func InputField(desc, placeholder, text string, v *string) *fyne.Container {
	return container.New(
		layout.NewFormLayout(), 
		Text(desc),
		InputText(placeholder, text, func(s string) {
			fmt.Printf("Updated %s from '%s' to '%s'\n", desc, *v, s)
			*v = s
		}))
}

func update(s *string, title, to string) {
	fmt.Printf("Updated %s from '%s' to '%s'", title, *s, to)
	*s = to
}

func center(c *fyne.Container) *fyne.Container {
	return container.New(layout.NewCenterLayout(), c)
}

func InitGUI() {
	var ip, port, password string
	minSize := fyne.NewSize(300, 300)

	a := app.New()
    w := a.NewWindow("Simple File Transfer")

    w.Resize(minSize)

    sendUI := center(container.NewVBox(
    	InputField("IP", "127.0.0.1", "", &ip),
    	InputField("Port", "1337", "", &port),
    	InputField("Password", "azerty", "", &password),
    	widget.NewLabel("Send UI (not done yet)")))

    receiveUI := center(container.NewVBox(
    	widget.NewLabel("Receive UI (not done yet)")))

    initUI := center(container.NewVBox(
    	widget.NewButton("Send", func() {w.SetContent(sendUI)}),
    	widget.NewButton("Receive", func() {w.SetContent(receiveUI)})))

    w.SetContent(initUI)
    w.ShowAndRun()
}