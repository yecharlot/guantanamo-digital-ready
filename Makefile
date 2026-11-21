.PHONY: run build tidy

run:
	go run ./cmd/gdr

build:
	go build -o gdr ./cmd/gdr

tidy:
	go mod tidy
