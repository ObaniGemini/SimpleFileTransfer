package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/layout"
    "fyne.io/fyne/v2/widget"
)

func center(c *fyne.Container) *fyne.Container {
	return container.New(layout.NewCenterLayout(), c)
}

func InitGUI() {
	minSize := fyne.NewSize(300, 300)

	a := app.New()
    w := a.NewWindow("Simple File Transfer")

    w.Resize(minSize)

    sendUI := center(container.NewVBox(
    	widget.NewLabel("Send UI (not done yet)"),
    ))

    receiveUI := center(container.NewVBox(
    	widget.NewLabel("Receive UI (not done yet)"),
    ))

    initUI := center(container.NewVBox(
    	widget.NewButton("Send", func() {w.SetContent(sendUI)}),
    	widget.NewButton("Receive", func() {w.SetContent(receiveUI)}),
    ))

    w.SetContent(initUI)
    w.ShowAndRun()
}