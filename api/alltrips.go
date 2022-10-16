package api

import (
	"log"
)

//go:generate genopts --params --function AllTrips --extends Trips,Base debug
func (c *Client) AllTrips(optss ...AllTripsOption) (chan getTripsInfoTrip, chan error) {
	opts := MakeAllTripsOptions(optss...)

	data, errs := make(chan getTripsInfoTrip), make(chan error)
	var nextCursor string
	go func() {
		for {
			os := opts.ToTripsOptions()
			if nextCursor != "" {
				os = append(os, TripsCursor(nextCursor))
			}
			res, err := c.Trips(os...)
			if err != nil {
				errs <- err
				continue
			}
			for _, t := range res.Trips {
				data <- t
			}

			if opts.Debug() {
				log.Printf("Got %d trips, hasMore:%t", len(res.Trips), res.PagingResult.HasMore)
			}

			if !res.PagingResult.HasMore {
				break
			}

			nextCursor = res.PagingResult.NextCursor
		}
		close(data)
		close(errs)
	}()

	return data, errs
}
