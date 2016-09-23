include golang.mk
.DEFAULT_GOAL := test # override default goal set in library makefile

SHELL := /bin/bash
PKG := github.com/Clever/resolve-ip
PKGS := $(shell go list ./... | grep -v /vendor)
EXECUTABLE := $(shell basename $(PKG))
.PHONY: test build vendor $(PKGS) $(SCRIPTS)

$(eval $(call golang-version-check,1.7))

all: test build

# builds every Go script found in scripts/. prefix is to prevent overlap w/ $(PKGS)
SCRIPTS :=  $(addprefix script/, $(shell go list ./... | grep /scripts))
$(SCRIPTS):
	go build -o bin/$(shell basename $@) $(@:script/%=%)

build: $(SCRIPTS)
	go build -o bin/$(EXECUTABLE) $(PKG)

clean:
	rm bin/*

test: $(PKGS)

$(PKGS): golang-test-all-deps
	$(call golang-test-all,$@)

vendor: golang-godep-vendor-deps
	$(call golang-godep-vendor,$(PKGS))

run:
	ENV=staging MONGO_DB=clever AWS_REGION=us-west-1 go run main.go -port=5007
