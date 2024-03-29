.PHONY:  setup build build-css clean dev test release

default: build

setup:
	go mod download
	npm install

build: clean build-css
	CGO_ENABLED=1 go build -o ./dist/onepage -a -ldflags '-s' -installsuffix cgo cmd/onepage/main.go

build-css:
	npm run build

clean:
	rm -rf ./dist/*

dev:
	gow -c -v -e=go -e=mod -e=html -e=css  run cmd/onepage/main.go

test:
	go test -v ./...

release:
	goreleaser --clean