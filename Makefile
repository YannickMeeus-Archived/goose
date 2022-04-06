help: ## Show all help information
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m\033[0m\n"} /^[$$()% 0-9a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: dev
dev: ## Start development server, watches for changes and restarts process
	watchexec -e go go run . 

##  Filtering out packages without tests, as per https://github.com/gotestyourself/gotestsum/issues/228
.PHONY: test
test: ## Start test server, watches for changes and restarts process
	@gotestsum --watch --format pkgname-and-test-fails  --no-color=false | grep -v '∅'

.PHONY: build
build: ## Build a package
	@go build -o ./dist/bin/

.PHONY: install-dev-deps
install-dev-deps: ## Install implicit dependencies for development
	go install gotest.tools/gotestsum@latest
	