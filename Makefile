.PHONY: golangci
golangci:
	golangci-lint run ./...

build:
	go build .

.PHONY: test
test:
	@go test -v ./... | sed ''/PASS/s//$$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$$(printf "\033[31mFAIL\033[0m")/''

.PHONY: test
test-report:
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out
