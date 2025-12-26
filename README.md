# Microservice Price Fetcher

A simple, robust microservice written in Go that simulates fetching cryptocurrency prices. This project demonstrates basic microservice patterns, including interface-based design and decorator patterns for cross-cutting concerns like logging.

## Features

- **Price Fetching**: Mock implementation for fetching prices of popular cryptocurrencies (e.g., BTC, ETH).
- **Logging Middleware**: Uses the decorator pattern to wrap the service with structured logging using `logrus`.
- **Easy Build & Run**: Includes a `Makefile` for streamlined development.

## Prerequisites

- [Go](https://golang.org/doc/install) (1.18+ recommended)
- `make` utility

## Getting Started

### Installation

1. Clone the repository (if applicable) or navigate to the project directory.
2. Download dependencies:
   ```bash
   go mod download
   ```

### Building the Project

To build the binary:
```bash
make build
```
This will create a binary named `pricefetcher` in the `bin/` directory.

### Running the Project

To build and run the service in one command:
```bash
make run
```

## Project Structure

- `main.go`: Entry point of the application.
- `service.go`: Contains the `Pricefetcher` interface and its implementation.
- `logging.go`: Implements the logging middleware for the service.
- `Makefile`: Contains build and run scripts.
- `go.mod` / `go.sum`: Go module dependency files.

## Example Output

When running the service, you should see output similar to this:
```text
INFO[0000] fetchPrice                                    err="<nil>" price=3000 took="18.5µs"
3000
```
