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

help: #!# Display this help screen
	@awk 'BEGIN {FS = ":.*#!#"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-].+:.*?#!#/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^#!#@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

# TASKS ============================================================================================================
all: pre-run build-linux build-windows build-darwin #!# run job pre-run, build-linux, build-windows, build-darwin

pre-run: ## job for deploy git
	@echo "> Go mod tidy"
	@go mod tidy
	@echo "> Go mod download"
	@go mod download
	@echo "> Go mod vendor"
	@go mod vendor
	@echo "> Generate cli docs"
	@go run main.go generate cli

# LINUX BUILD
build-linux: build-linux-386 build-linux-amd64 build-linux-arm build-linux-arm64 #!# build linux 386, amd64, arm, arm64

build-linux-386: ## build binary file linux-386
	@echo "> Build linux 386"
	@GOOS=linux GOARCH=386 \
		go build \
		-a \
		-mod vendor \
		-ldflags "-s -w" \
		-installsuffix cgo \
		-trimpath \
		-o bin/linux/pinger-linux-386 ./main.go

build-linux-amd64: ## build binary file linux-amd64
	@echo "> Build linux amd64"
	@GOOS=linux GOARCH=amd64 \
		go build \
		-a \
		-mod vendor \
		-ldflags "-s -w" \
		-installsuffix cgo \
		-trimpath \
		-o bin/linux/pinger-linux-amd64 ./main.go

build-linux-arm: ## build binary file linux-arm
	@echo "> Build linux arm"
	@GOOS=linux GOARCH=arm \
		go build \
		-a \
		-mod vendor \
		-ldflags "-s -w" \
		-installsuffix cgo \
		-trimpath \
		-o bin/linux/pinger-linux-arm ./main.go

build-linux-arm64: ## build binary file linux-arm64
	@echo "> Build linux arm64"
	@GOOS=linux GOARCH=arm64 \
		go build \
		-a \
		-mod vendor \
		-ldflags "-s -w" \
		-installsuffix cgo \
		-trimpath \
		-o bin/linux/pinger-linux-arm64 ./main.go

# WINDOWS BUILD
build-windows: build-windows-386 build-windows-amd64 build-windows-arm #!# build windows 386, amd64, arm

build-windows-386: ## build binary file windows-386
	@echo "> Build windows 386"
	@GOOS=windows GOARCH=386 \
		go build \
		-a \
		-mod vendor \
		-ldflags "-s -w" \
		-installsuffix cgo \
		-trimpath \
		-o bin/windows/pinger-windows-386.exe ./main.go

build-windows-amd64: ## build binary file windows-amd64
	@echo "> Build windows amd64"
	@GOOS=windows GOARCH=amd64 \
		go build \
		-a \
		-mod vendor \
		-ldflags "-s -w" \
		-installsuffix cgo \
		-trimpath \
		-o bin/windows/pinger-windows-amd64.exe ./main.go

build-windows-arm: ## build binary file windows-arm
	@echo "> Build windows arm"
	@GOOS=windows GOARCH=arm \
		go build \
		-a \
		-mod vendor \
		-ldflags "-s -w" \
		-installsuffix cgo \
		-trimpath \
		-o bin/windows/pinger-windows-arm ./main.go

# MACOS BUILD
build-darwin: build-darwin-amd64 build-darwin-arm64 #!# build darwin amd64, arm64

build-darwin-amd64: ## build binary file darwin-amd64
	@echo "> Build darwin amd64"
	@GOOS=darwin GOARCH=amd64 \
		go build \
		-a \
		-mod vendor \
		-ldflags "-s -w" \
		-installsuffix cgo \
		-trimpath \
		-o bin/darwin/pinger-darwin-amd64 ./main.go

build-darwin-arm64: ## build binary file darwin-arm64
	@echo "> Build darwin arm64"
	@GOOS=darwin GOARCH=arm64 \
		go build \
		-a \
		-mod vendor \
		-ldflags "-s -w" \
		-installsuffix cgo \
		-trimpath \
		-o bin/darwin/pinger-darwin-arm64 ./main.go