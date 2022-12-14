package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/wooknight/jh-weather/app/weather"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Healthy %v", time.Now())
	})
	http.HandleFunc("/weather", weather.Handler)

	http.ListenAndServe(":8000", nil)
}
