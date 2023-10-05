.PHONY: default run build test docs clean

#Variables

APP_NAME = go-rest-api

# Task
default: run

run:
	@go run main.go

build:
	@go build -o $(APP_NAME) main.go

test: 
	@go test ./ ...

docs:
	@swag init

clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs
