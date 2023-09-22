# Development process

## Service architecture and project layout

![Alt](diagrams/rtb-bidder.svg)

[Info on components](../README.md#architecture-overview)

The rtb-service contains various functional packages. The usage of the packages are decoupled from the implementation by leveraging the interfaces and resolving them at a startup.

## Dependency injection

Each package provides domain-specific functionality. 

- app package is responsible for wiring all dependencies and starting the service;
- observability package is reponsible for metrics, logs, traces;
- serializer package incapsulates any marshaling and unmarshalling functionality;
- health package provides health monitoring of the servise components;
- bidding package contains the business logic of the real time bidding;
- server package provides routing and serving of web requests;

Packages providing and consuming functionality of each other. Consuming of services done through dependency injection. The code in the package uses required services declared as interfaces. The bootstrap sets this implementation of those interfaces during startup procedure. This architecture allows to create decoupled and testable code. 

The following code of the health package is shown to demonstrate this architecture. Other packages in the service following exactly same pattern.

The package contains those files:

- handler.go - the file contains the definition of the health related web requests handlers and implements the Handler interface. The package consumes the `Observer` interface;
- interface.go - the file contains the definition of the interfaces provided by the package;
- observer.go - the file contains the implementation of `Observer` interface;
- router.go - the file contains the implementation of `Router` interface and consumes the `Handler` interface;
- wire.go - the file contains providers of package interfaces;

interface.go content 
```go
type Observer interface {
	IsLive() bool
	IsReady() bool
	StartMonitors(context.Context)
	StopMonitors()
}

// Handler Interface is for the health web requests handler.
type Handler interface {
	GetLivenessHandler(ctx *fasthttp.RequestCtx)
	GetReadinessHandler(ctx *fasthttp.RequestCtx)
}

// Router is proxy type of bidder.Router required by wire resolver.
type Router bidder.Router
```

handler.go
```go
package health

import (
	"fmt"

	"github.com/adgear/go-commons/pkg/log"
	"github.com/valyala/fasthttp"
)

// Enforce `Handler` interface implementation.
var _ Handler = (*controller)(nil)

type controller struct {
	observer Observer
	logger  log.Service
}

// GetLivenessHandler handles `liveness probe` requests.
func (c *controller) GetLivenessHandler(ctx *fasthttp.RequestCtx) {
	setResponseContext(ctx)
	isLive := c.observer.IsLive()
	fmt.Fprintf(ctx, `{"is_live":%t}`, isLive)
}

// GetReadinessHandler handles `readiness probe` requests.
func (c *controller) GetReadinessHandler(ctx *fasthttp.RequestCtx) {
	setResponseContext(ctx)
	isReady := c.observer.IsReady()
	fmt.Fprintf(ctx, `{"is_ready":%t}`, isReady)
}

func setResponseContext(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Connection", "keep-alive")
	ctx.SetContentType("application/json")
}
```
Important information to mention.

The `controller` is a structure which contains the consumed services - observer and logger.
In this case the `observer` provides the functionality described in the `Observer` interface.
The `logger` service provide the log.Service functionality.
The controller is a private structure and its content is not accessible outside of the package.
```go
type controller struct {
	observer Observer
	logger  log.Service
}
```

The `controller` structure has attached methods which implementing the `Handler` interface.

This line enforcing the implementation of the `Handler` interface.
```go
var _ Handler = (*controller)(nil)
```

The implemented instances of `Observer` and `log.Service` services are setting up in the providers, defined in the `wire.go` file.

wire.go content
```go
package health

import (
	"time"

	"github.com/adgear/go-commons/pkg/log"
	"github.com/adgear/rtb-bidder/config"
	"github.com/adgear/rtb-bidder/internal/bidder/health/probe"
)

// ProvideService provide health `Observer` implementation.
func ProvideService(logger log.Service, probes []probe.Probe, cfg *config.Config) Observer {
	loopInterval := time.Duration(cfg.HealthCheck.LoopIntervalSeconds) * time.Second
	return &service{
		logger:       logger,
		loopInterval: loopInterval,
		probes:       probes,
		status:       initStatus(probes),
	}
}

// ProvideHandler provide `Handler` implementation.
func ProvideHandler(logger log.Service, observer Observer) Handler {
	return &controller{observer, logger}
}

// ProvideRouter provide `bidder.Router` implementation.
func ProvideRouter(handler Handler) Router {
	return &router{
		handler: handler,
	}
}
```

As we can see the `ProvideHandler` is consuming `log.Service` and `Observer` interfaces and provide the `Handler` interface.

```go
func ProvideHandler(logger log.Service, observer Observer) Handler {
	return &controller{observer, logger}
}
```

The `ProvideRouter` function is consuming the `Handler` interface and providing the `Router` interface.

```go
func ProvideRouter(handler Handler) Router {
	return &router{
		handler: handler,
	}
}
```

We are using [wire](https://github.com/google/wire) code generation tool to generate `Dependency Injection bootstrap`.

The `app` package is an application container. The wire.go file looks like this:

wire.go in the `app` package content
```go
//go:build wireinject
// +build wireinject

package app

import (
	"github.com/adgear/rtb-bidder/config"
	"github.com/adgear/rtb-bidder/internal/bidder"
	"github.com/adgear/rtb-bidder/internal/bidder/bidding"
	"github.com/adgear/rtb-bidder/internal/bidder/health"
	"github.com/adgear/rtb-bidder/internal/bidder/health/probe"
	"github.com/adgear/rtb-bidder/internal/bidder/server"
	"github.com/adgear/rtb-bidder/internal/bidder/server/handler"
	"github.com/adgear/rtb-bidder/internal/observability"
	"github.com/adgear/rtb-bidder/internal/serializer"
	"github.com/google/wire"
)

// ProvideRouters is the function providing the list of routers from all packages, which have `bidder.Router` interface implemented.
func ProvideRouters(health health.Router, bidding bidding.Router, handler handler.Router, observability observability.Router) *[]bidder.Router {
	return &[]bidder.Router{
		health,
		bidding,
		handler,
		observability,
	}
}

// Using by the wire to generate gependency tree.
// This file is not include into release.
func Wire() *App {
	panic(wire.Build(
		// observability package providers
		observability.ProvideLogger,
		observability.ProvideMetrics,
		observability.ProvideRouter,

		// application Config providers
		config.ProvideConfig,

		// server package providers
		server.ProvideRestServer,
		server.ProvideFastHttpRouter,

		// bidding package providers
		bidding.ProvideBidder,
		bidding.ProvideHandler,
		bidding.ProvideRouter,

		// health package providers
		health.ProvideService,
		health.ProvideHandler,
		health.ProvideRouter,

		handler.ProvideRouter,

		// probe package providers
		probe.ProvideProbes,

		// serializer package providers
		serializer.ProvideSerializer,

		ProvideRouters,
		wire.Struct(new(App), "*"),
	))
}
```

The `Wire` function put together providers from all packages and allow to `wire tool` to build dependency tree and generate bootstrap function.

To install [wire](https://github.com/google/wire) run:

```shell
go install github.com/google/wire/cmd/wire@lates
```

To generate the wire bootstrap module, run from project root directory:

```shell
wire gen ./internal/bidder/app
```

This command will generate `wire_gen.go` file in the `app` package.

The content of `wire_gen.go` file looks like this:

```go
// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"github.com/adgear/rtb-bidder/config"
	"github.com/adgear/rtb-bidder/internal/bidder"
	"github.com/adgear/rtb-bidder/internal/bidder/bidding"
	"github.com/adgear/rtb-bidder/internal/bidder/health"
	"github.com/adgear/rtb-bidder/internal/bidder/health/probe"
	"github.com/adgear/rtb-bidder/internal/bidder/server"
	"github.com/adgear/rtb-bidder/internal/bidder/server/handler"
	"github.com/adgear/rtb-bidder/internal/observability"
	"github.com/adgear/rtb-bidder/internal/serializer"
)

// Injectors from wire.go:

// Using by the wire to generate gependency tree.
// This file is not include into release.
func Wire() *App {
	configConfig := config.ProvideConfig()
	service := observability.ProvideLogger(configConfig)
	metricService := observability.ProvideMetrics(configConfig)
	router := server.ProvideFastHttpRouter()
	v := probe.ProvideProbes(metricService)
	observer := health.ProvideService(service, v, configConfig)
	healthHandler := health.ProvideHandler(service, observer)
	healthRouter := health.ProvideRouter(healthHandler)
	bidder := bidding.ProvideBidder(service, metricService)
	serializerService := serializer.ProvideSerializer()
	biddingHandler := bidding.ProvideHandler(service, bidder, serializerService)
	biddingRouter := bidding.ProvideRouter(biddingHandler)
	handlerRouter := handler.ProvideRouter(service, metricService, configConfig)
	observabilityRouter := observability.ProvideRouter()
	v2 := ProvideRouters(healthRouter, biddingRouter, handlerRouter, observabilityRouter)
	serverServer := server.ProvideRestServer(service, metricService, router, v2, configConfig)
	app := &App{
		logger:        service,
		metrics:       metricService,
		restServer:    serverServer,
		healthService: observer,
	}
	return app
}

// wire.go:

// ProvideRouters is the function providing the list of routers from all packages, which have `bidder.Router` interface implemented.
func ProvideRouters(health2 health.Router, bidding2 bidding.Router, handler2 handler.Router, observability2 observability.Router) *[]bidder.Router {
	return &[]bidder.Router{health2, bidding2, handler2, observability2}
}
```

The `Wire() *App` initializes and resolves all dependency in the right order and returns the pointer to the `App` struct.

## Adding endpoint routes to the fasthttp router

If the package responsible for the handling web request, the package will contain the `router.go` file.
As an example here is the content of the `router.go` of the `health` package:

```go
package health

import (
	"github.com/adgear/rtb-bidder/internal/bidder"
	"github.com/buaazp/fasthttprouter"
)

// Enforce `bidder.Router` interface implementation.
var _ bidder.Router = (*router)(nil)

type router struct {
	// Health monitoring controller.
	handler Handler
}

// AddRoutes add `bidding package` routes.
func (r *router) AddRoutes(router *fasthttprouter.Router) {
	// k8s health probes handlers. This probes used to control health of
	// the application.
	// * Failing liveness probe will restart the container.
	// * Failing readiness probe will stop the application from serving traffic.
	router.GET("/health/liveness", r.handler.GetLivenessHandler)
	router.GET("/health/readiness", r.handler.GetReadinessHandler)
}
```

The router module is implementing the `bidder.Router` interface. The interface has `AddRoutes` function, which receives the pointer to the web server router and adds handlers for `/health/liveness` and `/helath/readiness` endpoints.

The `wire.go` module has 

```go
func ProvideRouter(handler Handler) Router {
	return &router{
		handler: handler,
	}
}
```

The `Router` interface is the proxy type to `bidder.Router`.

```go
type Router bidder.Router
```

We have to do this to allow wire to do the proper dependency resoulution. 

The `wire.go` module in the `app` package contains the function which consumes `Router` interfaces from all modules and provide `[]bidder.Router` list of all modules which implemented `bidder.Router` interface. 
`[]bidder.Router` list of routers consumed by `server` package. Before server start, it itterate over all routers and regester routes from each package.

## Healthchcks(liveness and readiness endpoints)

The `liveness endpoint probe` checks the container health. If liveness probe fails, the k8s will restarts the container.

The `readiness endpoint probe` checks the application ability to serve incomming requests. The probe will pass when some conditions are met e.g. initializing the db connection pool, populating cache, waiting for some other service to be alive etc.

The service API layer is using the `observer` service, which is responsible for observing the application health.
The service implements `Observer` interface.

```go
type Observer interface {
	IsLive() bool
	IsReady() bool
	StartMonitors(context.Context)
	StopMonitors()
}
```

`IsLive` function is responsible to check liveness health of the application and using in `liveness endpoint`.
`IsReady` function is responsible to check the readiness of the service to server incomming requests. The `observer` service has list of probes. The probe is implementing the interface of `Probe`.

```go
type Probe interface {
	Name() string
	Check() bool
	IsCritical() bool
}
```

`Name()` function returns the name of the probe.
`Check()` function provide the readiness check, e.g. running some empty query from the db(SELECT GETDATE();)
`IsCritical()` function return `true` if probe component is critical for serving request. If outage of the component is not affecting the request, will return `false`(f.e. The kafka component with offline capabilities).

The `observer` service will run periodic checks of all probes and in case of outage of critical component, the `readiness probe` will return false and k8s will stop serving requests to the service.

Here is the example of metrics readiness probe.

```go
package probe

import "github.com/adgear/go-commons/pkg/metric"

// Enforce Probe interface implementation.
var _ Probe = (*metricProbe)(nil)

// The metric probe struct.
type metricProbe struct {
	metrics metric.Service
}

// The function returns metric probe name.
func (*metricProbe) Name() string {
	return "metrics"
}

// The function process metric service health check.
func (p *metricProbe) Check() bool {
	return p.metrics.IsReady()
}

func (p *metricProbe) IsCritical() bool {
	return false
}
```