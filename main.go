package main

import (
	"log"
	"net/http"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	_ "fyne.io/fyne/v2/widget"
)

type Config struct {
	App                      fyne.App
	InfoLog                  *log.Logger
	ErrorLog                 *log.Logger
	MainWindow               fyne.Window // reference to main screen
	ClimateDataContainer     *fyne.Container
	HTTPClient               http.Client
	ForecastGraphicContainer *fyne.Container
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
	appConfig.MainWindow = fyneApp.NewWindow("Eco ktchen-garden App")
	appConfig.MainWindow.Resize(fyne.NewSize(800, 600))
	appConfig.MainWindow.SetFixedSize(true)
	appConfig.MainWindow.SetMaster()

	appConfig.makeUI() //

	// Show and execute fyne app
	appConfig.MainWindow.ShowAndRun()

	// go mod tidy // force update dependencies
}
