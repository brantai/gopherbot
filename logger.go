package main

import (
	"fmt"
	"log"
	"os"
)

func InitLogger() {

	file, err := os.OpenFile(DATA_DIR+"gopher.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Cannot open log file. ERR: %v", err)
	}

	debugLog = log.New(file, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
}
