package main

import "log"

const (
	logErr   = iota
	logWarn  = iota
	logInfo  = iota
	logDebug = iota
)

func dLog(format string, v ...interface{}) {
	log.Printf(format, v)
}

func lg(level int, format string, v ...interface{}) {
	if logLevel >= logDebug {
		log.Printf(format, v)
	} else if logLevel >= logInfo {
		log.Printf(format, v)
	} else if logLevel >= logWarn {
		log.Printf(format, v)
	} else if logLevel >= logErr {
		log.Printf(format, v)
	}
}
