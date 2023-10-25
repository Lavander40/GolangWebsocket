build:
	go build -o bin/wsserver.exe ./cmd/wsserver/main.go

run:
	./bin/wsserver

.DEFAULT_GOAL := build
