package configs

import (
	"log"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//return os.Getenv("MONGOURI")
	return "mongodb+srv://abdullahb:kMj8978zz6EMShCE@cards.ztgcu.mongodb.net/beyazhoroz?retryWrites=true&w=majority"
}
