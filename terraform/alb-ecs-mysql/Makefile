build:
	GOOS=linux GOARCH=arm64 go build -o api main.go
	zip lambda.zip hello

dc-up:
	docker-compose up -d localstack

dc-down:
	docker-compose down
