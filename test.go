package main

import (
	"fmt"
	"os"

	systray "github.com/getlantern/systray"
	icon "github.com/juliardi/go-prayer-reminder/assets/icon"
	config "github.com/juliardi/go-prayer-reminder/config"
)

var mQuit *systray.MenuItem

func main() {
	configuration, err := config.LoadConfig("config/")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	fmt.Println(configuration.GetConfig("city"))
	fmt.Println(configuration.GetLocalePrayerName("fajr"))
	configuration.SetLocale("en_US")
	fmt.Println(configuration.GetLocalePrayerName("fajr"))

	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(icon.Data)
	systray.SetTitle("Islamic Prayer Reminder")
	systray.SetTooltip("Islamic Prayer Reminder v1.0")

	mQuit = systray.AddMenuItem("Quit", "Quit the whole app")

	go onExit()
}

func onExit() {
	<-mQuit.ClickedCh
	systray.Quit()
	os.Exit(0)
}
