package utils

import (
	"fmt"
	"log"
	"os"
	"reflect"
)

// NOTE: Utility function for checking if a struct is empty
// x: interface => struct for empty check
// returns: bool => true or false based on result
func IsEmptyStruct(x interface{}) bool {
	v := reflect.ValueOf(x)

	for i := 0; i < v.NumField(); i++ {
		if !v.Field(i).IsZero() {
			return false
		}
	}

	return true
}

// NOTE: Utility function for database logging
func LogDB(msg string) {
	file, err := os.OpenFile("../database/db.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		msg := fmt.Sprintf("Error opening db.log file: %v", err)
		panic(msg)
	}

	defer file.Close()

	log.SetOutput(file)
	log.SetFlags(log.Ldate | log.Ltime)

	log.Println(msg)
}
