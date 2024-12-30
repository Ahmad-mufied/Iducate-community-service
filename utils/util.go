package utils

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

// StringToUint converts a string to a uint.
// If the conversion fails, it returns 0.
func StringToUint(s string) uint {
	val, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		// Return 0 if there's an error
		return 0
	}
	return uint(val)
}

// ParsePostgresTimestamp parses a PostgreSQL timestamp string in ISO 8601 format
// (e.g., '2024-12-17T19:00:11+07:00') into a Go time.Time object.
func ParsePostgresTimestamp(timestamp string) (time.Time, error) {
	// Define the format used in the PostgreSQL timestamp string with timezone offset
	const timeFormatWithOffset = "2006-01-02T15:04:05-07:00"
	const timeFormatWithUTC = "2006-01-02T15:04:05Z" // Special format for UTC 'Z'

	// Handle the 'Z' character by replacing it with '+00:00' to treat it as UTC
	if timestamp[len(timestamp)-1] == 'Z' {
		timestamp = timestamp[:len(timestamp)-1] + "+00:00"
	}

	// Try parsing the timestamp with offset format
	parsedTime, err := time.Parse(timeFormatWithOffset, timestamp)
	if err != nil {
		log.Printf("Error parsing timestamp '%s': %v", timestamp, err)
		return time.Time{}, fmt.Errorf("unable to parse timestamp: %v", err)
	}

	return parsedTime, nil
}
