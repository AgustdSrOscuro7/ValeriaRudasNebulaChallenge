package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const baseURL = "https://api.ssllabs.com/api/v2/analyze"

// HostResponse represents the main response from SSL Labs
type HostResponse struct {
	Host          string     `json:"host"`
	Status        string     `json:"status"`
	StatusMessage string     `json:"statusMessage"`
	Endpoints     []Endpoint `json:"endpoints"`
}

// Endpoint represents a single endpoint result
type Endpoint struct {
	IPAddress     string `json:"ipAddress"`
	StatusMessage string `json:"statusMessage"`
	Grade         string `json:"grade"`
	Progress      int    `json:"progress"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <domain>")
		os.Exit(1)
	}

	domain := os.Args[1]
	fmt.Println("Starting SSL Labs analysis for:", domain)

	// Start a new assessment (we don't need the response here)
	_, err := analyze(domain, true)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Poll until READY or ERROR
	for {
		time.Sleep(10 * time.Second)

		response, err := analyze(domain, false)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		fmt.Println("Current status:", response.Status)

		if response.Status == "READY" {
			printResults(response)
			break
		}

		if response.Status == "ERROR" {
			fmt.Println("Analysis failed:", response.StatusMessage)
			break
		}
	}
}

// analyze calls the SSL Labs analyze endpoint
func analyze(domain string, startNew bool) (*HostResponse, error) {
	url := fmt.Sprintf("%s?host=%s&publish=off", baseURL, domain)

	if startNew {
		url += "&startNew=on&all=done"
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result HostResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// printResults prints a summary of the TLS results
func printResults(result *HostResponse) {
	fmt.Println("\nTLS Analysis completed for:", result.Host)
	fmt.Println("--------------------------------------------------")

	if len(result.Endpoints) == 0 {
		fmt.Println("No endpoints found.")
		return
	}

	for _, ep := range result.Endpoints {
		fmt.Println("IP Address:", ep.IPAddress)
		fmt.Println("Status:", ep.StatusMessage)

		if ep.Grade != "" {
			fmt.Println("Grade:", ep.Grade)
		} else {
			fmt.Println("Grade: Not available")
		}

		fmt.Println("--------------------------------------------------")
	}
}
