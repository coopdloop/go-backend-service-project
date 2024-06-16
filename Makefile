# run:
# 	@templ generate
# 	@go run api/main.go

build:
	@go build -o bin/main cmd/server/main.go

test:
	@go test -v ./...

run: build
	@./bin/main
