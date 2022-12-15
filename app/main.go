package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/wooknight/jh-weather/app/weather"
)

var build = "develop"

var startTime time.Time

func uptime() time.Duration {
	return time.Since(startTime)
}

func init() {
	startTime = time.Now()
}

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["health"] = "healthy"
		resp["version"] = build
		resp["uptime"] = uptime().String()
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Printf("Error happened in JSON marshal. Err: %s\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "encountered an error")
			return
		}
		w.Write(jsonResp)
	})
	http.HandleFunc("/weather", weather.Handler)

	http.ListenAndServe(":80", nil)
}
