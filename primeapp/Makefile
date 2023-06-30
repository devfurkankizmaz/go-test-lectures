build:
	@go build -o bin/primeapp main.go

run: build
	@./bin/primeapp

test:
	@go test -v ./...

cover:
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html

individual_test:
	@go test -v -run "Test_isPrime" .
