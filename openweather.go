package openweather

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	u "net/url"
	"strings"
)

const (
	BASEURL = "http://api.openweathermap.org/data/2.5/weather"
	API     = "api.txt"
)

//{"coord":{"lon":-0.13,"lat":51.51},"sys":{"message":0.0031,"country":"GB","sunrise":1402199053,"sunset":1402258504},"weather":[{"id":801,"main":"Clouds","description":"few clouds","icon":"02d"}],"base":"cmc stations","main":{"temp":296.59,"pressure":1019,"humidity":43,"temp_min":292.59,"temp_max":301.48},"wind":{"speed":5.7,"deg":200,"var_beg":160,"var_end":230,"gust":11.8},"clouds":{"all":20},"dt":1402228200,"id":2643743,"name":"London","cod":200}

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
	key, err := getApiKey()
	if err != nil {
		return nil, err
	}

	url, err := u.Parse(BASEURL)
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
	key, err := getApiKey()
	if err != nil {
		return nil, err
	}

	url, err := u.Parse(BASEURL)
	if err != nil {
		return nil, err
	}

	params := u.Values{}
	params.Add("id", id)
	params.Add("APPID", key)
	url.RawQuery = params.Encode()

	return getForecast(url)
}

func getApiKey() (string, error) {
	keybytes, err := ioutil.ReadFile(API)
	if err != nil {
		return "", err
	}
	key := string(keybytes)
	key = strings.TrimSpace(key)

	return key, nil
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
