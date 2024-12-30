package utils

import (
	"fmt"
	"github.com/SerhiiCho/timeago"
	"time"
)

func init() {
	// Set the locale to English
	timeago.SetConfig(
		timeago.Config{
			Location: "Asia/Jakarta",
		})
}

// TimeAgo function converts a timestamp into a human-readable "time ago" format
func TimeAgo(timestampStr string) (string, error) {
	// Parse the timestamp string
	timestamp, err := time.Parse(time.RFC3339, timestampStr)
	if err != nil {
		return "", fmt.Errorf("error parsing timestamp: %v", err)
	}

	// Initialize the timeago instance
	ta := timeago.Parse(timestamp)

	// Return the formatted "time ago" string
	return ta, nil
}
