package commons

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

const configPath = "../etc/config.json"

type (
	AppError struct {
		Error      string `json:"error"`
		Message    string `json:"message"`
		HttpStatus int    `json:"status"`
	}
	ErrorResource struct {
		Data AppError `json:"data"`
	}
	AppConfig struct {
		Server      string
		MongoDBHost string
		DBUser      string
		DBPwd       string
		Database    string
	}
)
var AppConf *AppConfig

func DisplayAppError(w http.ResponseWriter, handleError error, message string, code int)  {
	appError := AppError{
		Error: handleError.Error(),
		Message: message,
		HttpStatus: code,
	}
	log.Printf("[AppErr]: %s\n", handleError)
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if resp, err := json.Marshal(ErrorResource{Data: appError}); err == nil {
		w.Write(resp)
	}
}

func loadAppConfig() {
	file, err := os.Open(configPath)
	defer file.Close()
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&AppConf)
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}
}

func initConfig() {
	loadAppConfig()
}



