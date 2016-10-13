include swagger.mk
include golang.mk

.PHONY: all test build run
SHELL := /bin/bash
APP_NAME ?= resolve-ip
EXECUTABLE = $(APP_NAME)
PKG = github.com/Clever/$(APP_NAME)
PKGS := $(shell go list ./... | grep -v /vendor | grep -v /gen-go)

SWAGGER_CONFIG := swagger.yml
SWAGGER_CLIENT_NPM_PACKAGE_NAME := @clever/resolve-ip
SWAGGER_CLIENT_NPM_PACKAGE_VERSION := 0.1.0
SWAGGER_CLIENT_NPM_PACKAGE_MODULE_NAME := resolve-ip

$(eval $(call golang-version-check,1.7))

all: test build

test: $(PKGS)
$(PKGS): golang-test-all-strict-deps
	$(call golang-test-all-strict,$@)

build:
	go build -o build/$(EXECUTABLE) $(PKG)

run: build
	build/$(EXECUTABLE)

validate: swagger-validate-deps
	$(call swagger-validate,$(SWAGGER_CONFIG))

generate: validate swagger-generate-go-deps swagger-generate-javascript-client-deps
	$(call swagger-generate-go,$(SWAGGER_CONFIG),$(PKG),$(PKG)/gen-go)
	$(call swagger-generate-javascript-client,$(SWAGGER_CONFIG),$(SWAGGER_CLIENT_NPM_PACKAGE_NAME),$(SWAGGER_CLIENT_NPM_PACKAGE_VERSION),$(SWAGGER_CLIENT_NPM_PACKAGE_MODULE_NAME))

$(GOPATH)/bin/glide:
	@go get github.com/Masterminds/glide

install_deps: $(GOPATH)/bin/glide
	@$(GOPATH)/bin/glide install
