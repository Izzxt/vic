.PHONY: start

start:
	go run cmd/main.go

dump:
	go run database/dump/dump.go 

gen:
	sqlc generate -f ./database/database.json