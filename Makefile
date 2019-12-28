.PHONY: build, test

Version=`git describe --abbrev=0 --tags`
GitSHA=`git rev-parse --short HEAD`

build:
	@echo "Building scooter..."
	@mkdir -p bin
	@go build -ldflags "-X main.Version=${Version} -X main.GitSHA=${GitSHA}" -o bin/scooter .

test:
	@echo "Running tests..."
	@go test -v ./...
