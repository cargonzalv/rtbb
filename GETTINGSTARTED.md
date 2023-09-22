# Getting Started

Setup development machine according to the [development environment setup istructions](docs/development.md).

### Build and Run

Run the following commands to build the service from the source code:

```shell
# get imported packages
go get ./...

# build binary(binary `bidder` file will be located in `bin` dir)
go build ./cmd/bidder -o ./bin/
```

Run the binary:

```shell
# run the service binary
./bidder
```

Run in docker container:

 ```shell
 docker run -it -p 8085:8085 rtb-bidder:0.0.1
 ```
 or with docker-compose

 ```shell
 docker-compose up
 ```

 To see information how to build the container, see [deployment doc](docs/deployment.md#docker)

Output build and version info:

```shell
./bidder --version
```

You can run the service from source code by running the following command:

```shell
# get imported packages
go get ./...

# run the service 
go run ./cmd/bidder
```

### Testing

To execute unit tests, run the command:

```shell
go test ./...
```

Run following command to generate test coverage report:

```shell
# run test
go test ./... -race -coverprofile=coverage.out -covermode=atomic

# generate report
go tool cover -html coverage.out -o coverage.html
```

In case, if you need to exclude auto generated files from coverage, run:

```shell
# run test
go test ./... -race -coverprofile=coverage.out.tmp -covermode=atomic

# filter data related to the generated files
grep -v ".gen.go" > coverage.out < coverage.out.tmp

# generate report
go tool cover -html coverage.out -o coverage.html
```
