go:
	go run cmd/main.go

tests:
	go test -v ./...

make build:
	docker build . -t tsarka-test-task

make run: 
	docker run -p 8080:8080 tsarka-test-task 