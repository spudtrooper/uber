package render

import (
	"github.com/spudtrooper/minimalcli/handler"
	"github.com/spudtrooper/uber/api"
)

func Trips(input any) ([]byte, handler.RendererConfig, error) {
	params := input.(*api.TripsInfo)

	return tripsFromTrips(params.Trips)
}
