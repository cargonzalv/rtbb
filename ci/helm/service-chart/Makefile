helm_bin = $(shell command -v helm3 2> /dev/null || command -v helm)
srv=rtb-bidder
namespace=rtb-bidder
THIS_FILE := $(lastword $(MAKEFILE_LIST))
export DYNAMIC_NAMESPACE="not-set"
export DYNAMIC_HOST="rtb-bidder"
export IMAGE_TAG="0.0.0"
export PR_FILE="not-set"

dry-run-dev: deps dry-run-chart-dev
dry-run-prod: deps dry-run-chart-prod
install-dev: deps install-chart-dev
install-prod: deps install-chart-prod

deps:
	$(helm_bin) dep update

dry-run-chart-dev:
	$(helm_bin) deploy --dry-run --name $(srv) -n $(namespace) --kube-context use1-rdev -f values.yaml -f dev.yaml .

dry-run-chart-prod:
	$(helm_bin) deploy --dry-run --name $(srv) -n $(namespace) --kube-context use1-rprod -f values.yaml -f prod.yaml .

install-chart-dev:
	$(helm_bin) deploy --name $(srv)  -n $(namespace) --kube-context use1-rdev -f values.yaml -f dev.yaml . --wet

install-chart-prod:
	$(helm_bin) deploy --name $(srv) -n $(namespace) --kube-context use1-rprod -f values.yaml -f prod.yaml . --wet

install-chart-dev:
	$(helm_bin) deploy --name $(srv) -n $(namespace) --kube-context use1-rdev -f values.yaml -f dev.yaml . --wet

install-chart-prod:
	$(helm_bin) deploy --name $(srv) -n $(namespace) --kube-context use1-rprod -f values.yaml -f prod.yaml . --wet

delete-dev:
	$(helm_bin) delete $(srv) -n $(namespace) --kube-context use1-rdev

delete-prod:
	$(helm_bin) delete $(srv) -n $(namespace) --kube-context use1-rprod

create-ns-info:
	sed 's/"kubernetes.io\/metadata.name": "namespace-placeholder"/"kubernetes.io\/metadata.name": "$(DYNAMIC_NAMESPACE)"/;s/"name": "namespace-placeholder"/"name": "$(DYNAMIC_NAMESPACE)"/'  pr-namespace.json >> namespace.json

pr-value-file:
    sed '/^image:/{n;s/tag:.*/tag: "$(IMAGE_TAG)"/;}' $(PR_FILE) >> env.yaml