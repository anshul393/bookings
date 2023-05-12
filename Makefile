build:
	@go build -o bin/webApp cmd/web/*.go

run: build
	@./bin/webApp


