package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/dialog"
	"os"
	"strconv"
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

func getSize(v int64) string {
	f := func(n, d int64, end string) string {
		return strconv.FormatInt(n/d, 10) + "," + strconv.Itoa(int(10 * (float64(n)/float64(d) - float64(n/d)))) + " " + end
	}

	if v < 1024 {
		return strconv.FormatInt(v, 10) + " o"
	} else if v < (1024 * 1024) {
		return f(v, 1024, "kb") 
	} else if v < (1024 * 1024 * 1024) {
		return f(v, 1024 * 1024, "Mb")
	} else if v < (1024 * 1024 * 1024 * 1024) {
		return f(v, 1024 * 1024 * 1024, "Gb")
	} else {
		return f(v, 1024 * 1024 * 1024 * 1024, "Tb")		
	}
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

	sizeText := Text("       ")

	sendUI := center(container.NewVBox(
		InputField("IP", "127.0.0.1", "", &ip),
		InputField("Port", "1337", "", &port),
		InputField("Password", "azerty", "", &password),
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
							sizeText.SetText(getSize(fStats.Size()))
							fmt.Printf("Selected file '%s'\n", file.URI().Path())
						}
					}
					fileWindow.Hide()
				}, fileWindow)}),
			sizeText),
		widget.NewButton("Send file", func() {fmt.Println("Should send but not really atm")})))

	receiveUI := center(container.NewVBox(
		widget.NewLabel("Receive UI (not done yet)")))

	initUI := center(container.NewVBox(
		widget.NewButton("Send", func() {w.SetContent(sendUI)}),
		widget.NewButton("Receive", func() {w.SetContent(receiveUI)})))

	w.SetContent(initUI)
	w.Show()

	a.Run()
}