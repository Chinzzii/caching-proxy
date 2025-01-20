package proxy

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"caching-proxy/internal/cache"
)

// Initialize a global cache instance to store responses.
var cacheStore = cache.NewCache()

// StartServer starts the caching proxy server on the specified port and forwards requests to the origin server.
func StartServer(port int, origin string) {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Construct a unique cache key using the HTTP method and request URL.
		key := r.Method + ":" + r.URL.Path + "?" + r.URL.RawQuery

		// Check if the request is already in the cache.
		value, found := cacheStore.Get(key)
		if found {
			w.Header().Set("X-Cache", "HIT")
			w.Write([]byte(value))
			return
		}

		// If the response is not cached, forward the request to the origin server.
		res, err := http.Get(origin + r.URL.Path + "?" + r.URL.RawQuery)
		if err != nil {
			// If an error occurs while forwarding the request, return a 502 Bad Gateway error.
			http.Error(w, "Error forwarding request", http.StatusBadGateway)
			return
		}
		defer res.Body.Close() // Ensure the response body is closed after reading.

		// Read the response body from the origin server.
		body, _ := io.ReadAll(res.Body)

		// Cache the response body using the constructed key.
		cacheStore.Set(key, string(body))

		// Set the X-Cache header to "MISS" to indicate the response is from the origin server.
		w.Header().Set("X-Cache", "MISS")

		// Write the response body to the client.
		w.Write(body)
	})

	// Start the HTTP server on the specified port.
	log.Printf("Caching proxy server is running on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

// ClearCache clears all entries from the cache.
func ClearCache() {
	cacheStore.Clear()
}
