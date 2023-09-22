/*
The Health Check Probe package.
*/
package probe

//go:generate mockgen  -destination=mocks.gen.go -source=interface.go -package=probe

// The Health check probe interface.
type Probe interface {
	// Function returns the name of the health check probe.
	Name() string
	// Health check process function.
	Check() bool
	// If function return true, the service readiness state switch to false.
	IsCritical() bool
}
