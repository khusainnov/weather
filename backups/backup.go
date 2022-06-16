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

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	logrus.Println("Reading config")

	if err := godotenv.Load(".env"); err != nil {
		logrus.Errorf("Cannot read .env file, due to error: %s", err.Error())
	}

	logrus.Infoln("Initializing new router")
	r := mux.NewRouter()

	r.HandleFunc("/weather/{city:[aA-zZ]+}", func(w http.ResponseWriter, r *http.Request) {
		var all entity.PageView

		/*err := r.ParseForm()
		if err != nil {
			logrus.Errorf("Cannot parse form, due to error: %s\n", err.Error())
		}*/

		urlPath := mux.Vars(r)
		logrus.Infof("%s\n", r.Form)
		city := urlPath["city"]

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

		all.Country = weatherBody.Location.Country
		all.City = weatherBody.Location.Region
		all.Airport = weatherBody.Location.Name
		all.Date = weatherBody.Location.Localtime
		all.TemperatureC = weatherBody.Current.TempC
		all.TemperatureF = weatherBody.Current.TempF

		t, err := template.ParseFiles(os.Getenv("templatePath"))
		if err != nil {
			logrus.Errorf("Cannon load and parse html files, due to error: %s\n", err.Error())
		}

		err = t.Execute(w, all)
		if err != nil {
			logrus.Errorf("Cannot execute html files, due to error: %s\n", err.Error())
		}
		/*_, err = fmt.Fprintf(w, "<b>Country: %v\n<br>City: %v\n<br>Airport: %v\n<br>Local Time: %v\n<br>Temperature: %v,C / %v,F\n</b>",
			weatherBody.Location.Country, weatherBody.Location.Region, weatherBody.Location.Name, weatherBody.Location.Localtime,
			weatherBody.Current.TempC, weatherBody.Current.TempF)
		if err != nil {
			logrus.Errorf("Cannot write data, due to error: %s", err.Error())
		}*/

	}).Methods("GET")

	r.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("method:", r.Method) //get request method
		if r.Method == "GET" {
			t, _ := template.ParseFiles("template.html")
			t.Execute(w, nil)
		} else {
			r.ParseForm()
			fmt.Println("city:", r.Form["city"])
		}
	})

	server := weather.Server{}
	logrus.Infof("Starting server on port:%s", os.Getenv("PORT"))
	if err := server.Run(os.Getenv("PORT"), r); err != nil {
		logrus.Errorf("Cannot run the server, due to error: %s\n", err.Error())
	}
}