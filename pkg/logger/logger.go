package logger

import (
	"log"
	"os"
)

var Logger = log.New(os.Stdout, "[workerpool] ", log.LstdFlags|log.Lshortfile)
