# executables
GO=go
SQLC=sqlc
GOLANGCI_LINT=golangci-lint
SQLFLUFF=sqlfluff

# directories
CMD_DIR=cmd
DIST_DIR=dist

# runserver config
RUNSERVER_BIN=${CMD_DIR}/runserver/main.go
RUNSERVER_OUT=${DIST_DIR}/runserver

.PHONY: all build

all: sqlc build

build:
	@echo "build binaries..."
	@${GO} build -o ${RUNSERVER_OUT} ${RUNSERVER_BIN}


sqlc:
	@echo "compiling sql..."
	@${SQLC} generate


test:
	@echo "running tests..."
	@${GO} test ./...

fmt:
	@echo "formatting go files..."
	@${GO} fmt ./...

lint:
	@echo "checking code..."
	@${GO} vet ./...
	@${GOLANGCI_LINT} run ./...

sqlfmt:
	@echo "formatting sql files..."
	@${SQLFLUFF} fix --dialect postgres ./database/**
