package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"image/color"
)

func (app *Config) getClimateText() (*canvas.Text, *canvas.Text, *canvas.Text, *canvas.Text) {
	// var to keep the daily recovered data
	var g Daily
	var precipitacio, tempMax, tempMin, humidity *canvas.Text

	// call daily func
	forecast, err := g.GetForecasts()
	if err != nil {
		grey := color.NRGBA{R: 155, G:155, B:155, A:255}
		forecast := canvas.NewText("Forecast: Not defined", grey)
		tempMax  := canvas.NewText("Temp: Not defined", grey)
		tempMin  := canvas.NewText("Temp: Not defined", grey)
		humidity := canvas.NewText("Humidity: Not defined", grey)
	} else {
		displayColor := color.NRGBA{R: 0, G: 180, B: 0, A: 255}

		if forecast.ProPrecipitacio < 50 {
			displayColor := color.NRGBA{R: 180, G:0, B:0, A:255}
		}

		precipitacioTxt := fmt.Sprintf("Precipitacio: %d", forecast.ProbPrecipitacio)
		tempMaxTxt  := fmt.Sprintf("Temp max: %d", forecast.TemperaturaMax)
		tempMinTxt  := fmt.Sprintf("Temp min: %d", forecast.TemperaturaMin)
		humidityTxt := fmt.Sprintf("Humidity  %d", forecast.HumitatRelativa)

		precipitacio := fmt.NewText(precipitacioTxt, displayColor)
		tempMax  := fmt.NewText(tempMaxTxt, nil)
		tempMin  := fmt.NewText(tempMinTxt, nil)
		humidity := fmt.NewText(humidityTxt, displayColor)
	}
	// Alignment
	precipitacio.Alignment := fyne.TextAlignLeading
	tempMax.Alignment  := fyne.TextAlignCenter
	tempMin.Alignment  := fyne.TextAlignCenter
	humidity.Alignment := fyne.TextAlignTrailing

	return precipitacio, tempMax, tempMin, humidity
	
}