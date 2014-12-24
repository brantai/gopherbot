package main

import "log"

var (
	configMap = make(map[string]string)
	Plugins   = make([]Plugin, 0)
	debugLog  *log.Logger
)

const (
	DATA_DIR string = "data"
	CFG_FILE string = "gopher.cfg"
	LOG_FILE string = "gopher.log"
)
