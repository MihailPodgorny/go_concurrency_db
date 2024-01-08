# https://github.com/SchwarzIT/go-template/blob/main/Makefile

SHELL=/bin/bash -e -o pipefail
PWD = $(shell pwd)
LINTER_VERSION = v1.55.1

out:
	@mkdir -pv "$(@)"


GO_BUILD = mkdir -pv "$(@)" && go build -ldflags="-w -s" -o "$(@)" ./...
.PHONY: out/bin
out/bin:
	$(GO_BUILD)

fmt: ## Formats all code with go fmt
	@go fmt ./...
	@go imports ./...

tidy:
	go mod tidy
	go mod vendor

run: fmt ## Run a controller from your host
	@go run ./main.go

test: ## Runs all tests
	@go test ./...

coverage: out/report.json ## Displays coverage per func on cli
	go tool cover -func=out/cover.out

html-coverage: out/report.json ## Displays the coverage results in the browser
	go tool cover -html=out/cover.out

test-coverage: out ## Creates a test coverage profile
	go test -v -cover ./... -coverprofile out/coverage.out -coverpkg ./...
	go tool cover -func out/coverage.out -o out/coverage.out

lint: fmt download ## Lints all code with golangci-lint
	@go run -v github.com/golangci/golangci-lint/cmd/golangci-lint@$(LINTER_VERSION) run


govulncheck: ## Vulnerability detection using govulncheck
	@go run golang.org/x/vuln/cmd/govulncheck ./...


###############################################################

test-build: ## Tests whether the code compiles
	@go build -o /dev/null ./...



generate: ## Generates files
	@go run cmd/options2md/main.go -o docs/options.md
	@go run github.com/nix-community/gomod2nix@latest --outdir nix

download: ## Downloads the dependencies
	@go mod download


tools:
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@$(LINTER_VERSION)
	go get golang.org/x/tools/cmd/goimports@latest
	golang.org/x/vuln/cmd/govulncheck@latest

clean: clean-test-project ## Cleans up everything
	@rm -rf bin out


git-hooks:
	@git config --local core.hooksPath .githooks/

.PHONY: release
release:  ## Create a new release version
	@./hack/release.sh

help:
	@echo 'Usage: make <OPTIONS> ... <TARGETS>'
	@echo ''
	@echo 'Available targets are:'
	@echo ''
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@echo ''