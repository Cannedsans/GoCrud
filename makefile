.PHONY: default run build test docs clean
#variables
APP_NAME=Crud-GoOportunitis

# tasks
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
