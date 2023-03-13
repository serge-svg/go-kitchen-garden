package main

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	_ "fyne.io/fyne/v2/widget"
)

type Config struct {
	App      fyne.App
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

var appConfig Config

func main() {

	// Create fyne app
	fyneApp := app.NewWithID("org.svg.ecokitchen-garden")
	appConfig.App = fyneApp

	// Call logs
	appConfig.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	appConfig.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Lshortfile)

	// Connect db

	// Create a db repository

	// Create and define size of the fyne screen
	win := fyneApp.NewWindow("Eco ktchen-garden App")
	// Show and execute fyne app
	win.ShowAndRun()
	/*
		a := app.New()
		w := a.NewWindow("Hi you")
		w.SetContent(widget.NewLabel("Hi you"))
		w.ShowAndRun()
	*/

	// go mod tidy // force update dependencies
}
