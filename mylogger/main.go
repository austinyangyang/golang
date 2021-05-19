package main

import (
	"mylogger/logger"
)

func main() {
	// log := logger.NewLog("TRACE")
	log := logger.NewFileLogger("INFO", "./", "cyy.log", 1024)
	id := 1010
	name := "access"
	for {
		// log.Trace("this is a Info log")
		log.Error("this is a Info log, id: %s, name: %d", id, name)
		// log.Warring("this is a Info log, id: %s, name: %d", id, name)
		// log.Error("this is a Info log, id: %s, name: %d", id, name)
		// time.Sleep(time.Second)

	}

}
