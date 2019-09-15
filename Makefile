# kernel-style V=1 build verbosity
ifeq ("$(origin V)", "command line")
       BUILD_VERBOSE = $(V)
endif

ifeq ($(BUILD_VERBOSE),1)
       Q =
else
       Q = @
endif

VERSION = $(shell git describe --dirty --tags --always)
SERVICE_NAME = orders

export CGO_ENABLED:=0

all: verify build

lint:
		$(Q)echo "linting...."
		$(Q)GO111MODULE=off go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
		$(Q)golangci-lint run -E gofmt -E golint -E goconst -E gocritic -E golint -E gosec -E maligned -E nakedret -E prealloc -E unconvert -E gocyclo -E scopelint -E goimports
		$(Q)echo linting OK

test:
		$(Q)echo "unit testing...."
		$(Q)go test ./pkg/...

verify: lint test

api:
		$(Q)protoc --go_out=plugins=grpc:. pkg/spec/v1/*.proto	

prepare: fmt generate verify

clean:
		$(Q)rm -rf build

fmt:
		$(Q)gofmt -w .

generate:
		$(Q)go get github.com/golang/mock/gomock
		$(Q)go install github.com/golang/mock/mockgen
		$(Q)go generate ./...

build:
		$(Q)$(GOARGS) go build -gcflags "all=-trimpath=${GOPATH}" -asmflags "all=-trimpath=${GOPATH}" -o ./artifacts/$(SERVICE_NAME) ./$(SERVICE_NAME).go
