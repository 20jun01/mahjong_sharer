# commands
GOCMD=go
GOMOCKHANDLER=gomockhandler
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOINSTALL=$(GOCMD) install
GOMOD=$(GOCMD) mod
GOGEN=$(GOCMD) generate
GOFMT=$(GOCMD)fmt
GOIMPORTS=$(GOCMD)imports
GOMOCKGEN=$(GOMOCKHANDLER) mockgen

# air
AIR_COMMAND=air
AIR_OPTION=-c .air.toml

# option variables
GOTEST_FLAGS := -v -cover 

# variables
BINARY_NAME := mahjong_backend

BINDIR:=build

BINARIES := $(BINDIR)/

MOCKGEN := github.com/golang/mock/mockgen
MOCKHANDLER := github.com/sanposhiho/gomockhandler

TEST_RESULT_DIR := ./tmp/test-results

TEST_RESULT_FILE := test-results.xml

# files
BINARY := $(BINDIR)/$(BINARY_NAME)
BINARY_UNIX := $(BINDIR)/$(BINARY_NAME_UNIX)

GOFILES=$(wildcard *.go **/*.go)

MAIN := main.go

# command with variables
INSTALL_MOCKGEN := $(GOINSTALL) $(MOCKGEN)@latest
INSTALL_MOCKHANDLER := $(GOINSTALL) $(MOCKHANDLER)@latest

.PHONY: ${shell egrep -o ^[a-zA-Z_-]+: ./Makefile | sed 's/://'}

all: clean mod build run

clean:
	@rm -f ${BINARY}
	@${GOCLEAN}

mod:
	@${GOMOD} tidy

build: ${GOFILES}
	@${GOBUILD} -o ${BINARY} -v ${MAIN}

run: build
	@${BINARY}

run_hot:
	@${AIR_COMMAND} ${AIR_OPTION}

vet:
	@${GOCMD} vet ./...

fmt:
	@${GOCMD} fmt ./...

gofmt:
	@${GOFMT} -w .

imports:
	@${GOIMPORTS} -w .

mockgen:
	@$(INSTALL_MOCKGEN)
	@$(INSTALL_MOCKHANDLER)
	@${GOMOCKGEN}

sql-gen:
	@cd db_config; sqlc generate

test: ${GOFILES} mockgen fmt
	@${GOTEST} ${GOTEST_FLAGS} ./...

test_with_junit: mockgen fmt
	@mkdir -p ${TEST_RESULT_DIR}
	@${GOTEST} ${GOTEST_FLAGS} ./...

linter:
	@golangci-lint run

golang_ci:
	@golangci-lint run -v

golang_ci_with_docker:
	@docker run -t --rm -v $(CURDIR):/app -w /app golangci/golangci-lint:v1.53.3 golangci-lint run -v