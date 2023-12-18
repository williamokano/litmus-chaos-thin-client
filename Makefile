.PHONY: help
help:
	@echo "Available commands:"
	@echo
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		sort | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: get-project
get-project: ## Gets a project from Litmus Chaos
	@ go run ./cmd/get-project -host ${LITMUS_URL} -token ${LITMUS_TOKEN} -projectId $(projectId)
