package configuration

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Configuration struct {
	DSN        string `json:"DSN"`
	Host       string `json:"Host"`
	Port       int    `json:"Port"`
	User       string `json:"User"`
	Password   string `json:"Password"`
	Database   string `json:"Database"`
	SslMode    string `json:"SslMode"`
	KafkaUrl   string `json:"KafkaUrl"`
	KafkaTopic string `json:"KafkaTopic"`
}

func GetConfiguration() *Configuration {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	var config Configuration

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal(byteValue, &config)

	return &config
}
