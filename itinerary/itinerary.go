package itinerary

import (
	"fmt"
)

// ItineraryFromTickets constructs an itinerary from a list of tickets
func ItineraryFromTickets(tickets [][]string) ([]string, error) {
	if len(tickets) == 0 {
		return nil, fmt.Errorf("no tickets provided")
	}

	// Build a graph of all connections
	graph := make(map[string][]string)
	inDegree := make(map[string]int)

	// Track all locations
	allLocations := make(map[string]bool)

	for _, ticket := range tickets {
		src, dst := ticket[0], ticket[1]
		graph[src] = append(graph[src], dst)
		inDegree[dst]++
		allLocations[src], allLocations[dst] = true, true
	}

	// Find potential starting points (nodes with no incoming edges)
	var startingPoints []string
	for loc := range allLocations {
		if inDegree[loc] == 0 {
			startingPoints = append(startingPoints, loc)
		}
	}

	// If no clear starting point, use any location
	if len(startingPoints) == 0 {
		// This could be a circular path - pick any node
		startingPoints = append(startingPoints, tickets[0][0])
	}

	// Check for multiple possible paths
	for _, dests := range graph {
		if len(dests) > 1 {
			return nil, fmt.Errorf("multiple possible paths")
		}
	}

	// Try each potential starting point
	var validItinerary []string
	for _, start := range startingPoints {
		itinerary := []string{start}
		visited := make(map[string]bool)
		visited[start] = true

		current := start
		for {
			if destinations, ok := graph[current]; ok && len(destinations) > 0 {
				next := destinations[0]

				if visited[next] {
					return nil, fmt.Errorf("circular itinerary")
				}

				itinerary = append(itinerary, next)
				visited[next] = true
				current = next
			} else {
				break
			}
		}

		// Check if this itinerary visits all locations
		if len(visited) == len(allLocations) {
			validItinerary = itinerary
			break
		}
	}

	// If we couldn't find a valid itinerary that visits all locations, the path is disconnected
	if validItinerary == nil {
		return nil, fmt.Errorf("disconnected itinerary")
	}

	return validItinerary, nil
}
