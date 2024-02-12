.PHONY: default run build tests docs clean prisma-migrate-dev prisma-generate prisma-studio
# Variables
APP_NAME=poll-voting
APP_ENTRY_POINT=./cmd/main.go

# Tasks
default: run

prisma-migrate-dev:
	@go run github.com/steebchen/prisma-client-go migrate dev

prisma-generate:
	@go run github.com/steebchen/prisma-client-go generate

prisma-studio:
	@go run github.com/steebchen/prisma-client-go studio

run:
	@go run $(APP_ENTRY_POINT)
run-with-docs:
	@swag init
	@go run $(APP_ENTRY_POINT)
build:
	@go build -o $(APP_NAME) $(APP_ENTRY_POINT)
tests:
	@go test -v ./...
docs:
	@swag init
clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs