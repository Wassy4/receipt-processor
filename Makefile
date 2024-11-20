format:
	gofmt -s -w .

lint:
	format golangci-lint run ./...

test:
	lint go test ./... -v --cover

build:
	docker build -t receipt-processor .

run:
	go run webservice.go
