package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"time"

	praytimes "github.com/juliardi/PrayTimes-Golang"
	"github.com/sqweek/dialog"
)

// timeTicker always check the time every minute and
// compares it with prayer time schedule. When the time is match
// with one of the prayer time, it calls playAzan function
func timeTicker(ptMap map[string]string) time.Ticker {
	azanFile := app.configuration.GetConfig("application.azan_file").(string)
	ticker := time.NewTicker(time.Minute)

	go func() {
		for t := range ticker.C {

			fmt.Println("Tick at ", t)
			strTime := strconv.Itoa(t.Hour()) + ":" + strconv.Itoa(t.Minute())

			for strSholat := range ptMap {
				if strTime == ptMap[strSholat] {
					message := fmt.Sprintf("Now is the time for %s prayer", strSholat)
					fmt.Println(message)
					go playAzan(azanFile)
					dialog.Message("%s", message).Title("Sholat").OkDialog()
					time.Sleep(time.Minute * 2)
				}
			}
		}
	}()

	return *ticker
}

// playAzan plays the Azan MP3 file using `mp123` library
func playAzan(azanFile string) {
	cmd := exec.Command("mpg123", azanFile)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

// getPrayerTimes retrieves prayer time schedule from PrayTimes(https://github.com/3ace/PrayTimes-Golang) library
func getPrayerTimes() map[string]string {
	method := app.configuration.GetConfig("application.calculation_method")
	cityLat := app.configuration.GetConfig("city.latitude")
	cityLong := app.configuration.GetConfig("city.longitude")
	cityTimezone := app.configuration.GetConfig("city.timezone")

	latitude, _ := strconv.ParseFloat((cityLat).(string), 64)
	longitude, _ := strconv.ParseFloat((cityLong).(string), 64)
	timezone, _ := strconv.ParseFloat(cityTimezone.(string), 64)

	praytimes.SetMethod(method.(string))
	coordinate := []float64{
		latitude,
		longitude,
	}
	pt := praytimes.GetTimes(time.Now(), coordinate, timezone)

	return pt
}

// printPrayerTimes prints out prayer time schedule
func printPrayerTimes(ptMap map[string]string) {
	prayerList := [7]string{"imsak", "fajr", "sunrise", "dhuhr", "asr", "maghrib", "isha"}

	for index := 0; index < len(prayerList); index++ {
		prayerName := prayerList[index]
		localePrayerName := app.configuration.GetLocalePrayerName(prayerName)
		fmt.Println(fmt.Sprintf("%s = %s", localePrayerName, ptMap[prayerName]))
	}
}
