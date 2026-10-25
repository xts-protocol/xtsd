package tenant

import (
	"encoding/json"
	"os"
)

const TenantFile = "data/tenants.json"

func SaveRegistry() error {

	file, err := os.Create(TenantFile)
	if err != nil {
		return err
	}

	defer file.Close()

	return json.NewEncoder(file).
		Encode(Registry)
}

func LoadRegistry() error {

	file, err := os.Open(TenantFile)
	if err != nil {
		return nil
	}

	defer file.Close()

	var data map[string]Tenant

	if err := json.NewDecoder(file).
		Decode(&data); err != nil {

		return nil
	}

	Registry = data

	return nil
}
