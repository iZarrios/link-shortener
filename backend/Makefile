build:
	@go build  -o bin ./cmd

run: build
	@./bin
createdb:
	docker exec -i  auth-psql psql -U postgres -c "CREATE DATABASE auth;"

dropdb:
	docker exec -i auth-psql psql -U postgres -c "DROP DATABASE auth;"

up:
	docker-compose up -d

down:
	docker-compose down

