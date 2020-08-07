.PHONY: build
build:
	@go build -o gpxtool cmd/gpxtool/main.go

.PHONY: clean
clean:
	rm -f gpxtool

.PHONY: lint
lint:
	@golangci-lint run --enable-all