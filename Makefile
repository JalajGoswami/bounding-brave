SHELL := /usr/local/bin/nu
export GOOS := windows
export GOARCH := amd64

default: build

build:
	go build -o tmp/main.exe .

run: build
	./tmp/main.exe
