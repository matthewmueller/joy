test:
	@go test -cover ./...
.PHONY: test

# Install the commands.
install:
	@go install ./cmd/...
.PHONY: install