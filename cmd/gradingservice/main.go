package main

import (
	"context"
	"fmt"
	stdlog "log"

	"github.com/moosch/DistributedGo/grades"
	"github.com/moosch/DistributedGo/log"
	"github.com/moosch/DistributedGo/registry"
	"github.com/moosch/DistributedGo/service"
)

func main() {
	// TODO(moosch): Pull this in from config file or env
	host, port := "localhost", "6000"
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	var reg registry.Registration
	reg.ServiceName = registry.GradingService
	reg.ServiceURL = serviceAddress
	reg.RequiredServices = []registry.ServiceName{registry.LogService}
	reg.ServiceUpdateURL = reg.ServiceURL + "/services"

	ctx, err := service.Start(
		context.Background(),
		host, port,
		reg,
		grades.RegisterHandlers,
	)
	if err != nil {
		stdlog.Fatal(err)
	}

	if logProvider, err := registry.GetProvider(registry.LogService); err == nil {
		fmt.Printf("Logging service was found at: %v\n", logProvider)
		log.SetClientLogger(logProvider, reg.ServiceName)
	}

	// Only continues past here when the service errors or shuts down.
	<-ctx.Done()
	fmt.Println("Shutting down grading service")
}
