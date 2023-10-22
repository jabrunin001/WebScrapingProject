package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

// Wikipedia URLs for topics of interest
var urls = []string{
	"https://en.wikipedia.org/wiki/Robotics",
	"https://en.wikipedia.org/wiki/Robot",
	"https://en.wikipedia.org/wiki/Reinforcement_learning",
	"https://en.wikipedia.org/wiki/Robot_Operating_System",
	"https://en.wikipedia.org/wiki/Intelligent_agent",
	"https://en.wikipedia.org/wiki/Software_agent",
	"https://en.wikipedia.org/wiki/Robotic_process_automation",
	"https://en.wikipedia.org/wiki/Chatbot",
	"https://en.wikipedia.org/wiki/Applications_of_artificial_intelligence",
	"https://en.wikipedia.org/wiki/Android_(robot)",
}

func main() {
	// Initialize Colly collector
	c := colly.NewCollector()

	// Set up error handling
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Set up scraping logic
	c.OnHTML("div.mw-parser-output", func(e *colly.HTMLElement) {
		// Extract text content from the div and clean it up
		text := strings.TrimSpace(e.Text)
		if text == "" {
			return
		}

		// Create a map to store the data
		data := map[string]interface{}{
			"url":  e.Request.URL.String(),
			"text": text,
		}

		// Convert the map to JSON
		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Failed to convert data to JSON:", err)
			return
		}

		// Print the JSON data
		fmt.Println(string(jsonData))
	})

	// Set up concurrency
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 10})

	// Start the timer
	start := time.Now()

	// Start the crawling process
	for _, url := range urls {
		c.Visit(url)
	}

	// Wait until the crawling is finished
	c.Wait()

	// Stop the timer and print the elapsed time
	elapsed := time.Since(start)
	fmt.Printf("Crawling took %s\n", elapsed)
}
