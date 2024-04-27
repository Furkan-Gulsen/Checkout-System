APP=checkout_app

### General

install:
	@echo "Installing..."
	go mod download

build:
	@echo "Building..."
	go build -o ${APP} cmd/main.go 

run: build
	@echo "Running..."
	./${APP}

dev:
	@echo "Running..."
	go run cmd/main.go 

swag:
	@echo "Generating swagger..."
	swag init -g cmd/main.go -d ./


### Docker
up:
	@echo "Running docker-compose..."
	docker-compose up -d --build

down:
	@echo "Stopping docker-compose..."
	docker-compose down


### Tests
persistance-tests:
	@echo "Running persistance tests..."
	go test -v ./internal/infrastructure/persistence


application-tests:
	@echo "Running application tests..."
	go test -v ./internal/application