package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	NodeName        string  `json:"node_name"`
	ListenPort      string  `json:"listen_port"`
	Environment     string  `json:"environment"`
	PricePerMessage float64 `json:"price_per_message"`
}

var App Config

func Load() error {

	file, err := os.Open("configs/config.json")
	if err != nil {
		return err
	}

	defer file.Close()

	return json.NewDecoder(file).Decode(&App)
}
