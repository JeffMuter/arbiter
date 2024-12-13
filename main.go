package main

import (
	"fmt"
	"log"

	"github.com/JeffMuter/arbiter/logging"
)

func main() {
	// tests are being run, let's create a log file
	logFile, err := logging.CreateLogFile()
	if err != nil {
		fmt.Printf("error in creating log file: %v\n", err)
	}
	defer logFile.Close()

	log.SetOutput(logFile) // sets the output of logging functions to the logfile.

}
