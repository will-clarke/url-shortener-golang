test:
	go test ./... -cover

build_linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

run:
	go run url-shortener.go -server true -redirect true

install:
	go install

lint:
	golangci-lint run ./...
