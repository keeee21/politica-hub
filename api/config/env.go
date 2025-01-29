package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv 環境変数をロード
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found, loading from system env")
	}
}

// GetEnv 環境変数を取得
func GetEnv(key string) string {
	return os.Getenv(key)
}