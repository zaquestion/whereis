package whereis

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type location struct {
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	LastUpdated  string  `json:"last_updated"`
	BatteryLevel int     `json:"battery_remaining"`
}

func Run() error {
	LOCATION_API := os.Getenv("LOCATION_API")
	PORT := os.Getenv("PORT")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("./static/index.html")
		if err != nil {
			log.Println(err)
			return
		}

		var l location
		resp, err := http.Get(LOCATION_API)
		if err != nil {
			log.Println(err)
			return
		}
		d := json.NewDecoder(resp.Body)
		err = d.Decode(&l)
		if err != nil {
			log.Println(err)
			return
		}

		loc, err := time.LoadLocation("America/Los_Angeles")
		if err != nil {
			log.Println(err)
			return
		}

		lut, err := time.Parse("2006-01-02T15:04:05.00Z", l.LastUpdated)
		if err != nil {
			log.Println(err)
			return
		}

		lut = lut.In(loc)

		l.LastUpdated = lut.Format("Mon 03:04PM MST")

		t.Execute(w, l)
	})
	return http.ListenAndServe(":"+PORT, nil)
}
