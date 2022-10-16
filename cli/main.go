package cli

import (
	"context"
	"flag"

	"github.com/spudtrooper/minimalcli/handler"
	"github.com/spudtrooper/uber/api"
	"github.com/spudtrooper/uber/handlers"
)

func Main(ctx context.Context) error {
	flag.String("cursor", "", "cursor")

	client, err := api.NewClientFromFlags()
	if err != nil {
		return err
	}

	return handler.RunCLI(ctx, handlers.CreateHandlers(client)...)
}
