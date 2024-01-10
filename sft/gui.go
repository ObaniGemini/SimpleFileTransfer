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

func RunGUI() {
	var ip, port, password string
	var file fyne.URIReadCloser
	var folder fyne.ListableURI

	spaceLeft := DiskSpaceLeft()
	fmt.Printf("Disk space left : %s\n", sizeToString(spaceLeft))

	a := app.New()

	w := a.NewWindow("Simple File Transfer")
	w.Resize(fyne.NewSize(500, 400))
	w.SetFixedSize(true)
	w.SetMaster()

	chooserWindow := a.NewWindow("")
	chooserWindow.Resize(fyne.NewSize(500, 400))
	chooserWindow.SetCloseIntercept(chooserWindow.Hide)

	ipField := InputField("IP", "127.0.0.1", "", &ip)
	portField := InputField("Port", "1337", "", &port)
	passwordField := InputField("Password", "azerty", "", &password)

	sizeText := Text("\t\t")
	fileSelector := container.New(layout.NewFormLayout(),
		widget.NewButton("Choose file", func() {
			chooserWindow.Show()
			dialog.ShowFileOpen(func(f fyne.URIReadCloser, e error) {
				if e != nil {
					fmt.Println("An error occured when opening the file")
				} else if f == nil {
					fmt.Println("File selection cancelled")
				} else {
					b, e2 := storage.Exists(f.URI())
					if e2 != nil || !b {
						fmt.Println("Didn't read any file")
					} else {
						fStats, e3 := os.Stat(f.URI().Path())
						if e3 != nil {
							fmt.Println("Couldn't open file")
						} else {
							file = f
							str := sizeToString(uint64(fStats.Size()))
							if len(str) >= 6 {
								sizeText.SetText(sizeToString(uint64(fStats.Size())) + "\t")
							} else {
								sizeText.SetText(sizeToString(uint64(fStats.Size())) + "\t\t")
							}
							fmt.Printf("Selected file '%s'\n", file.URI().Path())
						}
					}
				}
				chooserWindow.Hide()
			}, chooserWindow)
		}),
		sizeText)

	const chooseOutput string = "Choose an output folder"
	folderSelectText := Text(chooseOutput)
	folderSelector := widget.NewButton("Choose folder", func() {
		chooserWindow.Show()
		dialog.ShowFolderOpen(func(f fyne.ListableURI, e error) {
			if e != nil {
				fmt.Println("An error occured when opening the folder")
				folderSelectText.SetText(chooseOutput)
			} else if f == nil {
				fmt.Println("Folder selection cancelled")
			} else {
				folder = f
				folderSelectText.SetText("")
			}
			chooserWindow.Hide()
		}, chooserWindow)
	})

	sendUI := center(container.NewVBox(
		ipField,
		portField,
		passwordField,
		fileSelector,
		widget.NewButton("Send file", func() {
			if file == nil {
				fmt.Println("No file loaded to send")
			} else {
				fmt.Println("Should send but not really atm")
			}
		})))

	receiveUI := center(container.NewVBox(
		ipField,
		portField,
		passwordField,
		folderSelector,
		widget.NewButton("Retrieve", func() {
			if folder == nil {
				fmt.Println("Choose a folder mate")
			} else {
				fmt.Println("Should connect but not really atm")
			}
		}),
		folderSelectText))

	initUI := center(container.NewVBox(
		widget.NewButton("Send", func() {
			chooserWindow.SetTitle("Pick a file to send")
			w.SetContent(sendUI)
		}),
		widget.NewButton("Receive", func() {
			chooserWindow.SetTitle("Pick a folder to store the file")
			w.SetContent(receiveUI)
		})))

	w.SetContent(initUI)
	w.Show()

	a.Run()
}
