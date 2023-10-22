up-migrate:
	goose -dir ./schema postgres "user=root password=qwerty host=localhost dbname=shop sslmode=disable  port=5432" up
run-dev:
	docker-compose up -d
run-dev-back:
	docker-compose up -d
	go run cmd/main.go
down:
	docker-compose down