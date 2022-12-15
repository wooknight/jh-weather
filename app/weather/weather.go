package weather

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

/*
{"coord":{"lon":-122.3255,"lat":37.563},"weather":[{"id":801,"main":"Clouds","description":"few clouds","icon":"02n"}],"base":"stations","main":{"temp":279.68,"feels_like":279.68,"temp_min":276.75,"temp_max":282.41,"pressure":1020,"humidity":82},"visibility":10000,"wind":{"speed":0.45,"deg":215,"gust":0.89},"clouds":{"all":20},"dt":1671000136,"sys":{"type":2,"id":2002590,"country":"US","sunrise":1670944543,"sunset":1670979097},"timezone":-28800,"id":5392423,"name":"San Mateo","cod":200}
*/
// 37.5630° N, 122.3255
// (281.56K − 273.15) × 9/5 + 32 - Formula for converting K to F
// const URL = "https://api.openweathermap.org/data/2.5/weather?lat=37.5630&lon=-122.3255&appid=c99abc5102157c57ceee680e0f332bb6"
const URL = "https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s"

var APP_ID = os.Getenv("APP_ID")

func Handler(w http.ResponseWriter, r *http.Request) {
	urlParams := r.URL.Query()
	if _, ok := urlParams["lat"]; !ok {
		log.Println("Latitude is not set")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if _, ok := urlParams["lng"]; !ok {
		log.Println("Longitude is not set")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var lat, lng float64
	var err error
	if lat, err = strconv.ParseFloat(urlParams["lat"][0], 64); err != nil {
		log.Printf("Latitude is invalid . Need a float - %s \n", urlParams["lat"])
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if lng, err = strconv.ParseFloat(urlParams["lng"][0], 64); err != nil {
		log.Printf("Longitude is invalid . Need a float - %s \n", urlParams["lng"])
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	url := fmt.Sprintf(URL, lat, lng, APP_ID)
	log.Println(url)
	content := struct {
		Weather []struct {
			Id          float64 `json:"id"`
			Main        string  `json:"main"`
			Description string  `json:"description"`
			Icon        string  `json:"icon"`
		}
		Base string `json:"base"`
		Main struct {
			Temp       float64 `json:"temp"`
			Feels_like float64 `json:"feels_like"`

			Temp_min float64 `json:"temp_min"`
			Temp_max float64 `json:"temp_max"`
			Pressure float64 `json:"pressure"`
			Humidity float64 `json:"humidity"`
		}
		// Alerts []struct { //only applicable for API 3.0
		// 	Description string `json:"description"`
		// }
	}{}
	for {

		resp, err := http.Get(url)
		if resp.StatusCode != 200 {
			log.Printf("Encountered an error with the request[%s] at time %v \n", resp.Status, time.Now())
			fmt.Fprintf(w, "Encountered an error with the request[%s] at time %v \n", resp.Status, time.Now())
			return
		}
		if err != nil {
			if err.(net.Error).Temporary() {
				continue
			}
			fmt.Fprintf(w, "Encountered an error %v at time %v \n", err, time.Now())
		}
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&content)
		if err != nil {
			fmt.Fprintf(w, "Encountered an error %v at time %v \n", err, time.Now())
			return
		}
		log.Println(content)

		fmt.Fprintf(w, "Content %v - \n%s\n", content.Main, curTemp(content.Main.Temp))

		fmt.Fprintf(w, "Weather  %v", content.Weather[0].Description)
		resp.Body.Close()
		break
	}
}

func curTemp(K float64) string {
	temp := (K-273.15)*9/5 + 32
	switch {
	case temp >= 80:
		return "It is hot right now"
	case temp <= 40:
		return "freezing!!"
	case temp <= 50:
		return "It is cold right now"
	case temp > 50 && temp < 80:
		return "pleasantly moderate"
	default:
		return "unknown temp"
	}
}
