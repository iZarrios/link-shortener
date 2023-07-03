build:
	@go build  -o bin ./cmd

run: build
	@./bin

# createdb and dropdb are  not necessary anymore 
# as GORM has can create the database automatically(autoMigrate)
createdb:
	docker exec -i  auth-psql psql -U postgres -c "CREATE DATABASE auth;"

dropdb:
	docker exec -i auth-psql psql -U postgres -c "DROP DATABASE auth;"

