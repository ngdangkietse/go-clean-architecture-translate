run:
	go run main.go
build:
	go build main.go
dev-up:
	docker-compose -f docker-compose.dev.yml up
dev-up-d:
	docker-compose -f docker-compose.dev.yml up -d
dev-down:
	docker-compose -f docker-compose.dev.yml down
prod-up:
	docker-compose -f docker-compose.prod.yml up
prod-up-d:
	docker-compose -f docker-compose.prod.yml up -d
prod-down:
	docker-compose -f docker-compose.prod.yml down
