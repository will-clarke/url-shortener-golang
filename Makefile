BINARY_NAME=url-shortener

test:
	go test ./... -cover

build_linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/linux/$(BINARY_NAME)

build:
	go build -o $(BINARY_NAME) -v

run_server_redirect:
	go run main.go -server true -redirect true

clean:
	go clean
	rm -f $(BINARY_NAME)

install:
	go mod download

lint:
	golangci-lint run ./...
