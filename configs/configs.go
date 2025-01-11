package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println(err)
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
