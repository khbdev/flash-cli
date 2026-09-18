BINARY  := flash
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
LDFLAGS := -s -w -X github.com/khbdev/flash-cli/internal/cli.Version=$(VERSION)
PREFIX  ?= /usr/local

.PHONY: build test lint fmt install clean release

build: ## Binarni ./bin ga yig'ish
	go build -trimpath -ldflags '$(LDFLAGS)' -o bin/$(BINARY) ./cmd/flash

test: ## Testlarni race detector bilan ishga tushirish
	go test -race ./...

lint: ## vet + formatlash tekshiruvi
	go vet ./...
	@test -z "$$(gofmt -l . )" || { echo "gofmt kerak:"; gofmt -l .; exit 1; }

fmt: ## Kodni formatlash
	gofmt -w .

install: build ## $(PREFIX)/bin ga o'rnatish
	install -m 0755 bin/$(BINARY) $(PREFIX)/bin/$(BINARY)

release: ## Linux/macOS uchun binarlar
	@for target in linux/amd64 linux/arm64 darwin/amd64 darwin/arm64; do \
		os=$${target%/*}; arch=$${target#*/}; \
		echo "→ $$os/$$arch"; \
		GOOS=$$os GOARCH=$$arch go build -trimpath -ldflags '$(LDFLAGS)' \
			-o dist/$(BINARY)-$$os-$$arch ./cmd/flash; \
	done

clean:
	rm -rf bin dist
