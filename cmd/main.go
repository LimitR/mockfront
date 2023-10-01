package main

import (
	"mock/internal/configs"
	"mock/internal/server"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Mock")
	inp_file := dialog.NewFileOpen(func(uc fyne.URIReadCloser, err error) {
		c := configs.NewConfig(3)
		c.Init(uc.URI().Path())
		a := c.GetAll()
		s := server.NewServer(":3000")
		go s.Init(a)
	}, myWindow)
	btn_file := widget.NewButton("Add file api", func() {
		inp_file.Resize(fyne.NewSize(600, 600))
		inp_file.Show()
	})
	myWindow.SetContent(btn_file)
	myWindow.Resize(fyne.NewSize(600, 600))
	myWindow.ShowAndRun()
}
