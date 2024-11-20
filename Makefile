format:
	gofmt -s -w .

test:
	go test ./... -v --cover

build:
	docker build -t receipt-processor .

run:
	go run webservice.go
