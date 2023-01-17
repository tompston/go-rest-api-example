package settings

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// load all of the env variables from .env
func LoadEnvFile() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		fmt.Println("Error loading .env file")
	}
}

// load a specified .env variable
func Config(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}
