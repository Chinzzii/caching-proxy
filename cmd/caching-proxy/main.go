package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Chinzzii/caching-proxy/internal/proxy"
)

func main() {
	// Define command-line flags for port, origin, and clearing the cache.
	PORT := flag.Int("port", 3000, "Port to run the proxy server")
	ORIGIN := flag.String("origin", "", "Origin server URL")
	CLEAR_CACHE := flag.Bool("clear-cache", false, "Clear the cache and exit")

	// Parse the command-line flags provided by the user.
	flag.Parse()

	// Check if the --clear-cache flag is set. If so, clear the cache and exit.
	if *CLEAR_CACHE {
		proxy.ClearCache()
		fmt.Println("Cache cleared successfully.")
		return
	}

	// Validate the --origin flag to ensure it's provided.
	if *ORIGIN == "" {
		log.Fatal("Error: Origin server URL --origin is required when starting the server")
	}

	// Log the starting of the proxy server with the provided port and origin.
	log.Printf("Starting caching proxy server on port %d forwarding to %s", *PORT, *ORIGIN)

	// Start the proxy server using the port and origin provided by the user.
	proxy.StartServer(*PORT, *ORIGIN)
}
