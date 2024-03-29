package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var Env Environment

type Environment struct {
	BotToken           string
	GeminiApiKey       string
	AllowedUsers       []string
	DebugFlag          bool
	AWSAccessKey       string
	AWSSecretAccessKey string
	RedisAddr          string
	RedisPassword      string
	RedisDB            int
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, trying to load from environment")
	}

	allowedUsersVar := getEnv("ALLOWED_USERS", "")
	var allowedUsers []string
	if allowedUsersVar != "" {
		allowedUsers = strings.Split(allowedUsersVar, ",")
	}

	Env = Environment{
		BotToken:           getEnv("BOT_TOKEN", ""),
		GeminiApiKey:       getEnv("GEMINI_API_KEY", ""),
		AllowedUsers:       allowedUsers,
		DebugFlag:          getEnv("BOT_DEBUG_MODE", "false") == "true",
		AWSAccessKey:       getEnv("AWS_ACCESS_KEY_ID", ""),
		AWSSecretAccessKey: getEnv("AWS_SECRET_ACCESS_KEY", ""),
	}

	if Env.BotToken == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN must be set in environment variables or .env file")
	}
	if Env.GeminiApiKey == "" {
		log.Fatal("GEMINI_API_KEY must be set in environment variables or .env file")
	}
	if Env.AWSAccessKey == "" {
		log.Fatal("AWS_ACCESS_KEY must be set in environment variables or .env file")
	}
	if Env.AWSSecretAccessKey == "" {
		log.Fatal("AWS_SECRET_ACCESS_KEY must be set in environment variables or .env file")
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func init() {
	loadEnv()
	log.Printf("Loaded env")
}
