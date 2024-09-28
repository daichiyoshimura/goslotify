PROJECT_ROOT = $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

export GOBIN ?= $(PROJECT_ROOT)/bin
export PATH := $(GOBIN):$(PATH)

GOVULNCHECK = $(GOBIN)/govulncheck
BENCH_FLAGS ?= -cpuprofile=cpu.pprof -memprofile=mem.pprof -benchmem

.PHONY: all
all: lint tidy format vuln test

.PHONY: lint
lint:
	@golangci-lint run

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: format
format:
	@goimports -w .

.PHONY: vuln
vuln:
	@govulncheck ./...

.PHONY: test
test:
	@CGO_ENABLED=1 go test -race ./...

.PHONY: bench
bench:
	@CGO_ENABLED=1 go test -bench=.

.PHONY: docs
docs:
	@echo "See http://localhost:6060/pkg/timeslots" &
	@godoc -http=:6060
	