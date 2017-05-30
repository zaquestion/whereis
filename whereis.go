package whereis

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	LOCATION_API string
	PORT         string
)

func init() {
	LOCATION_API = os.Getenv("LOCATION_API")
	PORT = os.Getenv("PORT")
}

type location struct {
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	LastUpdated string  `json:"last_updated"`
	Battery     int32   `json:"battery"`
	Charging    bool    `json:"charging"`

	Destination string
}

func GetLocation(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(LOCATION_API)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	text, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = w.Write(text)
	if err != nil {
		log.Println(err)
		return
	}
}

func Run() error {
	http.HandleFunc("/getLocation", GetLocation)
	http.HandleFunc("/zaq.gif", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/zaq.gif")
	})
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

		l.Destination = r.URL.Query().Get("dest")
		if l.Destination == "" {
			l.Destination = r.URL.Query().Get("to")
		}

		t.Execute(w, l)
	})
	return http.ListenAndServe(":"+PORT, nil)
}
