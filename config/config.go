package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringConnectionBase = ""
	APIPort              = ""

	// chave que vai ser usada para assinar o token
	SecretKey []byte
)

func Load() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	port, err := strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		port = 8088
	}
	APIPort = fmt.Sprintf(":%d", port)

	StringConnectionBase = fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
