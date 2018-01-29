package config

import (
	"fmt"
	"os"
	"testing"
)

var configuration *Config
var originalConfig string

func TestLoadConfig(t *testing.T) {
	fmt.Println("Test 1. Test `LoadConfig` function")

	fmt.Print("\t")
	fmt.Println("Step 1 : Call `LoadConfig` function")
	config, err := LoadConfig("./")

	fmt.Print("\t")
	fmt.Println("Step 2 : Checks if there is any error")
	if err != nil {
		t.Error(err.Error())
	}

	fmt.Print("\t")
	fmt.Println("Step 3 : Get `application` configuration value")
	fmt.Println(config.GetConfig("application"))
	fmt.Println()

	configuration = config
	org, _ := config.ToString()
	originalConfig = org
}

func TestConfiguration_GetConfig(t *testing.T) {
	fmt.Println("Test 2. Test `GetConfig` function")

	fmt.Print("\t")
	fmt.Println("Step 1 : Get `city.name` configuration value")
	fmt.Println(configuration.GetConfig("city.name"))

	fmt.Print("\t")
	fmt.Println("Step 2 : Change `city.name` value to `Surakarta`")
	configuration.SetConfig("city.name", "Surakarta")

	fmt.Print("\t")
	fmt.Println("Step 3 : Get new `city.name` configuration value")
	fmt.Println(configuration.GetConfig("city.name"))

	fmt.Print("\t")
	fmt.Println("Step 4 : Get non-existent configuration key")
	fmt.Println(configuration.GetConfig("asdfgqwe"))

	fmt.Println()
}

func TestConfiguration_SetConfig(t *testing.T) {
	fmt.Println("Test 3. Test `SetConfig` function")

	fmt.Print("\t")
	fmt.Println("Step 1 : Set `city.name` value to `Surakarta`")
	configuration.SetConfig("city.name", "Surakarta")

	fmt.Print("\t")
	fmt.Println("Step 2 : Get `city.name` configuration value")
	fmt.Println(configuration.GetConfig("city.name"))

	fmt.Print("\t")
	fmt.Println("Step 3 : Set non-existent configuration key")
	configuration.SetConfig("paparapa", "this is a test")

	fmt.Print("\t")
	fmt.Println("Step 4 : Get `paparapa` configuration value")
	fmt.Println(configuration.GetConfig("paparapa"))

	fmt.Println()
}

func TestConfiguration_SetLocale(t *testing.T) {
	fmt.Println("Test 4. Test `SetLocale` function")

	fmt.Print("\t")
	fmt.Println("Step 1 : Set `locale` value to `id_ID`")
	configuration.SetLocale("id_ID")

	fmt.Print("\t")
	fmt.Println("Step 2 : Get prayer name for `fajr`")
	fmt.Println(configuration.GetLocalePrayerName("fajr"))

	fmt.Print("\t")
	fmt.Println("Step 3 : Set `locale` value to `en_US`")
	configuration.SetLocale("en_US")

	fmt.Print("\t")
	fmt.Println("Step 4 : Get prayer name for `fajr`")
	fmt.Println(configuration.GetLocalePrayerName("fajr"))

	fmt.Print("\t")
	fmt.Println("Step 5 : Set `locale` value to non-existent locale settings")
	configuration.SetLocale("ar_AR")

	fmt.Print("\t")
	fmt.Println("Step 6 : Get prayer name for `fajr`")
	fmt.Println(configuration.GetLocalePrayerName("fajr"))

	fmt.Println()
}

func TestConfiguration_GetLocalePrayerName(t *testing.T) {
	fmt.Println("Test 5. Test `GetLocalePrayerName` function")

	fmt.Print("\t")
	fmt.Println("Step 1 : Set `locale` value to `id_ID`")
	configuration.SetLocale("id_ID")

	fmt.Print("\t")
	fmt.Println("Step 2 : Get prayer name for `fajr`")
	fmt.Println(configuration.GetLocalePrayerName("fajr"))

	fmt.Print("\t")
	fmt.Println("Step 3 : Get locale for non-existent prayer name")
	fmt.Println(configuration.GetLocalePrayerName("asdfg"))

	fmt.Println()
}

func TestConfiguration_ToString(t *testing.T) {
	fmt.Println("Test 6. Test `ToString` function")

	fmt.Print("\t")
	fmt.Println("Step 1 : Call `ToString` function")
	configString, err := configuration.ToString()

	fmt.Print("\t")
	fmt.Println("Step 2 : Check for any error")
	if err != nil {
		t.Error(err.Error())
	}

	fmt.Print("\t")
	fmt.Println("Step 3 : Print `ToString` result")
	fmt.Println(configString)

	fmt.Println()
}

func TestConfiguration_SaveConfig(t *testing.T) {
	fmt.Println("Test 7. Test `SaveConfig` function")

	fmt.Print("\t")
	fmt.Println("Step 1 : Set `city.name` configuration to Semarang")
	configuration.SetConfig("city.name", "Semarang")

	fmt.Print("\t")
	fmt.Println("Step 2 : Save Configuration")
	r, err := configuration.SaveConfig()

	fmt.Print("\t")
	fmt.Println("Step 3 : Check for any error")
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println(r)
	}

	fmt.Print("\t")
	fmt.Println("Step 4 : Load New Configuration")
	newConfig, _ := LoadConfig("./")

	fmt.Print("\t")
	fmt.Println("Step 5 : Get `city.name` configuration value")
	fmt.Println(newConfig.GetConfig("city.name"))
}

func TestEnd(t *testing.T) {
	file, err1 := os.OpenFile(configuration.configFile, os.O_WRONLY|os.O_TRUNC, os.ModeAppend)
	defer file.Close()

	if err1 != nil {
		fmt.Println(err1.Error())
	}

	r, err2 := file.WriteString(originalConfig)

	if err2 != nil {
		fmt.Println(err1.Error())
	}

	fmt.Println(r)
}
