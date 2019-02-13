#--- Variables ---
EMAIL_ADDRESS=1
EMAIL_PASSWORD=2

#--- Help ---
help:
	@echo 
	@echo Makefile targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' Makefile | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@echo
.PHONY: help

#--- deps ---
deps: ## Get and update all dependencies
	go get -u=patch ./...
.PHONY: deps

#--- build-linux ---
build-linux: ## Build a Linux executable
	@GOOS=linux go build -o bootchecker -ldflags "-X github.com/retgits/bootchecker/cmd.emailAddress=${EMAIL_ADDRESS} -X github.com/retgits/bootchecker/cmd.emailPassword=${EMAIL_PASSWORD}" main.go

#--- build-macos ---
build-macos: ## Build a macOS executable
	@GOOS=darwin go build -o bootchecker -ldflags "-X github.com/retgits/bootchecker/cmd.emailAddress=${EMAIL_ADDRESS} -X github.com/retgits/bootchecker/cmd.emailPassword=${EMAIL_PASSWORD}" main.go