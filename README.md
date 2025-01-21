# Caching Proxy Server

A lightweight and efficient caching proxy server built with Go. This server forwards requests to a specified origin server, caches the responses, and serves subsequent identical requests from the cache. It includes a CLI for starting the server and clearing the cache.

## Features

- **Proxy Requests**: Forwards HTTP requests to an origin server.
- **Response Caching**: Caches responses to identical requests for faster subsequent retrievals.
- **Custom Headers**:
  - `X-Cache: HIT` if the response is served from the cache.
  - `X-Cache: MISS` if the response is forwarded from the origin server.
- **CLI Options**:
  - Start the server with custom port and origin.
  - Clear the cache directly from the CLI.
- **Concurrency**: Supports multiple simultaneous requests with thread-safe caching.

---

## Getting Started

### Prerequisites

- Go (1.19 or later)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/caching-proxy.git
   cd caching-proxy
   ```

2. Build the project:

   ```bash
   go build -o caching-proxy ./cmd/caching-proxy
   ```

3. Verify the build:
   ```bash
   ./caching-proxy --help
   ```

## Usage

### Start the Proxy Server

Run the server with a custom port and origin URL:

```bash
./caching-proxy --port <number> --origin <url>
```

**Example:**
Start the server on port 3000 and forward requests to http://dummyjson.com:

```bash
./caching-proxy --port 3000 --origin http://dummyjson.com
```

**Behavior:**

- A request to http://localhost:3000/products will forward to http://dummyjson.com/products and cache the response.

- Subsequent requests to http://localhost:3000/products will return the cached response.

### Clear the Cache

Clear all cached responses without starting the server:

```bash
./caching-proxy --clear-cache
```

## Project Structure

```csharp
caching-proxy/
├── cmd/
│   └── caching-proxy/
│       └── main.go          # CLI entry point
├── internal/
│   ├── cache/
│   │   └── cache.go         # Cache implementation
│   ├── proxy/
│   │   └── server.go        # Proxy server logic
├── go.mod                   # Go module file
├── README.md                # Project documentation
```

## How It Works

1. **Request Handling**:

   - The server listens on the specified port and intercepts HTTP requests.
   - For each request:
     - Cache Hit: Return the cached response with `X-Cache: HIT`.
     - Cache Miss: Forward the request to the origin server, cache the response, and return it with `X-Cache: MISS`.

2. **Cache Storage**:

   - Responses are cached in-memory using a thread-safe mechanism (`sync.Map`).
   - The cache key is constructed using the HTTP method and request URL.

3. **Cache Management**:
   - A `ClearCache` function is provided to reset the cache.

## Future Enhancements

- **Time-to-Live (TTL)**: Add support for cache expiration.
- **Persistent Cache**: Save cache to disk for persistence across restarts.
- **Configurable Options**: Allow users to configure cache size and TTL.
- **HTTPS Support**: Add support for forwarding secure HTTPS requests.

## Contributing

Contributions are welcome! Feel free to submit issues or pull requests to improve the project.
