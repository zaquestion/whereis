package whereis

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
)

type location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
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

		t.Execute(w, l)
	})
	return http.ListenAndServe(":"+PORT, nil)
}
