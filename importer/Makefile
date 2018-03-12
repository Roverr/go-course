test: ## Runs tests on the package
	go test -cover ./...
install: ## Installs dependencies
	dep ensure && cd ./mock && npm install
docker-compose: ## Alias for docker compose up
	docker-compose up
coverProfile: ## Creates cover profile
	go test -coverprofile=coverage.out ./...
coverage: ## Generates coverage html
	$(MAKE) coverProfile && go tool cover -html=coverage.out

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
.DEFAULT_GOAL := help