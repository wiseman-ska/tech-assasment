package commons

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
)

const configPath = "user-manager-api/etc/config.json"

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
		Server      string `json:"server"`
		MongoDBHost string `json:"mongoDBHost"`
		DBUser      string `json:"dbUser"`
		DBPwd       string `json:"pwd"`
		Database    string `json:"database"`
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

func initConfig() {
	loadAppConfig()
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

func IsValidSAIdNumber(idNum string) bool {
	if len(idNum) != 13 {
		return false
	}
	number, _ := strconv.Atoi(idNum)
	if (number%10+checksum(number/10))%10 != 0 {
		return false
	}
	return true
}

func checksum(number int) int {
	var luhn int
	for i := 0; number > 0; i++ {
		cur := number % 10
		if i%2 == 0 {
			cur = cur * 2
			if cur > 9 {
				cur = cur%10 + cur/10
			}
		}
		luhn += cur
		number = number / 10
	}
	return luhn % 10
}

func JSONMarshal(v interface{}, safeEncoding bool) ([]byte, error) {
	b, err := json.Marshal(v)
	if safeEncoding {
		b = bytes.Replace(b, []byte("\\u003c"), []byte(""), -1)
		b = bytes.Replace(b, []byte("\\u003e"), []byte(""), -1)
		b = bytes.Replace(b, []byte("\\u0026"), []byte(""), -1)
	}
	return b, err
}
