package main

import (
	"net/http"

	"github.com/kidus-tiliksew/professional-itinerary/itinerary"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)

	e.POST("/itinerary", func(c echo.Context) error {
		// Parse the request body into a slice of slices of strings
		var tickets [][]string
		if err := c.Bind(&tickets); err != nil {
			return err
		}

		// Create the itinerary
		i, err := itinerary.ItineraryFromTickets(tickets)
		if err != nil {
			c.Logger().Warn("failed to create itinerary: ", err, "input: ", tickets)
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, i)
	})
	
	e.Logger.Fatal(e.Start(":8080"))
}
