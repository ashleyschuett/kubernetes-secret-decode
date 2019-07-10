GO111MODULE=on
SHELL=/bin/bash
BINARY_NAME:=kubectl-ksd
GOPATH:=${HOME}/go

.PHONY: install
install: ## Install the binary 
	@go build -o "${GOPATH}/bin/${BINARY_NAME}"
