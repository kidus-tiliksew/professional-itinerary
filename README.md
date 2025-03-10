# Professional Itinerary

A Go-based API service that constructs travel itineraries from a collection of tickets.

## Features

- Constructs a valid itinerary from unordered tickets
- Detects and reports various error conditions:
  - Disconnected paths
  - Multiple possible paths
  - Duplicate tickets
  - Circular paths
- RESTful API interface
- Comprehensive test coverage

## Getting Started

### Prerequisites

- Go
- Git

### Installation

```bash
git clone https://github.com/kidus-tiliksew/professional-itinerary.git
cd professional-itinerary
go mod tidy
```

### Running the Service

```bash
go run main.go
```

The service will be available at http://localhost:8080.

## API Usage

### Create an Itinerary

Endpoint: `POST /itinerary`
Request Body: An array of ticket pairs, where each ticket is represented as [source, destination].
Example Request:

```json
[
  ["SFO", "LAX"],
  ["LAX", "JFK"]
]
```

Example Response:

```json
["SFO", "LAX", "JFK"]
```

Error Responses:
400 Bad Request: When the itinerary cannot be constructed due to issues like disconnected paths, multiple possible paths, or duplicate tickets.

## Running Tests

```bash
go test ./...
```

## Design Decisions

### Libraries and Tools

- Echo Framework
- Testing: Using Go's built-in testing framework with testify for assertions.

### Algorithm

The itinerary construction algorithm uses a graph-based approach:

- Build a directed graph from the tickets
- Identify potential starting points (locations with no incoming tickets)
- Construct the itinerary by following the graph from the starting point
- Validate that all locations are visited

This approach efficiently handles various edge cases:

- Circular paths are detected by tracking visited locations
- Disconnected paths are identified when not all locations can be visited
- Multiple possible paths are detected when a location has multiple outgoing tickets