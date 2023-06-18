package main

import (
	"context"
	"fmt"
	stdlog "log"

	"github.com/moosch/DistributedGo/log"
	"github.com/moosch/DistributedGo/registry"
	"github.com/moosch/DistributedGo/service"
)

func main() {
	log.Run("app.log")

	// TODO(moosch): Pull this in from config file or env
	host, port := "localhost", "4000"
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	var reg registry.Registration
	reg.ServiceName = registry.LogService
	reg.ServiceURL = serviceAddress
	// Log service doesn't need any services.
	reg.RequiredServices = make([]registry.ServiceName, 0)
	reg.ServiceUpdateURL = reg.ServiceURL + "/services"
	reg.HeartbeatURL = reg.ServiceURL + "/heartbeat"

	ctx, err := service.Start(
		context.Background(),
		host, port,
		reg,
		log.RegisterHandlers,
	)
	if err != nil {
		stdlog.Fatal(err)
	}
	// Only continues past here when the service errors or shuts down.
	<-ctx.Done()
	fmt.Println("Shutting down log service")
}
