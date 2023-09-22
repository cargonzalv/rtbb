package server

//go:generate mockgen  -destination=mocks.gen.go -source=$GOFILE -package=$GOPACKAGE

// The web server configurable parameters.
type RestServerParams struct {
	// Web server port.
	Port                int
	MaxConnsPerIp       int
	MaxReqPerConn       int
	ReadTimeoutSeconds  int
	WriteTimeoutSeconds int
	IdleTimeoutSeconds  int
	Name                string
}

// Server Interface is to define the rest server behaviour.
type Server interface {
	Start() error
	Shutdown() error
}
