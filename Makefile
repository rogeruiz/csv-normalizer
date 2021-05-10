.DEFAULT_GOAL := build

BIN_FILE=normalizer
build_dir = tmp

$(build_dir):
	@echo "Folder ./tmp doesn't exist."
	mkdir -vp $@

.PHONY: build
build: $(build_dir) ## Builds the normalizer command-line tool.
	@go build -o "tmp/${BIN_FILE}"

.PHONY: clean
clean: ## Runs clean-up.
	go clean
	@rm -rvf tmp

.PHONY: run
run: build ## Runs the normalizer command-line tool.
	tmp/${BIN_FILE} < data/sample.csv

.PHONY: help
help: ## Outputs this help message.
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
