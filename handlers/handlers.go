// Package handlers is a bridge between the API and handlers to create a CLI/API server.
package handlers

import (
	"context"
	_ "embed"

	"github.com/spudtrooper/minimalcli/handler"
	"github.com/spudtrooper/uber/api"
)

//go:generate minimalcli gsl --input handlers.go --uri_root "github.com/spudtrooper/uber/blob/main/handlers" --output handlers.go.json
//go:embed handlers.go.json
var SourceLocations []byte

func CreateHandlers(client *api.Client) []handler.Handler {
	b := handler.NewHandlerBuilder()

	b.NewHandler("GetStatus",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.GetStatusParams)
			return client.GetStatus(p.Options()...)
		},
		api.GetStatusParams{},
		handler.NewHandlerExtraRequiredFields([]string{"sid", "csid"}),
	)

	b.NewHandler("GetUser",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.GetUserParams)
			return client.GetUser(p.Options()...)
		},
		api.GetUserParams{},
		handler.NewHandlerExtraRequiredFields([]string{"sid", "csid"}),
	)

	return b.Build()
}
