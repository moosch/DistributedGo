package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/moosch/DistributedGo/registry"
)

func main() {
	http.Handle("/services", &registry.RegistryService{})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var server http.Server
	server.Addr = registry.ServerPort

	go func() {
		// TODO(moosch): Transport Layer Security
		log.Println(server.ListenAndServe())
		// ListenAndServe only returns if there's an error starting.
		cancel()
	}()

	go func() {
		fmt.Println("Registry service started. Press any key to stop.")
		var s string
		fmt.Scanln(&s)
		server.Shutdown(ctx)
		cancel()
	}()

	<-ctx.Done()
	fmt.Println("Shutting down Registry service.")
}
