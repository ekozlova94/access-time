.PHONY: build
build:
	go build cmd/access-time/main.go

.PHONY: run
run:
	go run cmd/access-time/main.go

.PHONY: test
test:
	go test ./...
