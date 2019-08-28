package drone

import (
	"log"
)

// LogError prints an error message to the command line
func LogError(reason string, err error) {
	log.Println("Drone error: ", reason)
	if err != nil {
		log.Println("  Cause:", err)
	}
}
