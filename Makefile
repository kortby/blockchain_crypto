build:
	@go build -o bin/blockchaincrypto

run: build
	@./bin/docker

test:
	@go test -v ./...