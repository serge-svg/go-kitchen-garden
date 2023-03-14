package main

import (
	"bytes"
	"errors"
	"fmt"
	"fyne"
	"image"
	"image/png"
	"io"
	"os"

	"fyne.io/fyne/container"
	"fyne.io/fyne/v2/canvas"
)

func (app *Config) forecastTab() *fyne.Container {
	grafic := app.getGraphic() 
	graficContainer := container.NewVBox(grafic) 	
	app.ForecastGraphicContainer = graficContainer
	return nil
}

func (app *Config) downloadFile(url, nameFile string) error {
	response, error := app.HTTPClient.Get(url)
	if error != nil {
		return error
	}

	if response.StatusCode != 200 {
		return errors.New("An error has ocurred")		
	}

	binary, error := io.ReadAll(response.Body)
	if error != nil {
		return error
	}

	defer response.Body.Close()

	img, _, error:= image.Decode(bytes.NewReader(binary))
	if error != nil {
		return error
	}
	output, error := os.Create(fmt.Sprintf("./%s, fileName"))
	if error != nil {
		return error
	}
	error = png.Encode(output, img)
	if error != nil {
		return nil
	}

}

func (app *Config) getGraphic() *canvas.Image {
	apiURL := fmt.Sprintf("https://my.meteoblue.com/visimage/meteogram_web?look=KILOMETER_PER_HOUR%2CCELSIUS%2CMILLIMETER%2Cdarkmode&apikey=5838a18e295d&temperature=C&windspeed=kmh&precipitationamount=mm&winddirection=3char&city=Montgat&iso2=es&lat=41.4686&lon=2.28001&asl=53&tz=Europe%2FMadrid&lang=es&sig=6fe10cc5c7cd9f0ddfc1c4129cd56515")
	var img *canvas.Image

	error := app.downloadFile(apiURL, "forecast.png")

	if error != nil {
		img = canvas.NewImageFromResource(resourceNodisponiblePng)
	} else {
		//Generem la imatge
		img = canvas.NewImageFromFile("pronostic.png")
	}

	img.SetMinSize(fyne.Size{
		Width: 770,
		Height: 480,
	})

	img.FillMode = canvas.ImageFillOriginal

	return img
}