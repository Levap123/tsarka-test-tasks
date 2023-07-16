go:
	go run cmd/main.go

tests:
	go test -v ./...

make run: 
	docker-compose up --build

make migrate:
	goose -dir=migrations postgres "user=dev password=dev dbname=dev sslmode=disable" up
