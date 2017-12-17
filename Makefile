DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))

INFOLOG := \033[34m ▸\033[0m
WARNLOG := \033[33m ▸\033[0m
ERROLOG := \033[31m ⨯\033[0m

TEST:="./..."

test: embed
	@echo "$(INFOLOG) Running go tests..."
	@go test -timeout=10m -parallel=10 $(TEST)
.PHONY: tests

# Install the commands.
install:
	@echo "$(INFOLOG) Installing Joy to the path..."
	@go install ./cmd/...
.PHONY: install

embed:
	@go-bindata -o ./internal/bindata/bindata.go -nometadata -pkg bindata ./internal/runtime ./macro/* ./stdlib/**/*
	@gofmt -w ./internal/bindata/
.PHONY: embed	

dom:
	@echo "$(INFOLOG) Generating the DOM..."
	@go run internal/dom/main.go
.PHONY: dom

vdom:
	@echo "$(INFOLOG) Generating the virtual DOM..."
	@go run internal/vdom/main.go
.PHONY: vdom

# Release binaries to GitHub.
release: embed
	@read -p "> What's the new version? (e.g. 0.1.0) " version; \
	echo "$(INFOLOG) Releasing... v$$version"; \
	git tag v$$version; \
	goreleaser -p 1 --rm-dist --config .goreleaser.yml;
.PHONY: release

# Show source statistics.
cloc:
	@cloc -exclude-dir=vendor,node_modules,dom,vdom,chrome .
.PHONY: cloc

# Show to-do items per file.
todo:
	@grep \
		--exclude-dir=vendor \
		--exclude-dir=node_modules \
		--exclude-dir=chrome \
		--exclude-dir=vdom \
		--exclude-dir=dom \
		--exclude=Makefile \
		--text \
		--color \
		-nRo -E ' TODO:.*|SkipNow' .
.PHONY: todo