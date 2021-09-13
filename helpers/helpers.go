package helpers

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func GetUriMQ() string {
	err := godotenv.Load(".env")
	FailOnError(err, "Failed to lod .env in current directory")
	return os.Getenv("RABBITMQ_ENDPOINT")
}
