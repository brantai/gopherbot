package main

import (
	"encoding/csv"
	"os"
)

func loadConfig() error {

	file, err := os.Open(configFile)
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
