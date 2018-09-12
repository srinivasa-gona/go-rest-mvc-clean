package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"reflect"
	"rest-api/models"
	"strconv"
)

func GetConfiguration() models.ConfigProps {

	config := models.ConfigProps{}
	file, err := os.Open("./configuration.json")

	if err != nil {
		log.Panicln("Error in reading file : ", err)
	}

	defer file.Close()
	decoder := json.NewDecoder(file)
	parseErr := decoder.Decode(&config)

	if parseErr != nil {
		log.Panicln("Error in parsing configuration file ", parseErr)
	}

	val := reflect.ValueOf(&config).Elem()

	for i := 0; i < val.NumField(); i++ {
		filedName := val.Type().Field(i).Name
		if os.Getenv(filedName) != "" {

			switch val.Type().Field(i).Type.Kind() {
			case reflect.String:
				val.Field(i).SetString(os.Getenv(filedName))
			case reflect.Int:
				intVal, err := strconv.ParseInt(os.Getenv(filedName), 10, 64)
				if err == nil {
					val.Field(i).SetInt(intVal)
				}
			case reflect.Bool:
				b, err := strconv.ParseBool(os.Getenv(filedName))
				if err == nil {
					val.Field(i).SetBool(b)
				}
			}
		}
	}
	return config

}

func RestError(message string, code int) *models.RestError {
	return &models.RestError{
		Err:     message,
		ErrCode: code,
	}
}

func RestRespond(w http.ResponseWriter, data interface{}, err error) {

	w.Header().Add("Content-Type", "application/json")
	if err != nil {
		if err, ok := err.(*models.RestError); ok {
			w.WriteHeader(err.ErrCode)
		}
		json.NewEncoder(w).Encode(err)
		return
	}

	json.NewEncoder(w).Encode(data)

}
