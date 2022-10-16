package api

import (
	"log"
)

//go:generate genopts --params --function AllTrips --extends Trips,Base debug
func (c *Client) AllTrips(optss ...AllTripsOption) (chan getTripsInfo, chan error) {
	opts := MakeAllTripsOptions(optss...)

	data, errs := make(chan getTripsInfo), make(chan error)
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

			data <- *res

			if opts.Debug() {
				log.Printf("Got %d trips, hasMore:%t", len(res.Trips), res.PagingResult.HasMore)
			}

			if !res.PagingResult.HasMore {
				return
			}

			nextCursor = res.PagingResult.NextCursor
		}
		close(data)
		close(errs)
	}()

	return data, errs
}
