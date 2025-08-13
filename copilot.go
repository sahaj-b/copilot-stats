package main

import (
	"log"
)

func main() {
	oauthToken, err := getOAuthToken()
	if err != nil {
		log.Fatalf("Failed to get OAuth token: %v", err)
	}

	stats, err := getCopilotStats(oauthToken)
	if err != nil {
		log.Fatalf("Failed to get Copilot stats: %v", err)
	}

	displayCopilotStats(stats)
}
