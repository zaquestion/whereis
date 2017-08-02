package whereis

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

var secrets = map[string]string{
	"zaq":    os.Getenv("ZAQ_SECRET"),
	"blaise": os.Getenv("BLAISE_SECRET"),
	"leland": "",
}

func GetLocation(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("user")
	secret, ok := secrets[user]
	if !ok {
		log.Println("User:", user, "not found")
	}
	url := LOCATION_API + "?secret=" + secret
	if user == "leland" {
		url = "http://whereis.lelandbatey.com/api/v1/entry"
	}
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return
	}
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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})
	http.HandleFunc("/zaq.gif", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/zaq.gif")
	})
	http.HandleFunc("/blaise.gif", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/blaise.gif")
	})
	http.HandleFunc("/leland.gif", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/leland.gif")
	})
	http.HandleFunc("/blaiseandzaq.gif", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/blaiseandzaq.gif")
	})
	return http.ListenAndServe(":"+PORT, nil)
}
