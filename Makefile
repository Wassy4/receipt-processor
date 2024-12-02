format:
	gofmt -s -w .

test:
	go test ./... -v --cover

build:
	docker build -t receipt-processor .

run:
	docker run -d -p 8080:8080 --name alex-wasserman-receipt-processor receipt-processor

stop:
	docker stop alex-wasserman-receipt-processor

delete:
	docker rm alex-wasserman-receipt-processor