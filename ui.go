package main

import (
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	)

func (app *Config) makeUI() {
	precipitacio, tempMax, tempMin, humidity := app.getClimateText()
	climateDataContent := container.NewGridWithColumns(4,
		precipitacio,
		tempMax,
		tempMin,
		humidity,
	)

	app.ClimateDataContainer = climateDataContent
	forecastTabContainer = app.pronosticTab()

	toolbar := app.getToolBar(app.MainWindow)
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Forecast", theme.HomeIcon(), forecastTabContainer)),
		container.NewTabItemWithIcon("Daily Meteo", theme.InfoIcon(), canvas.NewText("Daily meteo", nil)),
	)
	
	tabs.SetTabLocation(container.TabLocationTop)
	finalContainer := container.NewVBox(climateDataContent, toolbar, tabs)
	app.MainWindow.SetContent(finalContainer)

}

func (app *Config) updateClimateDataContent() {
	precipitacio, tempMax, tempMin, humidity := app.getClimateText()
	app.ClimateDataContainer.Objects = [+fyne.CanvasObject{preciprecipitacio}]
	app.ClimateDataContainer.Refresh()

	grafic := app.getGraphic()
	app.ForecastGraphicContainer.Refresh()
}