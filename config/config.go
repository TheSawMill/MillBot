package config

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	BotToken   string `json:"bot_token"`
	GifyApiKey string `json:"gify_api_key"`
}

func GetConfiguration() *Configuration {
	// Check to see if a configuration exists
	status := configExists()

	if !status {
		createConfigFile()
		log.Fatalln("Blank configuration was created. please fill in configs and run again")
	}

	// read in configuration file

	file, err := os.Open("./config/config.json")
	if err != nil {
		log.Fatal("Error getting program confiugraionts:", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	configuration := Configuration{}

	err = decoder.Decode(&configuration)
	if err != nil {
		log.Fatalln("Error decoding JSON config:", err)
	}

	return &configuration
}

func configExists() bool {
	_, err := os.Stat("./config/config.json")
	if err == nil {
		// File exists
		return true
	}
	if os.IsNotExist(err) {
		// File does not exist
		return false
	}

	// Some other permissions error (permissions, etc,)
	log.Println("Error checking file:", err)
	return false
}

func createConfigFile() {
	// Create struct with empty values
	config := Configuration{}

	// Marshall to JSON with indentation (makes it human readable)
	data, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		log.Println("Error marshaling JSON:", err)
		return
	}

	// Create a new file
	file, err := os.Create("./config/config.json")
	if err != nil {
		log.Fatalln("Error creating configuration file:", err)
	}
	defer file.Close()

	// Write JSON to file
	_, err = file.Write(data)
	if err != nil {
		log.Fatalln("Error writing empty config to file:", err)
		return
	}

	log.Println("Empty config.json created succesfully")
}
