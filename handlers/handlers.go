// Package handlers is a bridge between the API and handlers to create a CLI/API server.
package handlers

import (
	"context"
	_ "embed"

	"github.com/spudtrooper/minimalcli/handler"
	"github.com/spudtrooper/uber/api"
	"github.com/spudtrooper/uber/render"
)

//go:generate minimalcli gsl --input handlers.go --uri_root "github.com/spudtrooper/uber/blob/main/handlers" --output handlers.go.json
//go:embed handlers.go.json
var SourceLocations []byte

func CreateHandlers(client *api.Client) []handler.Handler {
	b := handler.NewHandlerBuilder()

	b.NewHandler("Status",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.StatusParams)
			return client.Status(p.Options()...)
		},
		api.StatusParams{},
		handler.NewHandlerExtraRequiredFields([]string{"sid", "csid"}),
		handler.NewHandlerRenderer(render.Status),
	)

	b.NewHandler("User",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.UserParams)
			return client.User(p.Options()...)
		},
		api.UserParams{},
		handler.NewHandlerExtraRequiredFields([]string{"sid", "csid"}),
	)

	b.NewHandler("Trips",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.TripsParams)
			return client.Trips(p.Options()...)
		},
		api.TripsParams{},
		handler.NewHandlerExtraRequiredFields([]string{"sid", "csid"}),
		handler.NewHandlerRenderer(render.Trips),
	)

	b.NewHandler("AllTrips",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.AllTripsBatchParams)
			return client.AllTripsBatch(p.Options()...)
		},
		api.AllTripsBatchParams{},
		handler.NewHandlerExtraRequiredFields([]string{"sid", "csid"}),
		handler.NewHandlerRenderer(render.AllTripsBatch),
	)

	b.NewHandler("FareEstimate",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.FareEstimateParams)
			return client.FareEstimate(p.Options()...)
		},
		api.FareEstimateParams{},
		handler.NewHandlerExtraRequiredFields([]string{"sid", "csid"}),
	)

	return b.Build()
}
