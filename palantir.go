package palantir

import (
	"os"
	"runtime"
	"runtime/debug"
)


type Report struct {
	// identifiers
	Category string
	Version string
	// os.Hostname()
	Hostname string
	// os.Environ()
	EnvironmentVariables []string
	// runtime.GOOS
	OS string
	Stack []byte
	Message string
}


func NewReport(category, message string) *Report {
	h, _ := os.Hostname()
	return &Report{
		Category: category,
		Message: message,
		Version: runtime.Version(),
		Hostname: h,
		EnvironmentVariables: os.Environ(),
		OS: runtime.GOOS,
		Stack: debug.Stack(),
	}
}