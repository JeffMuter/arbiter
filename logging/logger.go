package logging

import (
	"log"
	"os"
)

func LogToFile(file *os.File, msg string) {
	log.Println(msg)
	if _, err := file.WriteString(msg + "\n"); err != nil {
		log.Fatalf("failed to write to log file: %v\n", err)
	}
}
