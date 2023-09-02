PROJECT := go-htmx 

BOLD := \033[1m
RESET := \033[0m

.DEFAULT_GOAL := help

.PHONY: refresh # Update the dependencies
refresh:
	@echo "$(BOLD)Refreshing dependencies...$(RESET)"
	@go mod tidy

PHONY: build # Build the project
build: refresh
	@echo "${BOLD}Building ${PROJECT}...${RESET}"
	@go build -o bin/${PROJECT} -ldflags '-s -w'

.PHONY: run # Run the project
run: build
	@echo "${BOLD}Running ${PROJECT}...${RESET}"
	@./bin/${PROJECT}

.PHONY: help # Display the help message
help:
	@echo "${BOLD}Available targets:${RESET}"
	@cat Makefile | grep '.PHONY: [a-z]' | sed 's/.PHONY: / /g' | sed 's/ #* / - /g'
