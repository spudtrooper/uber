package api

import (
	goutilerrors "github.com/spudtrooper/goutil/errors"
	"github.com/spudtrooper/goutil/parallel"
)

type AllTripsBatchInfo struct {
	Trips []TripsInfoTrip `json:"trips"`
}

// TODO: genopts should transitively extend options
//go:generate genopts --params --function AllTripsBatch --extends AllTrips,Trips,Base
func (c *Client) AllTripsBatch(optss ...AllTripsBatchOption) (*AllTripsBatchInfo, error) {
	opts := MakeAllTripsBatchOptions(optss...)

	data, errs := c.AllTrips(opts.ToAllTripsOptions()...)
	var trips []TripsInfoTrip
	errBuilder := goutilerrors.MakeErrorCollector()
	parallel.WaitFor(func() {
		for d := range data {
			trips = append(trips, d)
		}
	}, func() {
		for e := range errs {
			errBuilder.Add(e)
		}
	})

	if err := errBuilder.Build(); err != nil {
		return nil, err
	}
	res := &AllTripsBatchInfo{
		Trips: trips,
	}
	return res, nil
}
