package config

import (
	"log"
	"os"
	"strings"
)

// ReadEnv reads the environment variable from the system or the .env file. When the environment variable is not found, the application will exit.
func ReadEnv(key string, defaultValue string) string {
	env := os.Getenv(key)
	if env != "" {
		return env
	}

	file, err := os.ReadFile(".env")
	if err != nil {
		if defaultValue != "" {
			return defaultValue
		}
		log.Fatal("Error reading the .env file")
	}

	fileContent := string(file)
	lines := strings.Split(fileContent, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, key) {
			returnValue := line[strings.Index(line, "=")+1:]
			// remove line breaks
			returnValue = strings.ReplaceAll(returnValue, "\n", "")
			returnValue = strings.ReplaceAll(returnValue, "\r", "")
			// And quotes
			returnValue = strings.ReplaceAll(returnValue, "\"", "")
			return returnValue
		}
	}

	if defaultValue != "" {
		return defaultValue
	}

	log.Fatal("Environment variable not found in .env or system environment variables (key: " + key + ")")
	return ""
}
