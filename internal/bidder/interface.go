package bidder

import "github.com/buaazp/fasthttprouter"

// Router is the interface for packages to register web requests handlers.
type Router interface {
	AddRoutes(router *fasthttprouter.Router)
}
