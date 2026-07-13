default:
    @just --list

start:
    cd src && go run main.go

test:
    cd test && go test

fmt:
    go fmt ./...
