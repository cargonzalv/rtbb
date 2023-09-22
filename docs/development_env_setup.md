# Development environment setup

## Docker desktop

To build and run docker images, install the [Docker desktop](https://www.docker.com/products/docker-desktop/)

## Setup access to k8s cluster

To be able to access the k8s cluster, follow the instructions in the [manual](https://docs.int.adgear.com/ep/k8s/kubectl_setup.html#rancher-cluster-setup-manual)

## GO
A development machine should have the following tools installed:

- [The go language](https://go.dev/doc/install);

## Wire - Automated Initialization Go tool 

Install [wire](https://github.com/golangci/golangci-lint) tool - a code generation tool that automates connecting components using dependency injection.

```shell
go install github.com/google/wire/cmd/wire@latest
```

## Golang linter

Install linter runner [golangci-lint](https://github.com/golangci/golangci-lint)

```shell
brew install golangci-lint
```

## Mock generator

Install mock generator tool [gomock](https://github.com/golang).

```shell
go install github.com/golang/mock/mockgen@v1.6.0
```

## Go version manager

It is recommended to install the version manager for the golang.

- [The Go version management](https://github.com/moovweb/gvm);
- [asdf runtime version manager](https://asdf-vm.com/);

#### IDE

- [Goland](https://www.jetbrains.com/go/promo/?source=google&medium=cpc&campaign=10160684326&term=goland&content=438684701656&gad=1&gclid=Cj0KCQjwoK2mBhDzARIsADGbjeoRBTnSMpvrR5V7fQiQr2UXL0ApQTgnjxo1UNYy8ZQd11aQN3u7DxsaApFnEALw_wcB);
- [VsCode](https://code.visualstudio.com/) with Go plugin;
