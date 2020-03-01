BINARY_NAME=url-shortener

test:
	go test ./... -cover

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
