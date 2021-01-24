package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type resp struct {
	Hourly []weatherdata `json:"hourly"`
}
type weatherdata struct {
	Temp float32 `json:"temp"`
}

func genResponse(lat, lon float32, date int64, key string) string {
	req := fmt.Sprintf("https://api.openweathermap.org/data/2.5/onecall/timemachine?lat=%f&lon=%f&units=imperial&dt=%d&appid=%s", lat, lon, date, key)
	return req
}

func getAverage(temps []weatherdata) float32 {
	var sum float32
	for _, entry := range temps {
		sum += entry.Temp
	}

	return sum / float32(len(temps))

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
	key := "de2e05a9ffcc0fcb5529b2f6d7f1e8db"
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

	var temp = getAverage(responseObject.Hourly)
	fmt.Println(getColor(temp))

}
