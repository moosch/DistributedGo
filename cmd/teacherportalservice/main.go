package main

import (
	"context"
	"fmt"
	stdlog "log"

	"github.com/moosch/DistributedGo/log"
	"github.com/moosch/DistributedGo/registry"
	"github.com/moosch/DistributedGo/service"
	"github.com/moosch/DistributedGo/teacherportal"
)

func main() {
	err := teacherportal.ImportTemplates()
	if err != nil {
		stdlog.Fatal(err)
	}

	host, port := "localhost", "5000"
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	var reg registry.Registration
	reg.ServiceName = registry.TeacherPortal
	reg.ServiceURL = serviceAddress
	reg.RequiredServices = []registry.ServiceName{
		registry.LogService,
		registry.GradingService,
	}
	reg.ServiceUpdateURL = reg.ServiceURL + "/services"
	reg.HeartbeatURL = reg.ServiceURL + "/heartbeat"

	ctx, err := service.Start(
		context.Background(),
		host,
		port,
		reg,
		teacherportal.RegisterHandlers)
	if err != nil {
		stdlog.Fatal(err)
	}
	if logProvider, err := registry.GetProvider(registry.LogService); err == nil {
		log.SetClientLogger(logProvider, reg.ServiceName)
	}

	<-ctx.Done()
	fmt.Println("Shutting down teacher portal")
}
