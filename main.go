package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"

	praytimes "github.com/3ace/PrayTimes-Golang"
	godotenv "github.com/joho/godotenv"
)

var cityName string
var cityLat float64
var cityLong float64
var cityTimezone int
var azanFile string

func main() {
	loadConfig()
	currentDate := getCurrentDateAsString()
	fmt.Println("Prayer time schedule for", cityName, "on", currentDate)
	ptMap := prayTime(cityLat, cityLong, cityTimezone)
	printPrayTime(ptMap)
	timeTicker(ptMap, azanFile)
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cityName = os.Getenv("CITY_NAME")
	cityLat, err = strconv.ParseFloat(os.Getenv("CITY_LAT"), 64)
	if err != nil {
		log.Fatal("CITY_LAT configuration is missing")
	}

	cityLong, err = strconv.ParseFloat(os.Getenv("CITY_LONG"), 64)
	if err != nil {
		log.Fatal("CITY_LONG configuration is missing")
	}

	cityTimezone, err = strconv.Atoi(os.Getenv("CITY_TIMEZONE"))
	if err != nil {
		log.Fatal("CITY_TIMEZONE configuration is missing")
	}

	azanFile = os.Getenv("AZAN_FILENAME")
}

func getCurrentDateAsString() string {
	objNow := time.Now()
	return fmt.Sprintf("%d %s %d", objNow.Day(), objNow.Month().String(), objNow.Year())
}

func timeTicker(ptMap map[string]string, azanFile string) {
	ticker := time.NewTicker(time.Minute)

	go func() {
		for t := range ticker.C {

			fmt.Println("Tick at ", t)
			strTime := strconv.Itoa(t.Hour()) + ":" + strconv.Itoa(t.Minute())

			for strSholat := range ptMap {
				if strTime == ptMap[strSholat] {
					fmt.Println("Now is the time for", strSholat, "prayer")
					playAzan(azanFile)
					time.Sleep(time.Minute * 2)
				}
			}
		}
	}()

	time.Sleep(time.Minute * 60)
	ticker.Stop()

}

func playAzan(azanFile string) {
	cmd := exec.Command("mpg123", azanFile)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func prayTime(latitude float64, longitude float64, timezone int) map[string]string {
	pt := praytimes.GetTimes(time.Now(), []float64{latitude, longitude}, timezone)

	return pt
}

func printPrayTime(ptMap map[string]string) {
	fmt.Print("midnight=", ptMap["midnight"], " | ")
	fmt.Print("imsak=", ptMap["imsak"], " | ")
	fmt.Print("fajr=", ptMap["fajr"], " | ")
	fmt.Print("sunrise=", ptMap["sunrise"], " | ")
	fmt.Print("dhuhr=", ptMap["dhuhr"], " | ")
	fmt.Print("asr=", ptMap["asr"], " | ")
	fmt.Print("maghrib=", ptMap["maghrib"], " | ")
	fmt.Print("isha=", ptMap["isha"])
	fmt.Print("\n")
}
