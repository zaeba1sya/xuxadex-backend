#!/usr/bin/make -f
.DEFAULT_GOAL := help

help:
	@echo " "
	@echo "XuXaDex Backend MVP Project"
	@echo " "
	@echo "Usage:"
	@echo " "
	@echo "> make run - Run Application"
	@echo " "
	@echo "> make dev - Run Application in dev mode"
	@echo " "
	@echo "> make db-status - Check database status"
	@echo " "
	@echo "> make db-up - Create database"
	@echo " "
	@echo "> make db-down - Drop database"
	@echo " "
	@echo "> make lint - Lint code via GolangCI-Lint"
	@echo " "
	@echo "> make swagger - Generate Swagger documentation"
	@echo " "
	@echo "Dependencies:"
	@echo " "
	@echo "Goose (go install github.com/pressly/goose/v3/cmd/goose@latest)"
	@echo "Swag (go install github.com/swaggo/swag/cmd/swag@latest)"
	@echo "Air (go install github.com/cosmtrek/air@latest)"

run:
	@echo " > Running Application"
	go run ./cmd/main/main.go

dev:
	@echo " > Running Application in dev mode"
	air -c .air.toml

db-status:
	@echo " > Checking database status"
	go run ./cmd/migrations/migrations.go --status

db-up:
	@echo " > Creating database"
	go run ./cmd/migrations/migrations.go --up

db-down:
	@echo " > Dropping database"
	go run ./cmd/migrations/migrations.go --down

ccompile:
	@echo " > Compiling smart contract"
	solc --abi --bin --overwrite -o ./pkg/blockchain/contracts/gen ./pkg/blockchain/contracts/src/Foo.sol
	abigen --pkg contracts --abi ./pkg/blockchain/contracts/gen/Foo.abi --bin ./pkg/blockchain/contracts/gen/Foo.bin --out ./pkg/blockchain/contracts/Foo.go

swagger:
	@echo " > Generating Swagger documentation"
	swag init --parseDependency -g ./cmd/main/main.go

lint:
	@echo " > Linting code via GolangCI-Lint"
	golangci-lint run --go 1.23
