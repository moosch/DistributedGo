package registry

type ServiceName string

type Registration struct {
	ServiceName ServiceName
	ServicesURL string
}

const (
	LogService = ServiceName("LogService")
)
