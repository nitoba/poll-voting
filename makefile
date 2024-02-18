.PHONY: default run build tests clean prisma-migrate-dev prisma-generate prisma-deploy env env-test docker-run
# Variables
APP_NAME=poll-voting
APP_ENTRY_POINT=./cmd/main.go

# Tasks
default: run-with-docs

prisma-migrate-dev:
	@go run github.com/steebchen/prisma-client-go migrate dev

prisma-generate:
	@go run github.com/steebchen/prisma-client-go generate

prisma-deploy:
	@go run github.com/steebchen/prisma-client-go migrate deploy

run:
	@go run $(APP_ENTRY_POINT)
run-with-docs:
	@go run github.com/swaggo/swag/cmd/swag@latest init -g $(APP_ENTRY_POINT) -o ./docs
	@go run $(APP_ENTRY_POINT)
docker-run:
	@docker-compose up -d
build:
	@go build -o $(APP_NAME) $(APP_ENTRY_POINT)
tests:
	@go test -v ./...
env:
	@cp .env.example .env
env-test:
	@cp .env.example .env.test
clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs