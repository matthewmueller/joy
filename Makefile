TEST:="./"

test:
	@echo "==> Running go tests..."
	@go list $(TEST) \
		| grep -v "/vendor/" \
		| xargs -n1 go test -timeout=5m -parallel=10 $(TESTARGS)
.PHONY: tests

jsx:
	@echo "==> Running jsx example..."
	@go run cmd/golly/golly.go build ./testdata/49-jsx/
.PHONY: jsx

hn:
	@echo "==> Running hackernews example..."
	@go run cmd/golly/golly.go serve ./_examples/hn
.PHONY: hn

# Install the commands.
install:
	@go install ./cmd/...
.PHONY: install

dom2:
	# @go run internal/dom/main.go
	# @go generate ./internal/gendom/main.go
	@go run cmd/golly/golly.go build ./_examples/dom2
	# @go run _examples/dom/dom.go
.PHONY: dom2

dom:
	# @go run internal/dom/main.go
	# @go generate ./internal/gendom/main.go
	@go run cmd/golly/golly.go build ./_examples/dom
	# @go run _examples/dom/dom.go
.PHONY: dom