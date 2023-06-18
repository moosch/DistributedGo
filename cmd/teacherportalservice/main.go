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

	var r registry.Registration
	r.ServiceName = registry.TeacherPortal
	r.ServiceURL = serviceAddress
	r.RequiredServices = []registry.ServiceName{
		registry.LogService,
		registry.GradingService,
	}
	r.ServiceUpdateURL = r.ServiceURL + "/services"

	ctx, err := service.Start(context.Background(),
		host,
		port,
		r,
		teacherportal.RegisterHandlers)
	if err != nil {
		stdlog.Fatal(err)
	}
	if logProvider, err := registry.GetProvider(registry.LogService); err == nil {
		log.SetClientLogger(logProvider, r.ServiceName)
	}

	<-ctx.Done()
	fmt.Println("Shutting down teacher portal")

}
