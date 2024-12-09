package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/JeffMuter/arbiter/logging"
)

func main() {
	// create a file to log in, by date and time in logs package.
	startTime := time.Now()
	timestamp := startTime.Format("02-01-15:04:05") // Use Go's reference time format
	logfile, err := os.Create("logs/test_" + timestamp)
	if err != nil {
		log.Fatalf("creating log file err: %w\n", err)
	}
	defer logfile.Close()

	err = runTests(logfile)
	if err != nil {
		logging.LogToFile(logfile, fmt.Sprintf("test suite had error while running: %w\n", err))
	}

	endTime := time.Now()
	durationOfTests := endTime.Sub(startTime)

	logging.LogToFile(logfile, fmt.Sprintf("tests finished after %v seconds\n", durationOfTests))
}

func runTests(logfile *os.File) error {

	return nil
}
