SHELL=/bin/bash
BINARY_NAME:="ksd"

.PHONY: install
install: ## Install the binary 
	@go build -o ${GOPATH}/bin/${BINARY_NAME}
