TEST:="./"

test:
	@echo "==> Running go tests..."
	@go list $(TEST) \
		| grep -v "/vendor/" \
		| xargs -n1 go test -timeout=5m -parallel=10 $(TESTARGS)
.PHONY: tests

hn:
	@echo "==> Running hackernews example..."
	@go run cmd/golly/golly.go serve ./testdata/56-hn-preact/

# Install the commands.
install:
	@go install ./cmd/...
.PHONY: install