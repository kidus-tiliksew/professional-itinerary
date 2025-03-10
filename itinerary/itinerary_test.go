package itinerary_test

import (
	"testing"

	"github.com/kidus-tiliksew/professional-itinerary/itinerary"
	"github.com/stretchr/testify/assert"
)

func TestItineraryFromTickets(t *testing.T) {
	tests := []struct {
		name          string
		tickets       [][]string
		wantItinerary []string
		wantErr       bool
		errMsg        string
	}{
		{
			name: "simple path",
			tickets: [][]string{
				{"SFO", "ATL"},
				{"ATL", "JFK"},
			},
			wantItinerary: []string{"SFO", "ATL", "JFK"},
			wantErr:       false,
		},
		{
			name: "different order",
			tickets: [][]string{
				{"ATL", "JFK"},
				{"SFO", "ATL"},
			},
			wantItinerary: []string{"SFO", "ATL", "JFK"},
			wantErr:       false,
		},
		{
			name: "longer path",
			tickets: [][]string{
				{"LAX", "DXB"},
				{"JFK", "LAX"},
				{"SFO", "SJC"},
				{"DXB", "SFO"},
			},
			wantItinerary: []string{"JFK", "LAX", "DXB", "SFO", "SJC"},
			wantErr:       false,
		},
		{
			name:          "empty tickets",
			tickets:       [][]string{},
			wantItinerary: nil,
			wantErr:       true,
			errMsg:        "no tickets provided",
		},
		{
			name: "disconnected path",
			tickets: [][]string{
				{"SFO", "ATL"},
				{"JFK", "LHR"},
			},
			wantItinerary: nil,
			wantErr:       true,
			errMsg:        "disconnected itinerary",
		},
		{
			name: "multiple starting points",
			tickets: [][]string{
				{"SFO", "ATL"},
				{"JFK", "SFO"},
				{"DXB", "JFK"},
			},
			wantItinerary: []string{"DXB", "JFK", "SFO", "ATL"},
			wantErr:       false,
		},
		{
			name: "circular path",
			tickets: [][]string{
				{"SFO", "ATL"},
				{"ATL", "JFK"},
				{"JFK", "SFO"},
			},
			wantItinerary: nil,
			wantErr:       true,
			errMsg:        "circular itinerary",
		},
		{
			name: "multiple possible paths",
			tickets: [][]string{
				{"SFO", "ATL"},
				{"SFO", "JFK"},
				{"ATL", "DXB"},
			},
			wantItinerary: nil,
			wantErr:       true,
			errMsg:        "multiple possible paths",
		},
		{
			name: "single ticket",
			tickets: [][]string{
				{"SFO", "ATL"},
			},
			wantItinerary: []string{"SFO", "ATL"},
			wantErr:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotItinerary, err := itinerary.ItineraryFromTickets(tt.tickets)

			// Check error expectations
			if (err != nil) != tt.wantErr {
				t.Errorf("ItineraryFromTickets() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// If expecting an error, check error message
			if tt.wantErr && err != nil && tt.errMsg != "" {
				if err.Error() != tt.errMsg {
					t.Errorf("ItineraryFromTickets() error message = %v, want to contain %v", err.Error(), tt.errMsg)
				}
				return
			}

			// Check itinerary
			if !assert.Equal(t, gotItinerary, tt.wantItinerary) {
				t.Errorf("ItineraryFromTickets() = %v, want %v", gotItinerary, tt.wantItinerary)
			}
		})
	}
}
