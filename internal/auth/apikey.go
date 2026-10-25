package auth

import (
	"log"

	"github.com/xts-protocol/xtsd/internal/tenant"
)

func ValidateAPIKey(key string) bool {

	_, ok := tenant.Get(key)

	log.Println(
		"apikey",
		key,
		"valid",
		ok,
	)

	return ok
}
