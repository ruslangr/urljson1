package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang/glog"
	_ "github.com/lib/pq"
)

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
type fromDb struct {
	ID                   int64
	WeatherStateName     string
	WindDirectionCompass string
	Created              time.Time
	ApplicableDate       string
	MinTemp              float64
	MaxTemp              float64
	TheTemp              float64
}

var database *sql.DB

func main() {
	Url := "https://www.metaweather.com/api/location/2122265"

	bodyByte := getUrl(Url)

	mWeather := getJson(bodyByte)

	userDb := os.Getenv("userdb")
	passDb := os.Getenv("passdb")
	nameDb := os.Getenv("namedb")
	sslmodeDb := os.Getenv("sslmodedb")
	hostDb := os.Getenv("hostdb")

	connStr := "user=" + userDb + " " + "password=" + passDb + " " + "dbname=" + nameDb + " " + "sslmode=" + sslmodeDb + " " + "host=" + hostDb
	//fmt.Println(connStr)

	db, err := sql.Open("postgres", string(connStr))
	if err != nil {
		panic(err)
	}
	database = db
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connection OK")
	}

	//first write to DB
	writeToDb(mWeather, db)

	go getAndWrite(Url, db)

	http.HandleFunc("/create", CreateHandler)
	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}

func getUrl(url string) []byte {
	client := http.Client{
		Timeout: 6 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	return body
}
func getJson(body []byte) Weather {
	var metaweather Weather

	jsonErr := json.Unmarshal(body, &metaweather)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return metaweather
}
func writeToDb(mW Weather, db *sql.DB) {
	for i := 0; i < len(mW.ConsolidatedWeather); i++ {
		_, err := db.Exec("insert into Weather (ID, WeatherStateName, WindDirectionCompass, Created, ApplicableDate, MinTemp, MaxTemp, TheTemp) values ($1, $2, $3, $4, $5, $6, $7, $8)", mW.ConsolidatedWeather[i].ID, mW.ConsolidatedWeather[i].WeatherStateName, mW.ConsolidatedWeather[i].WindDirectionCompass, mW.ConsolidatedWeather[i].Created, mW.ConsolidatedWeather[i].ApplicableDate, mW.ConsolidatedWeather[i].MinTemp, mW.ConsolidatedWeather[i].MaxTemp, mW.ConsolidatedWeather[i].TheTemp)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Write to DB OK!")
}
func IndexHandler(w http.ResponseWriter, r *http.Request) {

	rows, err := database.Query("SELECT * FROM Weather WHERE ApplicableDate = $1 ORDER BY Created", "2021-05-09")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	qWeather := []fromDb{}

	for rows.Next() {
		p := fromDb{}
		errRows := rows.Scan(&p.ID, &p.WeatherStateName, &p.WindDirectionCompass, &p.Created, &p.ApplicableDate, &p.MinTemp, &p.MaxTemp, &p.TheTemp)
		if errRows != nil {
			fmt.Println(errRows)
			continue
		}
		qWeather = append(qWeather, p)
	}
	tmpl, _ := template.ParseFiles("template/index.html")
	tmpl.Execute(w, qWeather)
}
func rowExists(query string, args int64) bool {
	var exists bool
	query = fmt.Sprintf("SELECT exists (%s)", query)
	err := database.QueryRow(query, args).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		glog.Fatalf("error checking if row exists '%s' %v", args, err)
	}
	return exists
}
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		date := r.FormValue("date")

		rows, err1 := database.Query("SELECT * FROM Weather WHERE ApplicableDate = $1 ORDER BY Created", date)
		if err1 != nil {
			log.Println(err)
		}
		defer rows.Close()
		qWeather := []fromDb{}

		for rows.Next() {
			p := fromDb{}
			errRows := rows.Scan(&p.ID, &p.WeatherStateName, &p.WindDirectionCompass, &p.Created, &p.ApplicableDate, &p.MinTemp, &p.MaxTemp, &p.TheTemp)
			if errRows != nil {
				fmt.Println(errRows)
				continue
			}
			qWeather = append(qWeather, p)
		}
		tmpl, _ := template.ParseFiles("template/index.html")
		tmpl.Execute(w, qWeather)
		http.Redirect(w, r, "/", 301)
		//		http.Redirect(w, r, "/", 301)
	} else {
		http.ServeFile(w, r, "template/create.html")
	}
}
func getAndWrite(url string, db *sql.DB) {
	for {
		t := time.Now()
		cTime := t.Format(time.Kitchen)

		if cTime == "9:01AM" || cTime == "12:01PM" || cTime == "3:01PM" || cTime == "6:01PM" || cTime == "9:01PM" || cTime == "00:01AM" || cTime == "3:01AM" || cTime == "6:01AM" {
			bodyByte := getUrl(url)

			mWeather := getJson(bodyByte)

			//check if row not exist - write to DB
			for i := 0; i <= len(mWeather.ConsolidatedWeather); i++ {
				fl := rowExists("SELECT ID FROM Weather WHERE ID=$1", mWeather.ConsolidatedWeather[i].ID)
				if fl == true {
					break
				}
				if fl == false {
					writeToDb(mWeather, db)
				}
			}
		}
	}
}
