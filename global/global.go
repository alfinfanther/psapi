package global

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		//panic(err)
		code = http.StatusInternalServerError
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func GetEnv(key string) string {

	if len(os.Args) > 1 && os.Args[1] == "local" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file GAAN")
		}
	}
	return os.Getenv(key)
}
