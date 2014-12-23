package main

import (
	"encoding/csv"
	"os"
)

var configMap map[string]string

func loadConfig() error {

	configMap = make(map[string]string)

	file, err := os.Open("gopher.cfg")
	if err != nil {
		debugLog.Printf("Could not open config file: %v", err)
		return err
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = '='
	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		debugLog.Printf("Could not read config file: %v", err)
		return err
	}

	for _, v := range rawCSVdata {
		configMap[v[0]] = v[1]
	}

	return nil
}
