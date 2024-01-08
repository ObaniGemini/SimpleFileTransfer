package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"os"
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
	var file fyne.URIReadCloser

	a := app.New()

	w := a.NewWindow("Simple File Transfer")
	w.Resize(fyne.NewSize(400, 400))
	w.SetFixedSize(true)
	w.SetMaster()

	fileWindow := a.NewWindow("Pick a file to send")
	fileWindow.Resize(fyne.NewSize(500, 400))
	fileWindow.SetCloseIntercept(fileWindow.Hide)

	ipField := InputField("IP", "127.0.0.1", "", &ip)
	portField := InputField("Port", "1337", "", &port)
	passwordField := InputField("Password", "azerty", "", &password)

	sizeText := Text("       ")

	sendUI := center(container.NewVBox(
		ipField,
		portField,
		passwordField,
		container.New(layout.NewFormLayout(),
			widget.NewButton("Choose file", func() {
				fileWindow.Show()
				dialog.ShowFileOpen(func(f fyne.URIReadCloser, e error) {
					if e != nil {
						fmt.Println("An error occured when opening the file")
					} else {
						fStats, err := os.Stat(f.URI().Path())
						if err != nil {
							fmt.Println("Couldn't open file")
						} else {
							file = f
							sizeText.SetText(sizeToString(fStats.Size()))
							fmt.Printf("Selected file '%s'\n", file.URI().Path())
						}
					}
					fileWindow.Hide()
				}, fileWindow)
			}),
			sizeText),
		widget.NewButton("Send file", func() { fmt.Println("Should send but not really atm") })))

	receiveUI := center(container.NewVBox(
		ipField,
		portField,
		passwordField,
		widget.NewButton("Retrieve", func() { fmt.Println("Should connect but not really atm") })))

	initUI := center(container.NewVBox(
		widget.NewButton("Send", func() { w.SetContent(sendUI) }),
		widget.NewButton("Receive", func() { w.SetContent(receiveUI) })))

	w.SetContent(initUI)
	w.Show()

	a.Run()
}
