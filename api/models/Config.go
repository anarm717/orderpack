package models

type Config struct {
	PackSizes   []int  `json:"packSizes"`
	LogFileName string `json:"LogFileName"`
	LogLevel    string `json:"LogLevel"`
	HttpPort    string `json:"httpPort"`
}