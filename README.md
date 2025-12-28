# Pricefetcher Microservice

A professional, layered Go microservice for fetching cryptocurrency prices. This project demonstrates the usage of the **Decorator Pattern**, **Interface-driven design**, and **Request ID tracking** in a clean architecture.

## Architecture

The project is built with a focus on maintainability and extensibility using layers:

1.  **Service Layer (`service.go`)**: Contains the core logic and interfaces. It uses a mock price fetcher with support for BTC, ETH, and GG.
2.  **Logging Decorator (`logging.go`)**: Wraps the service to provide JSON logging (via `logrus`), request timing, and request ID tracking.
3.  **Metrics Decorator (`metrics.go`)**: An extensible layer for pushing metrics to monitoring systems like Prometheus.
4.  **Transport Layer (`api.go`)**: Implements a standard JSON HTTP API.
5.  **Client package (`client/`)**: A reusable Go client for interacting with the service.

## Features

- **Decorator Pattern**: Seamlessly add logging and metrics without modifying core business logic.
- **Request ID Tracking**: Context-aware logging that includes unique request IDs.
- **Mock Data**: Pre-configured prices for testing:
  - BTC: 20,000.0
  - ETH: 200.0
  - GG: 100,000.0
- **Docker Support**: Ready for containerized deployment.
- **Makefile**: Simplified build and run commands.

## Getting Started

### Prerequisites

- Go 1.25 or higher
- Docker (optional)

### Running Locally

1.  Clone the repository and navigate to the project directory.
2.  Build and run using the Makefile:
    ```bash
    make run
    ```
    Alternatively, run directly:
    ```bash
    go run . --listenAddr :3000
    ```

### Running with Docker

1.  Build the image:
    ```bash
    docker build -t pricefetcher .
    ```
2.  Run the container:
    ```bash
    docker run -p 3000:3000 pricefetcher
    ```

## API Usage

### Fetch Price

**Endpoint**: `GET /?ticker={ticker}`

**Example**:
```bash
curl "http://localhost:3000/?ticker=BTC"
```

**Response**:
```json
{"ticker":"BTC","price":20000}
```

## Using the Go Client

You can use the built-in client package in your other Go projects:

```go
import "github.com/szoumoc/client"

func main() {
    c := client.New("http://localhost:3000/")
    price, err := c.FetchPrice(context.Background(), "ETH")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Price: %+v\n", price)
}
```
