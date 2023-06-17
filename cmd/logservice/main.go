package main

import (
	"context"
	"fmt"
	stdlog "log"

	"github.com/moosch/DistributedGo/log"
	"github.com/moosch/DistributedGo/service"
)

func main() {
	log.Run("app.log")

	// TODO(moosch): Pull this in from config file or env
	host, port := "localhost", "4000"

	ctx, err := service.Start(
		context.Background(),
		"Log Service",
		host, port,
		log.RegisterHandlers,
	)
	if err != nil {
		stdlog.Fatal(err)
	}
	// Only continues past here when the service errors or shuts down.
	<-ctx.Done()
	fmt.Println("Shutting down log service")
}
