package storage

import (
	"encoding/json"
	"os"
)

const UsageFile = "data/usage.json"

func SaveUsage(data map[string]uint64) error {
	file, err := os.Create(UsageFile)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(data)
}

func LoadUsage() (map[string]uint64, error) {
	file, err := os.Open(UsageFile)
	if err != nil {
		return map[string]uint64{}, nil
	}
	defer file.Close()

	var usage map[string]uint64

	if err := json.NewDecoder(file).Decode(&usage); err != nil {
		return map[string]uint64{}, nil
	}

	return usage, nil
}
