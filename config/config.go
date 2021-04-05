package config

import (
	"encoding/json"
	"os"
)

//configuration file for service
type Config struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func LoadConfiguration(fileName string) (Config, error)  {
	var config Config
	configFile, err := os.Open(fileName)
	defer configFile.Close()
	if err != nil{
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	return config, err

}
