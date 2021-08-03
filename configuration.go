package main

import (
	"encoding/json"
	"io/ioutil"
)

type Configuration struct {
	ApiUrl string `json:"api-url"`
	MongoURI string `json:"mongo-uri"`
}

var config Configuration

func getConfig() {
	defer wg.Done()
	file, err := ioutil.ReadFile("configuration.json")
	if err != nil {
		setConfig()
		return
	}
	_ = json.Unmarshal(file, &config)
}

func setConfig()  {
	data := Configuration{ApiUrl: "/api",
		MongoURI: "uri"}

	file, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile("configuration.json", file, 0644)

	failOnError(nil, "Restart with the config completed")
}
