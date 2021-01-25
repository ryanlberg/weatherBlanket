package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"./config"
)

type resp struct {
	Daily []weatherdata `json:"daily"`
}
type weatherdata struct {
	Wd temperature `json:"temp"`
}

type temperature struct {
	Max float32 `json:"max"`
}

func genResponse(lat, lon float32, date int64, key string) string {
	req := fmt.Sprintf("https://api.openweathermap.org/data/2.5/onecall?lat=%f&lon=%f&exclude=current,minutely,hourly,alerts&units=imperial&appid=%s", lat, lon, key)
	return req
}

func getColor(temp float32) string {
	switch {
	case temp >= 96:
		return "Parchment"
	case temp >= 91:
		return "Sunshine"
	case temp >= 86:
		return "Spice"
	case temp >= 81:
		return "Matador"
	case temp >= 76:
		return "Khaki"
	case temp >= 71:
		return "Stone"
	case temp >= 66:
		return "Empire"
	case temp >= 61:
		return "Petrol"
	default:
		return "Midnight"
	}
}

func getDate() int64 {
	time := time.Now()
	return time.AddDate(0, 0, -1).Unix()
}

func main() {
	file, err := os.Open("./config/key.config")
	if err != nil {
		panic(err)
	}
	var con config.Configuration
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&con)
	if err != nil {
		panic(err)
	}

	key := con.Key
	request := genResponse(28.5383, -81.3792, getDate(), key)

	response, err := http.Get(request)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject resp
	json.Unmarshal(data, &responseObject)

	var temp = responseObject.Daily[0].Wd.Max

	fmt.Println(temp, getColor(temp))

}
