package auth

func ValidateAPIKey(key string) bool {
	return key == "xts_demo_key"
}
