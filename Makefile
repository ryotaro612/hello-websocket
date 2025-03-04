##@ Run
.PHONY: server
run-server: server/dist/server
	cd server && ./dist/server

##@ Build
server/dist/server: $(shell find server)
	mkdir -p server/dist
	cd server && go build -o dist/server main.go

##@ Clean
.PHONY: clean
clean: ## Clean up all generated files.
	rm -rf server/dist

##@ Help
.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
