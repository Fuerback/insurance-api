img := insurance-api

docker-up: build docker-image docker-run

docker-image:
	docker build -t $(img) .

docker-run:
	docker run -p 8080:8000 $(img)

docker-tests:
	docker run -p 8080:8000 $(img) go test ./... -cover

run-local:
	go run .

build:
	go build .

run-tests:
	go test ./... -cover