package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	var cityID = os.Getenv("owmCityID")
	var appID = os.Getenv("owmAppID")
	if cityID == "" || appID == "" {
		os.Exit(3)
	}
	resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?id=" + cityID + "&APPID=" + appID)
	if err != nil {
		log.Println(err)
		os.Exit(4)
	}
	cp, err := current(resp.Body)
	if err != nil {
		log.Println(err)
		os.Exit(5)
	}
	resp, err = http.Get("https://api.openweathermap.org/data/2.5/forecast?id=" + cityID + "&APPID=" + appID)
	if err != nil {
		log.Println(err)
		os.Exit(6)
	}
	fp, err := forecast(resp.Body)
	if err != nil {
		log.Println(err)
		os.Exit(7)
	}
	if fp+1 < cp {
		fmt.Println("🤢👎")
	} else {
		fmt.Println("😀👍")
	}
}

func current(r io.Reader) (int, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return 0, err
	}
	var j interface{}
	err = json.Unmarshal(b, &j)
	p := j.(map[string]interface{})["main"].(map[string]interface{})["pressure"].(float64)
	return int(p), nil
}

func forecast(r io.Reader) (int, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return 0, err
	}
	var j interface{}
	err = json.Unmarshal(b, &j)
	if err != nil {
		return 0, err
	}
	p := j.(map[string]interface{})["list"].([]interface{})[1].(map[string]interface{})["main"].(map[string]interface{})["pressure"].(float64)
	return int(p), nil
}
