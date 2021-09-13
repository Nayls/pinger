SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

# BASE CONFIG ==========================================================================================================
.SILENT: ;               # no need for @
.ONESHELL: ;             # recipes execute in same shell
.NOTPARALLEL: ;          # wait for this target to finish
.EXPORT_ALL_VARIABLES: ; # send all vars to shell
default: help;           # default target
Makefile: ;              # skip prerequisite discovery

# HELP =================================================================================================================
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

# TASKS ============================================================================================================
all: pre_hook build-linux build-windows ## run pre_hook and build

pre_hook: ## job for deploy git
	@echo "> Go mod tidy"
	@go mod tidy
	@echo "> Go mod download"
	@go mod download
	@echo "> Go mod vendor"
	@go mod vendor

# LINUX BUILD
build-linux: build-linux-i386 build-linux-amd64 ## build linux i386 and amd64

build-linux-i386: ## build binary file
	@echo "> Build linux amd64"
	@GOOS=linux GOARCH=amd64 \
		go build \
		-a \
		-mod vendor \
		-ldflags "-s -w" \
		-installsuffix cgo \
		-trimpath \
		-o bin/pinger-linux-amd64 ./cmd/pinger

build-linux-amd64: ## build binary file
	@echo "> Build linux i386"
	@GOOS=linux GOARCH=386 \
		go build \
		-a \
		-mod vendor \
		-ldflags "-s -w" \
		-installsuffix cgo \
		-trimpath \
		-o bin/pinger-linux-i386 ./cmd/pinger

# WINDOWS BUILD
build-windows: build-windows-i386 build-windows-amd64 ## build windows i386 and amd64

build-windows-amd64: ## build 
	@echo "> Build windows amd64"
	@GOOS=windows GOARCH=amd64 \
		go build \
		-a \
		-mod vendor \
		-ldflags "-s -w" \
		-installsuffix cgo \
		-trimpath \
		-o bin/pinger-windows-amd64 ./cmd/pinger

build-windows-i386: ## build 
	@echo "> Build windows i386"
	@GOOS=windows GOARCH=386 \
		go build \
		-a \
		-mod vendor \
		-ldflags "-s -w" \
		-installsuffix cgo \
		-trimpath \
		-o bin/pinger-windows-i386 ./cmd/pinger