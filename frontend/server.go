package frontend

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/spudtrooper/minimalcli/handler"
	"github.com/spudtrooper/uber/api"
	"github.com/spudtrooper/uber/handlers"
)

func ListenAndServe(ctx context.Context, client *api.Client, port int, host string) error {
	var hostPort string
	if host == "localhost" {
		hostPort = fmt.Sprintf("http://localhost:%d", port)
	} else {
		hostPort = fmt.Sprintf("https://%s", host)
	}

	mux := http.NewServeMux()
	handler.Init(mux)
	if err := handler.AddHandlers(ctx, mux, handlers.CreateHandlers(client),
		handler.AddHandlersPrefix("api"),
		handler.AddHandlersKey("uber"),
		handler.AddHandlersIndexTitle("unofficial uber API"),
		handler.AddHandlersFooterHTML(`Details: <a target="_" href="//github.com/spudtrooper/uber">github.com/spudtrooper/uber</a>`),
		handler.AddHandlersSourceLinks(true),
		handler.AddHandlersSerializedSourceLocations(handlers.SourceLocations),
	); err != nil {
		return err
	}
	mux.Handle("/", http.RedirectHandler("/api", http.StatusSeeOther))

	log.Printf("listening on %s", hostPort)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux); err != nil {
		return err
	}

	return nil
}
