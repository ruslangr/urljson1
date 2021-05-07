package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

/*type Weather struct {
	//	Id                 string `json:"id"`
	//	Weather_state_name string `json:"weather_state_name"`
	ConsolidatedWeather []string `json:"consolidated_weather"`
}*/

type ConsolidatedWeather struct {
	Id                   int     `json:id`
	WeatherStateName     string  `json:weather_state_name`
	WindDirectionCompass string  `json:wind_direction_compass`
	Created              string  `json:created`
	ApplicableDate       string  `json:applicable_date`
	MinTemp              float32 `json:min_temp`
	MaxTemp              float32 `json:max_temp`
}

type Weather struct {
	ConsolidatedWeather ConsolidatedWeather `json:consolidated_weather`
}

func main() {

	url := "https://www.metaweather.com/api/location/2122265"

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	//fmt.Println(res.Body)
	//io.Copy(os.Stdout, res.Body)

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	//fmt.Println(body)

	//metaweather := Weather{}

	var metaweather Weather
	jsonErr := json.Unmarshal(body, &metaweather)
	if jsonErr != nil {
		fmt.Print(jsonErr)
		//	log.Fatal(jsonErr)
	}
	fmt.Printf("%+v", metaweather)
	fmt.Println(metaweather.ConsolidatedWeather)
}
