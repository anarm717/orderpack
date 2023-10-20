package config

import (
	"GoOrderPackProject/api/models"
	"encoding/json"
	"log"
	"os"
)

// OrderRequest represents the model for an OrderRequest
type OrderRequest struct {
	OrderQuantity int `json:"OrderQuantity"`
}

// OrderResponse represents the model for an OrderResponse
type OrderResponse struct {
	PackSize int `json:"PackSize"`
	Count    int `json:"Count"`
}

func GetPackSizes() []int {
	return ReadConfig().PackSizes
}

func SetPackSizes(newPackSizes []int) []int {
	configParams := ReadConfig()
	configParams.PackSizes = newPackSizes
	file, _ := json.MarshalIndent(configParams, "", " ")

	err := os.WriteFile("config.json", file, 0644)
	if err != nil {
		log.Println(err)
	}
	return ReadConfig().PackSizes
}

func ReadConfig() *models.Config {
	data, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	var config models.Config
	if err := json.Unmarshal(data, &config); err != nil {
		log.Fatal(err)
	}
	configureLogger(config)
	return &config
}

func configureLogger(logConfig models.Config) (*os.File, error) {
	logFile, err := os.OpenFile(logConfig.LogFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	log.SetOutput(logFile)

	// Set the log level based on the configuration.
	switch logConfig.LogLevel {
	case "info":
		log.SetFlags(log.LstdFlags)
	case "debug":
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	}

	return logFile, nil
}
