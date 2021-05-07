package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"time"
)

/*
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
	Time time.Time
	SunRize time.Time
	SunSet time.Time
	TimezoneName string
	Sources Sources
	Title string
	LocationType string
	Woeid int
	LattLong float32
	Timezone string

}
*/
/*
type Autogenerate struct {
	ConsolidatedWeather []struct {
		ID                   int64     `json:"id"`
		WeatherStateName     string    `json:"weather_state_name"`
		WeatherStateAbbr     string    `json:"weather_state_abbr"`
		WindDirectionCompass string    `json:"wind_direction_compass"`
		Created              time.Time `json:"created"`
		ApplicableDate       string    `json:"applicable_date"`
		MinTemp              float64   `json:"min_temp"`
		MaxTemp              float64   `json:"max_temp"`
		TheTemp              float64   `json:"the_temp"`
		WindSpeed            float64   `json:"wind_speed"`
		WindDirection        float64   `json:"wind_direction"`
		AirPressure          float64   `json:"air_pressure"`
		Humidity             int       `json:"humidity"`
		Visibility           float64   `json:"visibility"`
		Predictability       int       `json:"predictability"`
	} `json:"consolidated_weather"`
	Time         time.Time `json:"time"`
	SunRise      time.Time `json:"sun_rise"`
	SunSet       time.Time `json:"sun_set"`
	TimezoneName string    `json:"timezone_name"`
	Parent       struct {
		Title        string `json:"title"`
		LocationType string `json:"location_type"`
		Woeid        int    `json:"woeid"`
		LattLong     string `json:"latt_long"`
	} `json:"parent"`
	Sources []struct {
		Title     string `json:"title"`
		Slug      string `json:"slug"`
		URL       string `json:"url"`
		CrawlRate int    `json:"crawl_rate"`
	} `json:"sources"`
	Title        string `json:"title"`
	LocationType string `json:"location_type"`
	Woeid        int    `json:"woeid"`
	LattLong     string `json:"latt_long"`
	Timezone     string `json:"timezone"`
}*/
/*
type ConsolidatedWeather []struct {
	ID                   int64     `json:"id"`
	WeatherStateName     string    `json:"weather_state_name"`
	WeatherStateAbbr     string    `json:"weather_state_abbr"`
	WindDirectionCompass string    `json:"wind_direction_compass"`
	Created              time.Time `json:"created"`
	ApplicableDate       string    `json:"applicable_date"`
	MinTemp              float64   `json:"min_temp"`
	MaxTemp              float64   `json:"max_temp"`
	TheTemp              float64   `json:"the_temp"`
	WindSpeed            float64   `json:"wind_speed"`
	WindDirection        float64   `json:"wind_direction"`
	AirPressure          float64   `json:"air_pressure"`
	Humidity             int       `json:"humidity"`
	Visibility           float64   `json:"visibility"`
	Predictability       int       `json:"predictability"`
}

type Weather struct {
	ConsolidatedWeather ConsolidatedWeather `json:"consolidated_weather"`
	Time                time.Time           `json:"time"`
	SunRise             time.Time           `json:"sun_rise"`
	SunSet              time.Time           `json:"sun_set"`
	TimezoneName        string              `json:"timezone_name"`
	Parent              struct {
		Title        string `json:"title"`
		LocationType string `json:"location_type"`
		Woeid        int    `json:"woeid"`
		LattLong     string `json:"latt_long"`
	} `json:"parent"`
	Sources []struct {
		Title     string `json:"title"`
		Slug      string `json:"slug"`
		URL       string `json:"url"`
		CrawlRate int    `json:"crawl_rate"`
	} `json:"sources"`
	Title        string `json:"title"`
	LocationType string `json:"location_type"`
	Woeid        int    `json:"woeid"`
	LattLong     string `json:"latt_long"`
	Timezone     string `json:"timezone"`
}
*/

type Weather struct {
	ConsolidatedWeather []ConsolidatedWeather `json:"consolidated_weather"`
	Time                time.Time             `json:"time"`
	SunRise             time.Time             `json:"sun_rise"`
	SunSet              time.Time             `json:"sun_set"`
	TimezoneName        string                `json:"timezone_name"`
	Parent              Parent                `json:"parent"`
	Sources             []Sources             `json:"sources"`
	Title               string                `json:"title"`
	LocationType        string                `json:"location_type"`
	Woeid               int                   `json:"woeid"`
	LattLong            string                `json:"latt_long"`
	Timezone            string                `json:"timezone"`
}
type ConsolidatedWeather struct {
	ID                   int64     `json:"id"`
	WeatherStateName     string    `json:"weather_state_name"`
	WeatherStateAbbr     string    `json:"weather_state_abbr"`
	WindDirectionCompass string    `json:"wind_direction_compass"`
	Created              time.Time `json:"created"`
	ApplicableDate       string    `json:"applicable_date"`
	MinTemp              float64   `json:"min_temp"`
	MaxTemp              float64   `json:"max_temp"`
	TheTemp              float64   `json:"the_temp"`
	WindSpeed            float64   `json:"wind_speed"`
	WindDirection        float64   `json:"wind_direction"`
	AirPressure          float64   `json:"air_pressure"`
	Humidity             int       `json:"humidity"`
	Visibility           float64   `json:"visibility"`
	Predictability       int       `json:"predictability"`
}
type Parent struct {
	Title        string `json:"title"`
	LocationType string `json:"location_type"`
	Woeid        int    `json:"woeid"`
	LattLong     string `json:"latt_long"`
}
type Sources struct {
	Title     string `json:"title"`
	Slug      string `json:"slug"`
	URL       string `json:"url"`
	CrawlRate int    `json:"crawl_rate"`
}

func main() {
	url := "https://www.metaweather.com/api/location/2122265"

	client := http.Client{
		Timeout: 6 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	//io.Copy(os.Stdout, resp.Body)

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	//fmt.Println(body)

	//metaweather := weather{}

	var Metaweather Weather

	jsonErr := json.Unmarshal(body, &Metaweather)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	slWeather := Metaweather.ConsolidatedWeather

	cWeather := ConsolidatedWeather{}

	for i := 0; i < 6; i++ {
		w := reflect.ValueOf(slWeather[i])
		cWeather.ID = reflect.Indirect(w).FieldByName("ID")

		fmt.Println(id)
	}

	fmt.Println(cWeather.ID)

	//	io.Copy(os.Stdout, resp.Body)
}
