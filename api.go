package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func getOAuthToken() (string, error) {
	// Check environment variables first (for GitHub Codespaces)
	if token := os.Getenv("GITHUB_TOKEN"); token != "" && os.Getenv("CODESPACES") != "" {
		return token, nil
	}

	// Find config path
	configPath := os.Getenv("XDG_CONFIG_HOME")
	if configPath == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		configPath = filepath.Join(homeDir, ".config")
	}

	// Try both hosts.json and apps.json
	filePaths := []string{
		filepath.Join(configPath, "github-copilot", "hosts.json"),
		filepath.Join(configPath, "github-copilot", "apps.json"),
	}

	for _, filePath := range filePaths {
		if _, err := os.Stat(filePath); err == nil {
			data, err := os.ReadFile(filePath)
			if err != nil {
				continue
			}

			var config map[string]any
			if err := json.Unmarshal(data, &config); err != nil {
				continue
			}

			for key, value := range config {
				if strings.Contains(key, "github.com") {
					if valueMap, ok := value.(map[string]interface{}); ok {
						if token, ok := valueMap["oauth_token"].(string); ok {
							return token, nil
						}
					}
				}
			}
		}
	}

	return "", fmt.Errorf("no OAuth token found")
}

func getCopilotStats(oauthToken string) (*CopilotStats, error) {
	req, err := http.NewRequest("GET", "https://api.github.com/copilot_internal/user", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+oauthToken)
	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-Agent", "CodeCompanion.nvim")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var stats CopilotStats
	if err := json.Unmarshal(body, &stats); err != nil {
		return nil, err
	}

	return &stats, nil
}
