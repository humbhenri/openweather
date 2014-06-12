package openweather

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	u "net/url"
	"os"
	"strings"
)

const (
	baseurl = "http://api.openweathermap.org/data/2.5/weather"
)

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Weather struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}

type Temperature struct {
	Temp     float64 `json:"temp"`
	Pressure float64 `json:"pressure"`
	Humidity float64 `json:"humidity"`
	TempMin  float64 `json:"temp_min"`
	TempMax  float64 `json:"temp_max"`
}

type Forecast struct {
	Location Coord       `json:"coord"`
	Weathers []Weather   `json:"weather"`
	Temp     Temperature `json:"main"`
}

func GetByCityName(city string) (*Forecast, error) {
	key := getApiKey()
	url, err := u.Parse(baseurl)
	if err != nil {
		return nil, err
	}

	params := u.Values{}
	params.Add("q", city)
	params.Add("APPID", key)
	url.RawQuery = params.Encode()

	return getForecast(url)
}

func GetById(id string) (*Forecast, error) {
	key := getApiKey()
	url, err := u.Parse(baseurl)
	if err != nil {
		return nil, err
	}

	params := u.Values{}
	params.Add("id", id)
	params.Add("APPID", key)
	url.RawQuery = params.Encode()

	return getForecast(url)
}

func getApiKey() string {
	return strings.TrimSpace(os.Getenv("OPENWEATHER_API_KEY"))
}

func getForecast(url *u.URL) (*Forecast, error) {
	res, err := http.Get(url.String())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var f Forecast
	err = json.Unmarshal(body, &f)
	if err != nil {
		return nil, err
	}
	return &f, nil
}
