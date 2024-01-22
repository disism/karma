# -*- coding: utf-8 -*-
# SET YOUR ORGANIZATION'S ADDRESS, WHICH IS ALSO USED AS THE GO PACKAGE URL.
ORG = "github.com/disism"

FSPATH=$(PWD)
GIT_REPO_URL = $(ORG)
REPO_NAME = $(shell basename `pwd`)
REPO_URL = "$(GIT_REPO_URL)/$(REPO_NAME)"

# OS
OS := $(shell uname -s)
ARCH := $(shell uname -m)
USER_HOME := $(shell echo ${HOME})

# VERSION
VERSION_FILE := VERSION
VERSION := $(shell cat ./$(VERSION_FILE))

# GO
GOPATH := $(shell go env GOPATH)

ifeq ($(VERSION),)
        VERSION := $(shell git describe --tags --abbrev=0)
endif

# MKDIR
MAKE_FSPATH=$(FSPATH)/make
MK_FILES := $(wildcard $(MAKE_FSPATH)/*.mk)

.PHONY: print mod ent run build

print:
	@echo
	@echo "REPO NAME: $(REPO_NAME)"
	@echo "REPO URL: $(REPO_URL)"
	@echo "FSPATH: $(FSPATH)"
	@echo "MAKE_FSPATH: $(MAKE_FSPATH)"
	@echo "VERSION: $(VERSION)"
	@echo ""

	@echo "OS: $(OS)"
	@echo "ARCH: $(ARCH)"
	@echo "GO_PATH: $(GOPATH)"
	@echo "USER_HOME: $(USER_HOME)"
	@echo ""

	@echo "Usage: "
	@echo "  make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  print        - Print the dev environment and help."
	@echo "  mod          - Generate go.mod and run go mod tidy."
	@echo "  ent          - Generate ent code. ref: https://entgo.io/"
	@echo "  run          - Run the server."
	@echo "  migrations   - version migrations."
	@echo ""

mod:
	@rm -f go.mod go.sum
	@go mod init $(GIT_REPO_URL)/$(REPO_NAME)
	@go mod tidy

ent:
	go generate $(FSPATH)/ent
	@echo "for version migrations, please use"
	@echo ""
	@echo "    make migrations	"

run:
	go run main.go run

build:
	#GOOS=linux GOARCH=amd64 go build -o saikan main.go
	go build -o saikan main.go

include $(MK_FILES)
