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
}

func GetLocation(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("user")
	secret, ok := secrets[user]
	if !ok {
		log.Println("User:", user, "not found")
	}
	resp, err := http.Get(LOCATION_API + "?secret=" + secret)
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
	http.Handle("/", http.FileServer(http.Dir("static")))
	return http.ListenAndServe(":"+PORT, nil)
}
