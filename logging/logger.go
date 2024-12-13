package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

// CreateLogFile is a func that creates a new file for logging,
// naming the file based on a formatted timestamp.
// formats the top of the file describing what kinds of tests have been run in this log.
func CreateLogFile() (*os.File, error) {
	// get current time,then format it into a readable file name.
	timeStamp := time.Now()
	formattedDate := timeStamp.Format("2006-01-02_15:04:05")

	// create a file, using formatted date as the name for logging.
	logFile, err := os.Create(formattedDate + ".log")
	if err != nil {
		return nil, fmt.Errorf("ERR in creating file with filename: %s... %w\n", formattedDate, err)
	}
	// TODO: append test suite details to the top of the CreateLogFile

	return logFile, nil
}

func LogToFile(file *os.File, msg string) {
	log.Println(msg)
	if _, err := file.WriteString(msg + "\n"); err != nil {
		log.Fatalf("failed to write to log file: %v\n", err)
	}
}
