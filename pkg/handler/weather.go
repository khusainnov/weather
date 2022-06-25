package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/khusainnov/weather/internal/entity"
	"github.com/sirupsen/logrus"
)

var tml *template.Template

func init() {
	tml = template.Must(template.ParseGlob("/Users/rustamkhusainov/DocumentsAir/go-workspace/github/russodream/weather/frontend/*.gohtml"))
}

func (h *Handler) Form(w http.ResponseWriter, r *http.Request) {
	logrus.Infoln("Executing Form")
	err := tml.ExecuteTemplate(w, "form.gohtml", nil)
	if err != nil {
		logrus.Errorf("Cannot execute template \"form.gohtml\", due to error: %s", err.Error())
	}
}

func (h *Handler) Weather(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Redirect(w, r, "/weather/form", http.StatusSeeOther)
		return
	}

	city := r.FormValue("city")

	API := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", os.Getenv("WEATHER_API_TOKEN"), city)

	logrus.Infof("Sending request to get data from %s\n", API)
	resp, err := http.Get(API)
	if err != nil {
		logrus.Errorf("Cannot get data from %s, due to error: %s", API, err.Error())
	}

	logrus.Infoln("Reading response body")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("Cannot read the response body, due to error: %s", err.Error())
	}

	weatherBody := &entity.Weather{}

	logrus.Infoln("Unmarshalling response body")
	err = json.Unmarshal(body, weatherBody)
	if err != nil {
		logrus.Errorf("Error in unmarshalling resoonse body, err: %s", err.Error())
	}

	err = tml.ExecuteTemplate(w, "index.gohtml", weatherBody)
	if err != nil {
		logrus.Errorf("Cannot load and parse html files, due to error: %s", err.Error())
	}
}
