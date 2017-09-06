package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	godotenv "github.com/joho/godotenv"
	praytimes "github.com/juliardi/PrayTimes-Golang"
)

var cityName string
var cityLat float64
var cityLong float64
var cityTimezone int
var azanFile string

func main() {
	loadConfig()
	currentDate := getCurrentDateAsString()
	ptMap := prayTime(cityLat, cityLong, cityTimezone)

	fmt.Println("Prayer time schedule for", cityName, "on", currentDate)
	printPrayTime(ptMap)
	ticker := timeTicker(ptMap, azanFile)
	mainloop(ticker)
}

// Function mainloop is used to block the program execution
// and provide a pretty way of program to quit
func mainloop(ticker time.Ticker) {
	exitSignal := make(chan os.Signal)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-exitSignal

	ticker.Stop()
	fmt.Println("Program Exit")
	os.Exit(0)
}

// Function loadConfig is used to load program configuration
// in .env file
// NOTE : You must always put your .env file in the same directory as the program
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
	if azanFile == "" {
		log.Fatal("AZAN_FILENAME configuration is missing")
	}
}

// This is used to get the current date as string
func getCurrentDateAsString() string {
	objNow := time.Now()
	return fmt.Sprintf("%d %s %d", objNow.Day(), objNow.Month().String(), objNow.Year())
}

// Function timeTicker always check the time every minute and
// compares it with prayer time schedule. When the time is match
// with one of the prayer time, it calls playAzan function
func timeTicker(ptMap map[string]string, azanFile string) time.Ticker {
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

	return *ticker
}

// Function playAzan will plays the Azan MP3 file
// It needs `mp123` library to plays MP3 file
func playAzan(azanFile string) {
	cmd := exec.Command("mpg123", azanFile)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

// Function prayTime retrieves prayer time schedule from
// PrayTimes(https://github.com/3ace/PrayTimes-Golang) library
func prayTime(latitude float64, longitude float64, timezone int) map[string]string {
	pt := praytimes.GetTimes(time.Now(), []float64{latitude, longitude}, timezone)

	return pt
}

// This function prints out prayer time schedule
func printPrayTime(ptMap map[string]string) {
	fmt.Println("midnight =", ptMap["midnight"])
	fmt.Println("imsak =", ptMap["imsak"])
	fmt.Println("fajr =", ptMap["fajr"])
	fmt.Println("sunrise =", ptMap["sunrise"])
	fmt.Println("dhuhr =", ptMap["dhuhr"])
	fmt.Println("asr =", ptMap["asr"])
	fmt.Println("maghrib =", ptMap["maghrib"])
	fmt.Println("isha =", ptMap["isha"])
}
