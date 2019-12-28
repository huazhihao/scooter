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

release:
	@for platform in darwin linux windows; do \
		if [ $$platform == windows ]; then extension=.exe; fi; \
		GOOS=$$platform GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w -X main.Version=${Version} -X main.GitSHA=${GitSHA}" -o bin/confd-${Version}-$$platform-amd64$$extension; \
	done
	@GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -ldflags="-s -w  -X main.Version=${Version} -X main.GitSHA=${GitSHA}" -o bin/confd-${Version}-linux-arm64;