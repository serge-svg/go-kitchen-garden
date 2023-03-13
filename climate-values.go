package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type PreUrl struct {
	Url    string `json:"datos"` // url
	Client *http.Client
}

func main() {
	result, _ := GetPreUrl()
	fmt.Println(result)

}

func GetPreUrl() (string, error) {
	url := "https://opendata.aemet.es/opendata/api/prediccion/especifica/municipio/diaria/08001/?api_key=eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJzZXJnaXZhcXVlQGdtYWlsLmNvbSIsImp0aSI6ImY5NDA4MTllLTM2MmMtNDQ5OC1iMDA3LTY1ODZmOWYyOGUxZCIsImlzcyI6IkFFTUVUIiwiaWF0IjoxNjc4NzI3NTE4LCJ1c2VySWQiOiJmOTQwODE5ZS0zNjJjLTQ0OTgtYjAwNy02NTg2ZjlmMjhlMWQiLCJyb2xlIjoiIn0.rJSXW8pOwtr4T_MC1YsHWQtRJJHoiZS4YXvSJFhKwss"

	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("cache-control", "no-cache")
	response, error := http.DefaultClient.Do(request)
	if error != nil {
		log.Println("Error connecting with AEMET")
	}

	defer response.Body.Close()
	body, error := ioutil.ReadAll(response.Body)

	if error != nil {
		log.Println("Error reading json")
	}

	preUrl := PreUrl{}
	error = json.Unmarshal(body, &preUrl)

	if error != nil {
		log.Println("Error unmarshal")
	}

	return preUrl.Url, error
}
