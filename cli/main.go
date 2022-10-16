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
	flag.Duration("timeout", 0, "timeout")
	flag.Bool("debug", false, "debug")

	client, err := api.NewClientFromFlags()
	if err != nil {
		return err
	}

	return handler.RunCLI(ctx, handlers.CreateHandlers(client)...)
}
