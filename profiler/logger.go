package profiler

import "log"

func debugLog(format string, e ...interface{}) {
	log.Printf(format, e...)
}