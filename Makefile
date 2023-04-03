include golang.mk
include wag.mk

# Temporarily pin to wag 6.4.5 until after migrated to go mod and Go 1.16
WAG_VERSION := latest

.PHONY: all test build run
SHELL := /bin/bash
APP_NAME ?= resolve-ip
EXECUTABLE = $(APP_NAME)
PKG = github.com/Clever/$(APP_NAME)/v4
PKGS := $(shell go list ./... | grep -v /vendor | grep -v /gen-go)

$(eval $(call golang-version-check,1.16))

all: test build

test: $(PKGS)
$(PKGS): golang-test-all-strict-deps
	$(call golang-test-all-strict,$@)

build:
	$(call golang-build,$(PKG),$(EXECUTABLE))

run: build
	bin/$(EXECUTABLE)

generate: wag-generate-deps
	$(call wag-generate-mod,./swagger.yml)
	go mod vendor




install_deps:
	go mod vendor
