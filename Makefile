include golang.mk
include wag.mk

WAG_VERSION := latest

.PHONY: all test build run
SHELL := /bin/bash
APP_NAME ?= resolve-ip
EXECUTABLE = $(APP_NAME)
PKG = github.com/Clever/$(APP_NAME)
PKGS := $(shell go list ./... | grep -v /vendor | grep -v /gen-go)

$(eval $(call golang-version-check,1.9))

all: test build

test: $(PKGS)
$(PKGS): golang-test-all-strict-deps
	$(call golang-test-all-strict,$@)

build:
	CGO_ENABLED=0 go build -installsuffix cgo -o build/$(EXECUTABLE) $(PKG)

run: build
	build/$(EXECUTABLE)

generate: wag-generate-deps
	$(call wag-generate,./swagger.yml,$(PKG))




install_deps: golang-dep-vendor-deps
	$(call golang-dep-vendor)