.PHONY: build, test

GIT_SHA=`git rev-parse --short HEAD || echo`

build:
	@echo "Building scooter..."
	@mkdir -p bin
	@go build -ldflags "-X main.GitSHA=${GIT_SHA}" -o bin/scooter .

test:
	@echo "Running tests..."
	@go test -v
