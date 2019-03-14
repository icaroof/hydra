package hlogger

import (
	"log"
	"os"
	"sync"
)

// HydraLogger structure
type HydraLogger struct {
	*log.Logger
	filename string
}

var hlogger *HydraLogger
var once sync.Once

// GetInstance gets a singleton instance of the hydraLogger
func GetInstance() *HydraLogger {
	once.Do(func() {
		hlogger = createLogger("./Hydra/hydralogger.log")
	})

	return hlogger
}

func createLogger(fname string) *HydraLogger {
	file, _ := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)

	return &HydraLogger{
		filename: fname,
		Logger:   log.New(file, "Hydra ", log.Lshortfile),
	}
}
