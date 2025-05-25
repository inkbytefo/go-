# GO-Minus HTTP Package

This package provides HTTP client and server implementations for the GO-Minus programming language. It allows you to create HTTP servers, make HTTP requests, and handle HTTP responses.

## Features

- HTTP client for making requests
- HTTP server for handling requests
- Support for HTTP/1.1 and HTTP/2
- Cookie handling
- Form data processing
- File uploads and downloads
- WebSocket support
- Middleware support
- TLS/SSL support
- Request and response compression

## Usage

### HTTP Client

```go
import (
    "fmt"
    "net/http"
)

func main() {
    // Create a new HTTP client
    client := http.Client.New()
    
    // Make a GET request
    response, err := client.Get("https://example.com")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    
    // Print the response status code
    fmt.Println("Status:", response.StatusCode)
    
    // Read the response body
    body, err := response.ReadBody()
    if err != nil {
        fmt.Println("Error reading body:", err)
        return
    }
    
    // Print the response body
    fmt.Println("Body:", string(body))
    
    // Close the response body
    response.Close()
}
```

### HTTP Server

```go
import (
    "fmt"
    "net/http"
)

// Handler function for the root path
func handleRoot(w http.ResponseWriter, r http.Request) {
    w.WriteHeader(200)
    w.Write("Hello, World!")
}

// Handler function for the /api path
func handleAPI(w http.ResponseWriter, r http.Request) {
    w.WriteHeader(200)
    w.Write(`{"message": "Welcome to the API"}`)
}

func main() {
    // Create a new HTTP server
    server := http.Server.New()
    
    // Register handlers
    server.HandleFunc("/", handleRoot)
    server.HandleFunc("/api", handleAPI)
    
    // Start the server on port 8080
    fmt.Println("Server starting on :8080")
    err := server.ListenAndServe(":8080")
    if err != nil {
        fmt.Println("Server error:", err)
    }
}
```

### Using Middleware

```go
import (
    "fmt"
    "net/http"
    "time"
)

// Logging middleware
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r http.Request) {
        start := time.Now()
        fmt.Printf("Request: %s %s\n", r.Method, r.URL.Path)
        
        // Call the next handler
        next(w, r)
        
        // Log the time taken
        fmt.Printf("Request completed in %v\n", time.Since(start))
    }
}

// Handler function
func handleRoot(w http.ResponseWriter, r http.Request) {
    w.WriteHeader(200)
    w.Write("Hello, World!")
}

func main() {
    // Create a new HTTP server
    server := http.Server.New()
    
    // Register handler with middleware
    server.HandleFunc("/", loggingMiddleware(handleRoot))
    
    // Start the server on port 8080
    fmt.Println("Server starting on :8080")
    err := server.ListenAndServe(":8080")
    if err != nil {
        fmt.Println("Server error:", err)
    }
}
```

## Classes and Interfaces

### Client

The `Client` class is used to make HTTP requests.

```go
class Client {
    // Create a new HTTP client
    static func New() *Client
    
    // Make a GET request
    func Get(url string) (*Response, error)
    
    // Make a POST request
    func Post(url string, contentType string, body io.Reader) (*Response, error)
    
    // Make a PUT request
    func Put(url string, contentType string, body io.Reader) (*Response, error)
    
    // Make a DELETE request
    func Delete(url string) (*Response, error)
    
    // Make a HEAD request
    func Head(url string) (*Response, error)
    
    // Make a custom request
    func Do(req *Request) (*Response, error)
}
```

### Server

The `Server` class is used to create HTTP servers.

```go
class Server {
    // Create a new HTTP server
    static func New() *Server
    
    // Register a handler function for a path
    func HandleFunc(pattern string, handler HandlerFunc)
    
    // Register a handler for a path
    func Handle(pattern string, handler Handler)
    
    // Start the server on the specified address
    func ListenAndServe(addr string) error
    
    // Start the server with TLS on the specified address
    func ListenAndServeTLS(addr, certFile, keyFile string) error
    
    // Shutdown the server gracefully
    func Shutdown() error
}
```

### Request

The `Request` class represents an HTTP request.

```go
class Request {
    // Create a new HTTP request
    static func New(method, url string, body io.Reader) (*Request, error)
    
    // Get a request header
    func Header(key string) string
    
    // Set a request header
    func SetHeader(key, value string)
    
    // Get a URL parameter
    func Param(key string) string
    
    // Get a form value
    func FormValue(key string) string
    
    // Get a file from a multipart form
    func FormFile(key string) (*FormFile, error)
    
    // Parse the request body as a form
    func ParseForm() error
    
    // Parse the request body as a multipart form
    func ParseMultipartForm(maxMemory int64) error
}
```

### Response

The `Response` class represents an HTTP response.

```go
class Response {
    // Status code of the response
    StatusCode int
    
    // Status message of the response
    Status string
    
    // Get a response header
    func Header(key string) string
    
    // Set a response header
    func SetHeader(key, value string)
    
    // Read the response body
    func ReadBody() ([]byte, error)
    
    // Close the response body
    func Close() error
}
```

## Error Handling

The HTTP package uses GO-Minus's exception handling mechanism for error handling.

```go
import (
    "fmt"
    "net/http"
)

func main() {
    try {
        client := http.Client.New()
        response, err := client.Get("https://example.com")
        if err != nil {
            throw err
        }
        
        body, err := response.ReadBody()
        if err != nil {
            throw err
        }
        
        fmt.Println("Body:", string(body))
        response.Close()
    } catch (err) {
        fmt.Println("Error:", err)
    }
}
```

## WebSocket Support

The HTTP package includes WebSocket support for real-time communication.

```go
import (
    "fmt"
    "net/http"
    "net/http/websocket"
)

func handleWebSocket(ws *websocket.Conn) {
    for {
        // Read a message
        message, err := ws.ReadMessage()
        if err != nil {
            fmt.Println("Error reading message:", err)
            break
        }
        
        fmt.Println("Received message:", string(message))
        
        // Send a response
        err = ws.WriteMessage([]byte("Received: " + string(message)))
        if err != nil {
            fmt.Println("Error writing message:", err)
            break
        }
    }
}

func main() {
    server := http.Server.New()
    
    // Register WebSocket handler
    server.HandleWebSocket("/ws", handleWebSocket)
    
    fmt.Println("Server starting on :8080")
    err := server.ListenAndServe(":8080")
    if err != nil {
        fmt.Println("Server error:", err)
    }
}
```
