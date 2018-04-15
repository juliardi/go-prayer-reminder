package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	systray "github.com/getlantern/systray"
	icon "github.com/juliardi/go-prayer-reminder/assets/icon"
	config "github.com/juliardi/go-prayer-reminder/config"
)

type Application struct {
	configuration *config.Config
}

type Menu struct {
	mQuit *systray.MenuItem
}

var app *Application
var menu *Menu

func main() {
	var configPath string

	flag.StringVar(&configPath, "config", "config/", "Configuration file path")

	flag.Parse()

	configuration, err := config.LoadConfig(configPath)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	app = &Application{configuration: configuration}

	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(icon.Data)
	systray.SetTitle("Islamic Prayer Reminder")
	systray.SetTooltip("Islamic Prayer Reminder v1.0")

	cityName := app.configuration.GetConfig("city.name")

	fmt.Println("Welcome to Islamic Prayer Reminder v1.0")
	fmt.Println("Current Date : " + getCurrentDateAsString())
	fmt.Println("Prayer times for the city of", cityName)

	ptMap := getPrayerTimes()
	printPrayerTimes(ptMap)

	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

	menu = &Menu{mQuit: mQuit}

	go onExit()
}

func onExit() {
	<-menu.mQuit.ClickedCh
	systray.Quit()
	os.Exit(0)
}

// getCurrentDateAsString is used to get the current date as string
func getCurrentDateAsString() string {
	objNow := time.Now()
	return fmt.Sprintf("%d %s %d", objNow.Day(), objNow.Month().String(), objNow.Year())
}
