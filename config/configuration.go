package config

import (
	"os"

	toml "github.com/pelletier/go-toml"
)

// Config is a struct for configuration
type Config struct {
	storage             *toml.Tree
	localizationStorage *toml.Tree
	configFile          string
}

// LoadConfig is used to load configuration from toml configuration file
func LoadConfig(configDir string) (*Config, error) {
	configFile := configDir + "main.toml"
	config, err := toml.LoadFile(configFile)

	if err != nil {
		return nil, err
	}

	localizationStorage := loadLocaleConfig(configDir)

	return &Config{storage: config, localizationStorage: localizationStorage, configFile: configFile}, nil
}

func loadLocaleConfig(configDir string) *toml.Tree {
	configFile := configDir + "locale.toml"
	localization, err := toml.LoadFile(configFile)

	if err != nil {
		panic(err.Error())
	}

	return localization
}

// SaveConfig is used to save configuration into configuration file
func (configuration *Config) SaveConfig() (int, error) {
	file, err1 := os.OpenFile(configuration.configFile, os.O_WRONLY|os.O_TRUNC, os.ModeAppend)
	defer file.Close()

	if err1 != nil {
		return 0, err1
	}

	strToml, _ := configuration.ToString()

	r, err2 := file.WriteString(strToml)

	if err2 != nil {
		return r, err2
	}

	return r, nil
}

// GetConfig is used to retrieves value of configuration by key
func (configuration *Config) GetConfig(key string) interface{} {
	return configuration.storage.Get(key)
}

// SetConfig is used to set configuration value
func (configuration *Config) SetConfig(key string, value string) {
	configuration.storage.Set(key, value)
}

// ToString is used to get the string representation of the configuration
func (configuration *Config) ToString() (string, error) {
	return configuration.storage.ToTomlString()
}

// SetLocale is used to set application language settings
func (configuration *Config) SetLocale(localeID string) {
	configuration.SetConfig("application.locale", localeID)
}

// GetLocalePrayerName is used to get prayer name according to locale settings
func (configuration *Config) GetLocalePrayerName(prayerName string) string {
	localeID := configuration.GetConfig("application.locale").(string)
	keyName := localeID + "." + prayerName

	if configuration.localizationStorage.Has(keyName) {
		return configuration.localizationStorage.Get(keyName).(string)
	}

	return prayerName
}
