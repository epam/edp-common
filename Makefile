.DEFAULT_GOAL:=help

help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

# Run all unit tests
.PHONY: test
# Run tests
test: fmt vet  ## run unit tests
	go test ./... -coverprofile=coverage.out `go list ./...`

.PHONY: fmt
fmt:  ## Run go fmt
	go fmt ./...

.PHONY: vet
vet:  ## Run go vet
	go vet ./...

.PHONY: lint
lint: ## Run go lint
	golangci-lint run
