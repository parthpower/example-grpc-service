TOP_DIR := $(PWD)
BUILD_DIR := $(TOP_DIR)/.build

GO     := $(shell which go)
GO_FLAGS := -a -tags netgo -ldflags '-w -extldflags "-static"'
PATH := $(PATH):$(BUILD_DIR)/protoc/bin:$(shell $(GO) env GOPATH)/bin
PROTOC := $(BUILD_DIR)/protoc/bin/protoc
PROTOC_INCLUDE := -I$(BUILD_DIR)/protoc/include -I$(TOP_DIR)/proto
PROTOC_GEN_GO := $(shell which protoc-gen-go)
PROTOC_VERSION := 3.12.4

DOCKER := $(shell which docker)

DOCKER_VERSION ?= latest
DOCKER_IMAGE_OUT := parthpower.me/service:$(DOCKER_VERSION)

all: build

# To build protobuf stubs if required
.PHONY: protoc
protoc: $(PROTOC)
$(PROTOC):
	mkdir -p $(BUILD_DIR)/protoc && \
	cd $(BUILD_DIR)/protoc && \
	curl -Lfso protoc.zip \
		https://github.com/protocolbuffers/protobuf/releases/download/v$(PROTOC_VERSION)/protoc-$(PROTOC_VERSION)-linux-x86_64.zip && \
	unzip -y protoc.zip

.PHONY: download
download:
	@echo Download go.mod dependencies
	@$(GO) mod download

.PHONY: install-tools
install-tools: download
	@echo Installing tools from tools.go
	@cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

protostub-a: protoc install-tools
	mkdir -p $(TOP_DIR)/pkg/api/serviceav1
	$(PROTOC) $(PROTOC_INCLUDE) \
		--go_out=$(TOP_DIR)/pkg/api/serviceav1 \
		--go-grpc_out=$(TOP_DIR)/pkg/api/serviceav1 \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		servicea.proto

protostub-b: protoc install-tools
	mkdir -p $(TOP_DIR)/pkg/api/servicebv1
	$(PROTOC) $(PROTOC_INCLUDE) \
		--go_out=$(TOP_DIR)/pkg/api/servicebv1 \
		--go-grpc_out=$(TOP_DIR)/pkg/api/servicebv1 \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		serviceb.proto


.PHONY: build
build: docker-build

.PHONY: go-build
go-build: install-tools
	$(GO) fmt
	$(GO) vet
	golint
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO) build $(GO_FLAGS) -o svc

.PHONY: docker-build
docker-build: go-build
	$(DOCKER) build -t $(DOCKER_IMAGE_OUT) $(TOP_DIR)

.PHONY: docker-tarball
docker-tarball: docker-build
	mkdir -p $(BUILD_DIR)
	$(DOCKER) save $(DOCKER_IMAGE_OUT) > $(BUILD_DIR)/svc.tar
	# sudo k3s ctr image import svc.tar