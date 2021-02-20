package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {
	endpoint := "http://localhost:8080/health"
	resp, err := http.Get(endpoint)
	if err != nil {
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Validate that all of the healthPoints are where they should be -- this
	// one is the simplest possible.
	var healthchecks struct {
		IsAlive string
	}

	err = json.NewDecoder(resp.Body).Decode(&healthchecks)
	if err != nil {
		log.Printf("Error decoding body: %v", err)
		os.Exit(1)
	}

	if healthchecks.IsAlive != "true" {
		log.Printf("IsAlive returned '%s'", healthchecks.IsAlive)
		os.Exit(1)
	}
	// Implicit exit with code 0
}
