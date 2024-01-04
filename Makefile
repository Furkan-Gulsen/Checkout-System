APP=TY_Case

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