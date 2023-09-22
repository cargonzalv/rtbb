# "rtb-bidder"

RTB Bidder service which serves O&O (TV Native) traffic

## Description

"rtb-bidder" is a next-generation Realtime Bidder service that serves O&O (TV Native) traffic.

## Important links

### Grafana dashboard

[RTB Bidder Golden Signals dashboard](https://grafana.int.adgear.com/d/0JYSf334k/rtb-bidder-golden-signals?orgId=1&refresh=10s&var-datasource=use1-rdev&var-cluster=use1-rdev&var-namespace=rtb-bidder&var-pod=rtb-bidder-0&from=now-5m&to=now)

### Logs search examples

[All logs from rtb-bidder service on `ues1-rdev` cluster](https://grafana.int.adgear.com/explore?orgId=1&left=%7B%22datasource%22:%22P9A714CFC031CC693%22,%22queries%22:%5B%7B%22refId%22:%22A%22,%22expr%22:%22%7Bcluster%3D%5C%22use1-rdev%5C%22,%20namespace%3D%5C%22rtb-bidder%5C%22%7D%20%7C%20json%22,%22queryType%22:%22range%22,%22datasource%22:%7B%22type%22:%22loki%22,%22uid%22:%22P9A714CFC031CC693%22%7D,%22editorMode%22:%22code%22%7D%5D,%22range%22:%7B%22from%22:%22now-30m%22,%22to%22:%22now%22%7D%7D)

Query:

```
{cluster="use1-rdev", namespace="rtb-bidder"} | json
```

[Search text in the logs](https://grafana.int.adgear.com/explore?orgId=1&left=%7B%22datasource%22:%22P9A714CFC031CC693%22,%22queries%22:%5B%7B%22refId%22:%22A%22,%22expr%22:%22%7Bcluster%3D%5C%22use1-rdev%5C%22,%20namespace%3D%5C%22rtb-bidder%5C%22%7D%20%7C%3D%20%60health%20service%20monitor%20check%60%20%7C%20json%22,%22queryType%22:%22range%22,%22datasource%22:%7B%22type%22:%22loki%22,%22uid%22:%22P9A714CFC031CC693%22%7D,%22editorMode%22:%22code%22%7D%5D,%22range%22:%7B%22from%22:%22now-30m%22,%22to%22:%22now%22%7D%7D)

Query:

```
{cluster="use1-rdev", namespace="rtb-bidder"} |= `health service monitor check` | json
```


## Architecture Overview

![Alt](docs/diagrams/rtb-bidder.svg)

## Getting Started

rtb-bidder is a service written on [the go programming [language](https://go.dev/) and follows established style guides and best practices](https://go.dev/) 

The service contains the following packages:

- main - Entry point of the project;
- app - Application container, responsible for wiring up and initializing dependencies;
- observability - The observability package contains logger, tracer, and metrics services. The package is responsible for handling metrics collector requests;
serializer - The Serializer package includes available marshal/unmarshal services for JSON and Protobuf data;
- bidder - Bidder package encapsulates business logic handling and processing [Open RTB](https://iabtechlab.com/standards/openrtb/) requests;
- health - The health package is responsible for monitoring the health of the components and handling readiness and liveness requests;
- server - The server is a package responsible for routing and serving web requests;
- handler - The handler package is responsible for setting the handler for the /info endpoint and handling invalid path requests;

### Tech Stack

#### "rtb-bidder" service
The service uses following major frameworks and tools:
- [fasthttp](https://github.com/valyala/fasthttp) - the high-performance web server and client library;
- [wire](https://github.com/google/wire) - the tool for automation of [dependency injection](https://en.wikipedia.org/wiki/Dependency_injection);
- [viper](https://github.com/spf13/viper) - the go configuration package; 
- [ginkgo](https://github.com/onsi/ginkgo) - BDD testing framework;
- [gomega](https://github.com/onsi/gomega) - matcher/assertion library;

#### Docker and k8s:
The service is hosted in the k8s cluster.

- [docker](https://www.docker.com/) - delivery container;
- [helm](https://helm.sh/) - helps manage Kubernetes applications;

### Development environment

[See details in "Development environment"](docs/development.md)

### Databases

The information on the DBs will be available soon

## Run the service 

See [Getting Started document](GETTINGSTARTED.md)

## Deployment

See [Deployment document](docs/deployment.md)