package api

import (
	"log"
)

//go:generate genopts --params --function AllTrips --extends Trips,Base debug totalLimit:int
func (c *Client) AllTrips(optss ...AllTripsOption) (chan TripsInfoTrip, chan error) {
	opts := MakeAllTripsOptions(optss...)

	data, errs := make(chan TripsInfoTrip), make(chan error)
	var nextCursor string
	go func() {
		for i := 0; ; i++ {
			if opts.TotalLimit() > 0 && i > opts.TotalLimit() {
				log.Printf("break because at limit of %d", opts.TotalLimit())
				break
			}
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
