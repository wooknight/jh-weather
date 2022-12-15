package main

import (
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
		weather.OutputJSON(resp, w)
	})
	http.HandleFunc("/weather", weather.Handler)

	http.ListenAndServe(":80", nil)
}
