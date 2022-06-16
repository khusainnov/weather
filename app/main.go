package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/khusainnov/weather"
	"github.com/khusainnov/weather/internal/entity"
	"github.com/sirupsen/logrus"
)

var tml *template.Template

func init() {
	tml = template.Must(template.ParseGlob("/Users/rustamkhusainov/DocumentsAir/go-workspace/github/russodream/weather/frontend/*.gohtml"))
}

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	logrus.Println("Reading config")

	if err := godotenv.Load(".env"); err != nil {
		logrus.Errorf("Cannot read .env file, due to error: %s", err.Error())
	}

	logrus.Infoln("Initializing new router")
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tml.ExecuteTemplate(w, "form.gohtml", nil)
	})

	r.HandleFunc("/weather", Weather)

	server := weather.Server{}
	logrus.Infof("Starting server on port:%s", os.Getenv("PORT"))
	if err := server.Run(os.Getenv("PORT"), r); err != nil {
		logrus.Errorf("Cannot run the server, due to error: %s\n", err.Error())
	}
}

func Weather(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	city := r.FormValue("city")

	API := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", os.Getenv("WEATHER_API_TOKEN"), city)

	logrus.Infof("Sending request to get data from %s\n", API)
	resp, err := http.Get(API)
	if err != nil {
		logrus.Errorf("Cannot get data from %s, due to error: %s\n", API, err.Error())
	}

	logrus.Infoln("Reading response body")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("Cannot read the response body, due to error: %s\n", err.Error())
	}

	weatherBody := &entity.Weather{}

	logrus.Infoln("Unmarshalling response body")
	err = json.Unmarshal(body, weatherBody)
	if err != nil {
		logrus.Errorf("Error in unmarshalling resoonse body, err: %s", err.Error())
	}

	tml.ExecuteTemplate(w, "index.gohtml", weatherBody)
	if err != nil {
		logrus.Errorf("Cannot load and parse html files, due to error: %s\n", err.Error())
	}
}
