package registry

type ServiceName string

type Registration struct {
	ServiceName      ServiceName
	ServiceURL       string
	RequiredServices []ServiceName
	ServiceUpdateURL string
}

const (
	LogService     = ServiceName("LogService")
	GradingService = ServiceName("GradingService")
	TeacherPortal  = ServiceName("TeacherPortal")
)

type patchEntry struct {
	Name ServiceName
	URL  string
}

type patch struct {
	Added   []patchEntry
	Removed []patchEntry
}
