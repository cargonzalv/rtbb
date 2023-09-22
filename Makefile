export DATE := $(shell date +%Y.%m.%d-%H%M)
export LATEST_COMMIT := $(shell git log --pretty=format:'%h' -n 1)
export GIT_TAG := ${LATEST_COMMIT}
export BRANCH := $(shell git branch |grep -v "no branch"| grep \*|cut -d ' ' -f2)
export BIN_DIR=./bin
export BIN_NAME=srv
export BUILD_VERSION="0.0.1"

export GOIMPORTS_EXECUTABLE := $(shell command -v goimports 2> /dev/null)

# Docker image base
export DOCKER_SERVICE_IMAGE_BASE=adgear-docker.jfrog.io/adgear/rtb-bidder

export BUILT_ON_OS=$(shell uname)
ifeq ($(BRANCH),)
BRANCH := master
endif

export COMMIT_CNT:=$(shell git rev-list HEAD | wc -l | sed 's/ //g' )
export BUILD_NUMBER:=${BRANCH}-${COMMIT_CNT}
export COMPILE_LDFLAGS=-s -X 'github.com/adgear/go-commons/pkg/buildinfo.BuildNumber=${BUILD_NUMBER}' \
                          -X 'github.com/adgear/go-commons/pkg/buildinfo.BuiltOnOs=${BUILT_ON_OS}' \
                          -X 'github.com/adgear/go-commons/pkg/buildinfo.GitDescribeTag=${GIT_TAG}' \
                          -X 'github.com/adgear/go-commons/pkg/buildinfo.DockerTag=${GIT_TAG}'

build_info: check_prereq ## Build the container
	@echo ''
	@echo '---------------------------------------------------------'
	@echo 'BUILT_ON_OS       $(BUILT_ON_OS)'
	@echo 'DATE              $(DATE)'
	@echo 'LATEST_COMMIT     $(LATEST_COMMIT)'
	@echo 'BRANCH            $(BRANCH)'
	@echo 'COMMIT_CNT        $(COMMIT_CNT)'
	@echo 'BUILD_NUMBER      $(BUILD_NUMBER)'
	@echo 'COMPILE_LDFLAGS   $(COMPILE_LDFLAGS)'
	@echo 'PATH              $(PATH)'
	@echo 'GOIMPORTS_EXECUTABLE $(GOIMPORTS_EXECUTABLE)'
	@echo '---------------------------------------------------------'
	@echo ''

####################################################################################################################
##
## help for each task - https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
##
####################################################################################################################
.PHONY: help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

####################################################################################################################
##
## Build of binaries
##
####################################################################################################################
all: rtb-bidder test ## build rtb-bidder and run tests

binaries: rtb-bidder ## build binaries in bin dir

create_dir:
	@mkdir -p $(BIN_DIR)

check_prereq: create_dir

ifndef GOIMPORTS_EXECUTABLE
	go install golang.org/x/tools/cmd/goimports@latest
endif
	$(warning "found goimports")

build-%:
# Given target OS.ARCH format
	CGO_ENABLED=0 GOOS=$(basename $(subst .exe,,$*)) GOARCH=$(subst .,,$(suffix $(subst .exe,,$*))) go build -o $(BIN_DIR)/$(BIN_NAME) -a -ldflags="$(COMPILE_LDFLAGS)" ./cmd/bidder

build-app: create_dir
		go build -mod vendor -o $(BIN_DIR)/$(BIN_NAME) -a -ldflags '$(COMPILE_LDFLAGS)' $(APP_PATH)

rtb-bidder: build_info ## build rtb-bidder binary in bin dir
	@echo "build rtb-bidder openrtb server"
	goimports -w -v ./cmd/bidder
	go mod tidy
	go mod vendor
	make BIN_NAME=srv APP_PATH=./cmd/bidder build-app
	@echo ''
	@echo ''

docker: build-linux.amd64
	docker build . -t ${DOCKER_SERVICE_IMAGE_BASE}:${BUILD_VERSION} --build-arg="BUILD_DATE=${DATE},GIT_HASH=${LATEST_COMMIT},GIT_TAG=${GIT_TAG},GIT_BRANCH=${BRANCH},SERVICE_NAME=rtb-bidder"

####################################################################################################################
##
## Cleanup of binaries
##
####################################################################################################################

clean_binaries: clean_rtb_bidder  ## clean all binaries in bin dir

generate: ## Generate dynamic code sources
	@go generate -v ./...

clean_binary: ## clean binary in bin dir
	rm -f $(BIN_DIR)/$(BIN_NAME)

clean_rtb_bidder: ## clean rtb-bidder 
	make BIN_NAME=srv clean_binary

test: ## run tests
	go test -v ./...

test_with_report: ## run tests with test and coverage reports
	mkdir -p coverage
	go test -v -race ./... -coverprofile=./coverage/coverage.out -json > ./coverage/test-report.json

coverage-app: ## Run tests with coverage
	@go test ./...  -coverprofile coverage.out
	@go tool cover -html=coverage.out -o coverage.html

coverage: generate coverage-app ## Generate dynamic sources and run tests with coverage

fmt: ## run fmt on project
	#go fmt $(PROJ_PATH)/...
	gofmt -s -d -w -l .

doc: ## launch godoc on port 6060
	godoc -http=:6060

deps: ## display deps for project
	go list -f '{{ join .Deps  "\n"}}' . |grep "/" | grep -v $(PROJ_PATH)| grep "\." | sort |uniq

lint: ## run lint on the project
	golangci-lint run --timeout 2m

docker_lint: ## run docker linter
	hadolint Dockerfile

staticcheck: ## run staticcheck on the project
	staticcheck -ignore "$(shell cat .checkignore)" .

vet: ## run go vet on the project
	go vet .

tools: ## install dependent tools for code analysis
	go get -u github.com/gogo/protobuf
	go get -u github.com/gogo/protobuf/proto
	go get -u github.com/gogo/protobuf/jsonpb
	go get -u github.com/gogo/protobuf/protoc-gen-gogo
	go get -u github.com/gogo/protobuf/gogoproto
	go get -u github.com/fzipp/gocyclo #detect cyclomatic complexity
	go get -u golang.org/x/lint/golint
