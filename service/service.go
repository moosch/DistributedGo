package service

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/moosch/DistributedGo/registry"
)

func Start(ctx context.Context, host, port string, reg registry.Registration, registerHandlersFunc func()) (context.Context, error) {
	registerHandlersFunc()
	ctx = startService(ctx, reg.ServiceName, host, port)
	err := registry.RegisterService(reg)
	if err != nil {
		return ctx, err
	}

	return ctx, nil
}

func startService(ctx context.Context, serviceName registry.ServiceName, host, port string) context.Context {
	ctx, cancel := context.WithCancel(ctx)

	var server http.Server
	server.Addr = ":" + port

	go func() {
		// TODO(moosch): Transport Layer Security
		log.Println(server.ListenAndServe())
		// ListenAndServe only returns if there's an error starting.
		cancel()
	}()

	go func() {
		fmt.Printf("%v service started. Press any key to stop.\n", serviceName)
		var s string
		fmt.Scanln(&s)
		err := registry.ShutdownService(fmt.Sprintf("http://%v:%v", host, port))
		if err != nil {
			// Just log error to allow shutdown process to continue.
			log.Println(err)
		}
		server.Shutdown(ctx)
		cancel()
	}()

	return ctx
}
