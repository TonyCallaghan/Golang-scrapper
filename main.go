package main

import (
	"github.com/tonyc/scrape/core"
	"log"
)

func main() {
	err := core.Execute()
	if err != nil {
		log.Fatalf("Failed to execute core: %v", err)
	}
}
