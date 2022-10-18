package render

import (
	"github.com/spudtrooper/minimalcli/handler"
	"github.com/spudtrooper/uber/api"
)

func AllTripsBatch(input any) ([]byte, handler.RendererConfig, error) {
	params := input.(*api.AllTripsBatchInfo)

	return tripsFromTrips(params.Trips)
}
