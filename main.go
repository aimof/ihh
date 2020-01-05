package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/aimof/jaason"
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
	v, err := jaason.Parse(b)
	if err != nil {
		return 0, err
	}
	p, err := v.Get("main").Get("pressure").Int()
	return p, err
}

func forecast(r io.Reader) (int, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return 0, err
	}
	v, err := jaason.Parse(b)
	if err != nil {
		return 0, err
	}
	p, err := v.Get("list").Get(1).Get("main").Get("pressure").Int()
	return p, err
}
