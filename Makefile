#
#	Makefile for Volume assignment
#

all: test build

build:
	go build -o bin/server

test:
	go test ./...