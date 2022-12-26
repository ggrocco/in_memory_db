.PHONY: test
test:
	go test -cover -race ./...

.PHONY: run
run:
	go run ./cmd/main.go