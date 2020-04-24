package utils

import (
	"os"

	"gopkg.in/yaml.v2"
)

//Config is a struct that allows for easy interation with yaml stored data.
type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	Database struct {
		URI      string `yaml:"uri"`
		Username string `yaml:"user"`
		Password string `yaml:"pass"`
	}
}

//GetConfigurations retries data from the yaml file in the local directory, and builds out a Config struct
func GetConfigurations(cfg *Config) {
	const configFile = "config.yaml"
	f, err := os.Open(configFile)
	if err != nil {
		DefaultLogger.Error.Printf("An error occured when opening file: %s", configFile)
		DefaultLogger.Error.Printf("%s may not exist in the directory where the progam was called.", configFile)
		ProcessError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		DefaultLogger.Error.Println("Error occured when decoding yaml file.")
		ProcessError(err)
	}
}

//ProcessError handles errors by printing them and performing an os.exit.
func ProcessError(err error) {
	DefaultLogger.Error.Println(err)
	os.Exit(2)
}
