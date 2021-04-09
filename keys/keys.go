package keys

import (
	"os"

	"github.com/joho/godotenv"
)

var MONGO_URL string
var PORT string

func Load() {
	godotenv.Load(".env")

	PORT = os.Getenv("PORT")
	MONGO_URL = os.Getenv("MONGO_URL")
}
