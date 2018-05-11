#
# Synse Emulator Plugin
#

PLUGIN_NAME		:= emulator
PLUGIN_VERSION	:= 1.0.1

IMAGE_NAME		:= vaporio/emulator-plugin

GIT_COMMIT ?= $(shell git rev-parse --short HEAD 2> /dev/null || true)
GIT_TAG    ?= $(shell git describe --tags 2> /dev/null || true)


PKG_CTX    := main
BUILD_DATE := $(shell date -u +%Y-%m-%dT%T 2> /dev/null)
#GIT_COMMIT := $(shell git rev-parse --short HEAD 2> /dev/null || true)
#GIT_TAG    := $(shell git describe --tags 2> /dev/null || true)
GO_VERSION := $(shell go version | awk '{ print $$3 }')

LDFLAGS="-w \
	-X ${PKG_CTX}.BuildDate=${BUILD_DATE} \
	-X ${PKG_CTX}.GitCommit=${GIT_COMMIT} \
	-X ${PKG_CTX}.GitTag=${GIT_TAG} \
	-X ${PKG_CTX}.GoVersion=${GO_VERSION} \
	-X ${PKG_CTX}.VersionString=${PLUGIN_VERSION} \
	${LDFLAGS}"

HAS_LINT := $(shell which gometalinter)
HAS_DEP  := $(shell which dep)


.PHONY: build
build:  ## Build the plugin Go binary
	go build -ldflags ${LDFLAGS} -o build/emulator

.PHONY: build-linux
build-linux:  ## Build the plugin for linux amd64
	GOOS=linux GOARCH=amd64 go build -ldflags ${LDFLAGS} -o build/emulator .

.PHONY: ci
ci:  ## Run CI checks locally (build, lint)
	@$(MAKE) build lint

.PHONY: clean
clean:  ## Remove temporary files
	go clean -v

.PHONY: dep
dep:  ## Ensure and prune dependencies
ifndef HAS_DEP
	go get -u github.com/golang/dep/cmd/dep
endif
	dep ensure -v

.PHONY: docker
docker:  ## Build the docker image
	docker build -f Dockerfile \
		-t $(IMAGE_NAME):latest \
		-t $(IMAGE_NAME):$(PLUGIN_VERSION) .

.PHONY: fmt
fmt:  ## Run goimports on all go files
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do goimports -w "$$file"; done

.PHONY: github-tag
github-tag:  ## Create and push a tag with the current plugin version
	git tag -a ${PLUGIN_VERSION} -m "${PLUGIN_NAME} version ${PLUGIN_VERSION}"
	git push -u origin ${PLUGIN_VERSION}

.PHONY: lint
lint:  ## Lint project source files
ifndef HAS_LINT
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install
endif
	@ # disable gotype: https://github.com/alecthomas/gometalinter/issues/40
	gometalinter ./... --tests --vendor --deadline=5m \
		--disable=gotype --disable=gocyclo

.PHONY: setup
setup:  ## Install the build and development dependencies and set up vendoring
	go get -u github.com/alecthomas/gometalinter
	go get -u github.com/golang/dep/cmd/dep
	gometalinter --install
ifeq (,$(wildcard ./Gopkg.toml))
	dep init
endif
	@$(MAKE) dep

.PHONY: version
version:  ## Print the version of the plugin
	@echo "$(PLUGIN_VERSION)"

.PHONY: help
help:  ## Print usage information
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort

.DEFAULT_GOAL := help
