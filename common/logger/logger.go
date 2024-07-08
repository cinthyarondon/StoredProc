package logger

import (
	"log"
	"os"
)

var Log = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)