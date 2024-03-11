clean:
	rm -rf ./bin
	rm coverage.out

build:
	go build -o bin/cmd ./cmd
	go build -o bin/4 ./drafts/4

clean_build: clean build

run:
	go run ./...

fmt:
	gofmt -w ./

test:
	go clean -testcache
	go test ./... -cover -coverprofile=coverage.out

cov:
	go tool cover -html=coverage.out
