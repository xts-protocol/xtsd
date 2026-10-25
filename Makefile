build:
	go build -o xtsd ./cmd/xtsd

run:
	go run ./cmd/xtsd

docker-build:
	docker build -t xtsd .

docker-run:
	docker run -p 8080:8080 xtsd

fmt:
	go fmt ./...

test:
	go test ./...
